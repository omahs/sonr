package store

import (
	"context"

	"berty.tech/go-orbit-db/iface"
	"github.com/sonrhq/core/pkg/node/config"
	vaultv1 "github.com/sonrhq/core/x/identity/types/vault/v1"
)

type IPFSStore struct {
	accConfig   *vaultv1.AccountConfig
	ipfsKVStore iface.KeyValueStore
}

func newIPFSStore(node config.IPFSNode, accCfg *vaultv1.AccountConfig) (WalletStore, error) {
	docs, err := node.LoadKeyValueStore(accCfg.Address)
	if err != nil {

		return nil, err
	}
	ds := &IPFSStore{
		accConfig:   accCfg,
		ipfsKVStore: docs,
	}
	return ds, nil
}

func (ds *IPFSStore) GetShare(name string) (*vaultv1.ShareConfig, error) {
	bz, err := ds.ipfsKVStore.Get(context.Background(), name)
	if err != nil {
		return nil, err
	}

	sc := &vaultv1.ShareConfig{}
	if err := sc.Unmarshal(bz); err != nil {
		return nil, err
	}
	return sc, nil
}

func (ds *IPFSStore) SetShare(sc *vaultv1.ShareConfig) error {
	bz, err := sc.Marshal()
	if err != nil {
		return err
	}
	_, err = ds.ipfsKVStore.Put(context.Background(), sc.SelfId, bz)
	if err != nil {
		return err
	}
	return nil
}

// JWKClaims returns the JWKClaims for the store to be signed by the identity
func (ds *IPFSStore) JWKClaims() (string, error) {
	return "", nil
}

// VerifyJWKClaims verifies the JWKClaims for the store
func (ds *IPFSStore) VerifyJWKClaims(claims string) error {
	return nil
}
