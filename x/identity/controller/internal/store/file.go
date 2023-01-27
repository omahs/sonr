package store

import (
	vaultv1 "github.com/sonrhq/core/x/identity/types/vault/v1"
	bolt "go.etcd.io/bbolt"
)

type FileStore struct {
	accConfig *vaultv1.AccountConfig
	path      string
	db        *bolt.DB
	bucketKey []byte
}

func newFileStore(p string, accCfg *vaultv1.AccountConfig) (WalletStore, error) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(p, 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	ds := &FileStore{
		accConfig: accCfg,
		path:      p,
		db:        db,
		bucketKey: []byte(accCfg.Address),
	}
	return ds, nil
}

func (ds *FileStore) GetShare(name string) (*vaultv1.ShareConfig, error) {
	var sc *vaultv1.ShareConfig
	ds.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(ds.bucketKey)
		v := b.Get([]byte(name))
		sc = &vaultv1.ShareConfig{}
		return sc.Unmarshal(v)
	})
	return sc, nil
}

func (ds *FileStore) SetShare(sc *vaultv1.ShareConfig) error {
	ds.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(ds.bucketKey)
		v, err := sc.Marshal()
		if err != nil {
			return err
		}
		return b.Put([]byte(sc.SelfId), v)
	})
	return nil
}

// JWKClaims returns the JWKClaims for the store to be signed by the identity
func (ds *FileStore) JWKClaims() (string, error) {
	return "", nil
}

// VerifyJWKClaims verifies the JWKClaims for the store
func (ds *FileStore) VerifyJWKClaims(claims string) error {
	return nil
}
