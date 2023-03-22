package controller

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/crypto/mpc"
	"github.com/sonrhq/core/pkg/node"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

type Controller interface {
	// Get the controller's DID
	Did() string

	// Get the controller's DID document
	DidDocument() *types.DidDocument

	// Authorize the client to access the controller's account
	Authorize(cred *crypto.WebauthnCredential) error

	// CreateAccount creates a new account for the controller
	CreateAccount(name string, coinType crypto.CoinType) error

	// GetAccount returns the controller's account
	GetAccount(did string) (Account, error)

	// ListAccounts returns the controller's accounts
	ListAccounts() ([]Account, error)

	// Sign signs a message with the controller's account
	Sign(did string, msg []byte) ([]byte, error)

	// Verify verifies a signature with the controller's account
	Verify(did string, msg []byte, sig []byte) (bool, error)
}

type didController struct {
	ipfsStore node.IPFSStore
	primary Account
	didDoc *types.DidDocument
}

func NewController(ctx context.Context, credential *crypto.WebauthnCredential) (Controller, error) {
	doneCh := make(chan Account)
	errCh := make(chan error)

	go generateInitialAccount(ctx, credential, doneCh, errCh)

	select {
	case acc := <-doneCh:
		return setupController(ctx, credential, acc)
	case err := <-errCh:
		return nil, err
	}
}

func LoadController(ctx context.Context, credential *crypto.WebauthnCredential, address string) (Controller, error) {
	doneCh := make(chan Account)
	errCh := make(chan error)

	go generateInitialAccount(ctx, credential, doneCh, errCh)

	select {
	case acc := <-doneCh:
		return setupController(ctx, credential, acc)
	case err := <-errCh:
		return nil, err
	}
}

func (dc *didController) Did() string {
	return dc.primary.DID()
}

func (dc *didController) DidDocument() *types.DidDocument {
	return dc.didDoc
}

func (dc *didController) Authorize(cred *crypto.WebauthnCredential) error {
	return nil
}

func (dc *didController) CreateAccount(name string, coinType crypto.CoinType) error {
	kss, err := dc.primary.ListKeyshares()
	if err != nil {
		return err
	}
	var cmpcnfs []*cmp.Config
	for _, ks := range kss {
		cmpcnfs = append(cmpcnfs, ks.Config())
	}

	var newKss []KeyShare
	for _, conf := range cmpcnfs {
		newConf, err := conf.DeriveBIP32(uint32(coinType.BipPath()))
		if err != nil {
			return err
		}
		ksb, err := newConf.MarshalBinary()
		if err != nil {
			return err
		}
		ks, err := NewKeyshare(string(newConf.ID), ksb, coinType, name)
		if err != nil {
			return err
		}
		newKss = append(newKss, ks)
	}
	newAcc := NewAccount(newKss)
	newAcc.Sync(dc.ipfsStore)
	return nil
}

// GetAccount returns the controller's account from the Address
func (dc *didController) GetAccount(address string) (Account, error) {
	mapkv := dc.ipfsStore.All()
	mapkv = filterByAccountName(mapkv, address)
	if len(mapkv) == 0 {
		return nil, fmt.Errorf("account not found")
	}
	var kss []KeyShare
	for k, v := range mapkv {
		ks, err := LoadKeyshareFromStore(k, v)
		if err != nil {
			return nil, err
		}
		kss = append(kss, ks)
	}
	return NewAccount(kss), nil
}

// ListAccounts returns the controller's accounts
func (dc *didController) ListAccounts() ([]Account, error) {
	mapkv := dc.ipfsStore.All()
	var accs []Account
	mapkv = filterByCoin(mapkv, crypto.SONRCoinType)
	for k := range mapkv {
		acc, err := dc.GetAccount(k)
		if err != nil {
			return nil, err
		}
		accs = append(accs, acc)
	}
	mapkv = filterByCoin(mapkv, crypto.ETHCoinType)
	for k := range mapkv {
		acc, err := dc.GetAccount(k)
		if err != nil {
			return nil, err
		}
		accs = append(accs, acc)
	}
	mapkv = filterByCoin(mapkv, crypto.BTCCoinType)
	for k := range mapkv {
		acc, err := dc.GetAccount(k)
		if err != nil {
			return nil, err
		}
		accs = append(accs, acc)
	}
	return accs, nil
}

func (dc *didController) Sign(did string, msg []byte) ([]byte, error) {
	return nil, nil
}

func (dc *didController) Verify(did string, msg []byte, sig []byte) (bool, error) {
	return false, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                          Helper Methods for Controller                         ||
// ! ||--------------------------------------------------------------------------------||


func generateInitialAccount(ctx context.Context, credential *crypto.WebauthnCredential, doneCh chan Account, errChan chan error) {
	shardName := crypto.PartyID(base64.RawStdEncoding.EncodeToString(credential.Id))
	// Call Handler for keygen
	confs, err := mpc.Keygen(shardName, 1, []crypto.PartyID{"vault"})
	if err != nil {
		errChan <- err
	}

	var kss []KeyShare
	for _, conf := range confs {
		ksb, err := conf.MarshalBinary()
		if err != nil {
			errChan <- err
		}
		ks, err := NewKeyshare(string(conf.ID), ksb, crypto.SONRCoinType, "Primary")
		if err != nil {
			errChan <- err
		}
		kss = append(kss, ks)
	}
	doneCh <- NewAccount(kss)
}

func setupController(ctx context.Context, credential *crypto.WebauthnCredential,  primary Account) (Controller, error) {
	didDoc := types.NewBlankDocument(primary.DID())

	st, err := node.NewIPFSStore(context.Background(), primary.PubKey())
	if err != nil {
		return nil, err
	}

	err = primary.Sync(st)
	if err != nil {
		return nil, err
	}
	didDoc.AddService(st.Service())

	return &didController{
		ipfsStore: st,
		primary: primary,
		didDoc: didDoc,
	}, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                              Map & Slice Filtering                             ||
// ! ||--------------------------------------------------------------------------------||

func filterMap(m map[string][]byte, f func(string) bool) map[string][]byte {
	n := make(map[string][]byte)
	for k, v := range m {
		if f(k) {
			n[k] = v
		}
	}
	return n
}

func filterByCoin(m map[string][]byte, ct crypto.CoinType) map[string][]byte {
	return filterMap(m, func(k string) bool {
		ksr, err := ParseKeyShareDid(k)
		if err != nil {
			return false
		}
		return ksr.CoinType == ct
	})
}

func filterByCoinAndIndex(m map[string][]byte, ct crypto.CoinType, idx int) map[string][]byte {
	i := 0
	return filterMap(m, func(k string) bool {
		ksr, err := ParseKeyShareDid(k)
		if err != nil {
			return false
		}
		if ksr.CoinType == ct {
			i++
		}
		return ksr.CoinType == ct && i == idx
	})
}

func filterByAccountName(m map[string][]byte, name string) map[string][]byte {
	return filterMap(m, func(k string) bool {
		ksr, err := ParseKeyShareDid(k)
		if err != nil {
			return false
		}
		return ksr.AccountName == name
	})
}
