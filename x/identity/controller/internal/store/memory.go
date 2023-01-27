package store

import (
	"fmt"
	"sync"

	vaultv1 "github.com/sonrhq/core/x/identity/types/vault/v1"
)

type MemoryStore struct {
	accConfig *vaultv1.AccountConfig
	configs   map[string]*vaultv1.ShareConfig
	sync.Mutex
}

func newMemoryStore(accCfg *vaultv1.AccountConfig) (WalletStore, error) {
	ds := &MemoryStore{
		accConfig: accCfg,
		configs:   make(map[string]*vaultv1.ShareConfig),
	}
	return ds, nil
}

func (ds *MemoryStore) GetShare(name string) (*vaultv1.ShareConfig, error) {
	ds.Lock()
	defer ds.Unlock()
	s, ok := ds.configs[name]
	if !ok {
		return nil, fmt.Errorf("share not found")
	}
	return s, nil
}

func (ds *MemoryStore) SetShare(sc *vaultv1.ShareConfig) error {
	ds.Lock()
	defer ds.Unlock()
	ds.configs[sc.SelfId] = sc
	return nil
}

// JWKClaims returns the JWKClaims for the store to be signed by the identity
func (ds *MemoryStore) JWKClaims() (string, error) {
	return "", nil
}

// VerifyJWKClaims verifies the JWKClaims for the store
func (ds *MemoryStore) VerifyJWKClaims(claims string) error {
	return nil
}
