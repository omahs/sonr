package controller

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/crypto/mpc"
	"github.com/sonrhq/core/pkg/node"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

const PrimaryAccountName = "primary"

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
	GetAccount(name string, coinType crypto.CoinType) (Account, error)

	// ListAccounts returns the controller's accounts
	ListAccounts(ct crypto.CoinType) ([]Account, error)

	// Sign signs a message with the controller's account
	Sign(name string, coinType crypto.CoinType, msg []byte) ([]byte, error)

	// Verify verifies a signature with the controller's account
	Verify(name string, coinType crypto.CoinType, msg []byte, sig []byte) (bool, error)
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

// LoadController loads a controller from the given DID document using the underlying IPFS store
func LoadController(ctx context.Context, credential *crypto.WebauthnCredential, didDoc *types.DidDocument) (Controller, error) {
	if len(didDoc.Service) == 0 {
		return nil, fmt.Errorf("no service found in DID document")
	}

	// Get the IPFS store service
	ipfsStore, err := node.NewIPFSStore(ctx, strings.Split( didDoc.Id, ":")[len(strings.Split( didDoc.Id, ":"))-1])
	if err != nil {
		return nil, err
	}

	// Get the primary account
	mapKv := ipfsStore.All()
	mapKv = filterByCoin(mapKv, crypto.SONRCoinType)
	if len(mapKv) == 0 {
		return nil, fmt.Errorf("no primary account found")
	}

	// Get the primary account
	var kss []KeyShare
	for k, v := range mapKv {
		ks, err := LoadKeyshareFromStore(k, v)
		if err != nil {
			return nil, err
		}
		kss = append(kss, ks)
	}
	primary := NewAccount(kss, crypto.SONRCoinType)
	return &didController{
		ipfsStore: ipfsStore,
		primary: primary,
		didDoc: didDoc,
	}, nil
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
	newAcc := NewAccount(newKss, coinType)
	newAcc.Sync(dc.ipfsStore)
	return nil
}

// GetAccount returns the controller's account from the Address
func (dc *didController) GetAccount(name string, coinType crypto.CoinType) (Account, error) {
	mapkv := dc.ipfsStore.All()
	mapkv = filterByCoinAndAccountName(mapkv, coinType, name)
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
	return NewAccount(kss, coinType), nil
}

// ListAccounts returns the controller's accounts
func (dc *didController) ListAccounts(ct crypto.CoinType) ([]Account, error) {
	mapkv := dc.ipfsStore.All()
	var accs []Account
	mapkv = filterByCoin(mapkv, ct)
	for k := range mapkv {
		acc, err := dc.GetAccount(k, ct)
		if err != nil {
			return nil, err
		}
		accs = append(accs, acc)
	}
	return accs, nil
}

// Sign signs a message with the controller's selected account
func (dc *didController) Sign(name string, coinType crypto.CoinType, msg []byte) ([]byte, error) {
	acc, err := dc.GetAccount(name, coinType)
	if err != nil {
		return nil, err
	}
	return acc.Sign(msg)
}

// Verify verifies a signature with the controller's selected account
func (dc *didController) Verify(name string, coinType crypto.CoinType, msg []byte, sig []byte) (bool, error) {
	acc, err := dc.GetAccount(name, coinType)
	if err != nil {
		return false, err
	}
	return acc.Verify(msg, sig)
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
	doneCh <- NewAccount(kss, crypto.SONRCoinType)
}

func setupController(ctx context.Context, credential *crypto.WebauthnCredential,  primary Account) (Controller, error) {
	didDoc := types.NewBlankDocument(primary.DID())

	st, err := node.NewIPFSStore(context.Background(), primary.Address())
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

func filterByCoinAndAccountName(m map[string][]byte, ct crypto.CoinType, name string) map[string][]byte {
	return filterMap(m, func(k string) bool {
		ksr, err := ParseKeyShareDid(k)
		if err != nil {
			return false
		}
		return ksr.CoinType == ct && ksr.AccountName == name
	})
}

