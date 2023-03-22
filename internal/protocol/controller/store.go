package controller

import (
	"context"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/node"
)

type WalletStore interface {
	// GetWallet returns the wallet with the given name.
	Address() string

	// GetWallet returns the wallet with the given name.
	GetAccount(name string) (Account, error)

	// ListWallets returns a list of all wallets.
	ListAccounts() ([]Account, error)

	// CreateWallet creates a new wallet with the given name.
	CreateAccount(name string, coinType crypto.CoinType) (Account, error)
}

type walletStore struct {
	ipfsStore node.IPFSStore
	primary Account
}

func NewWalletStore(primary Account) (WalletStore, error) {
	st, err := node.NewIPFSStore(context.Background(), primary.PubKey())
	if err != nil {
		return nil, err
	}

	err = primary.Sync(st)
	if err != nil {
		return nil, err
	}
	return &walletStore{
		ipfsStore: st,
		primary: primary,
	}, nil
}

func (s *walletStore) Address() string {
	return s.primary.Address()
}

func (s *walletStore) GetAccount(name string) (Account, error) {
	return nil, nil
}

func (s *walletStore) ListAccounts() ([]Account, error) {
	return nil, nil
}

func (s *walletStore) CreateAccount(name string, coinType crypto.CoinType) (Account, error) {

	return nil, nil
}

func (s *walletStore) SignTransaction(name string, tx []byte) ([]byte, error) {
	return nil, nil
}

func (s *walletStore) VerifyMessage(name string, message []byte, signature []byte) (bool, error) {
	return false, nil
}

func (s *walletStore) GetKeyShare(cid string) (KeyShare, error) {
	return nil, nil
}

func (s *walletStore) PutKeyShare(ks KeyShare) error {
	return nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                       Helper Functions for Map and Slice                       ||
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
