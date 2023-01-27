package store

import (
	"context"

	"berty.tech/go-orbit-db/iface"
	"github.com/sonrhq/core/pkg/common"
	vaultv1 "github.com/sonrhq/core/x/identity/types/vault/v1"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

type IPFSStore struct {
	accConfig   *vaultv1.AccountConfig
	ipfsKVStore iface.KeyValueStore
}

func newIPFSStore(node common.IPFSNode, accCfg *vaultv1.AccountConfig) (WalletStore, error) {
	docs, err := node.LoadKeyValueStore(accCfg.DID())
	if err != nil {

		return nil, err
	}
	ds := &IPFSStore{
		accConfig:   accCfg,
		ipfsKVStore: docs,
	}
	return ds, nil
}

func (ds *IPFSStore) GetShare(name string) (*cmp.Config, error) {
	bz, err := ds.ipfsKVStore.Get(context.Background(), name)
	if err != nil {
		return nil, err
	}

	sc := cmp.EmptyConfig(curve.Secp256k1{})
	if err := sc.UnmarshalBinary(bz); err != nil {
		return nil, err
	}
	return sc, nil
}

func (ds *IPFSStore) SetShare(sc *cmp.Config) error {
	bz, err := sc.MarshalBinary()
	if err != nil {
		return err
	}
	_, err = ds.ipfsKVStore.Put(context.Background(), string(sc.ID), bz)
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
