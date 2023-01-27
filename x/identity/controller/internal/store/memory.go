package store

import (
	"fmt"
	"sync"

	vaultv1 "github.com/sonrhq/core/x/identity/types/vault/v1"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

type MemoryStore struct {
	accConfig *vaultv1.AccountConfig
	configs   map[string]*cmp.Config
	sync.Mutex
}

func newMemoryStore(accCfg *vaultv1.AccountConfig) (WalletStore, error) {
	ds := &MemoryStore{
		accConfig: accCfg,
		configs:   make(map[string]*cmp.Config),
	}
	return ds, nil
}

func (ds *MemoryStore) GetShare(name string) (*cmp.Config, error) {
	ds.Lock()
	defer ds.Unlock()
	s, ok := ds.configs[name]
	if !ok {
		return nil, fmt.Errorf("share not found")
	}
	return s, nil
}

func (ds *MemoryStore) SetShare(sc *cmp.Config) error {
	ds.Lock()
	defer ds.Unlock()
	ds.configs[string(sc.ID)] = sc
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
