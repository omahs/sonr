package resolver

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/pkg/node"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                          Global Resolver Store Methods                         ||
// ! ||--------------------------------------------------------------------------------||

// InsertKeyShare inserts a record into the IPFS store for the given controller
func InsertKeyShare(key string, value interface{}) error {
	err := setupKeyshareStore()
	if err != nil {
		return err
	}

	var vBiz []byte
	switch value.(type) {
	case string:
		v, err := base64.StdEncoding.DecodeString(value.(string))
		if err != nil {
			return err
		}
		vBiz = v
	case []byte:
		vBiz = value.([]byte)
	default:
		return fmt.Errorf("value must be a string or []byte")
	}
	return ksStore.Put(key, vBiz)
}

// GetKeyShare gets a record from the IPFS store for the given controller
func GetKeyShare(key string) ([]byte, error) {
	err := setupKeyshareStore()
	if err != nil {
		return nil, err
	}
	vBiz, err := ksStore.Get(key)
	if err != nil {
		return nil, err
	}
	return vBiz, nil
}

// DeleteKeyShare deletes a record from the IPFS store for the given controller
func DeleteKeyShare(key string) error {
	err := setupKeyshareStore()
	if err != nil {
		return err
	}
	return ksStore.Delete(key)
}

// ListKeyShares lists all records in the IPFS store for the given controller
func ListKeyShares() (map[string][]byte, error) {
	err := setupKeyshareStore()
	if err != nil {
		return nil, err
	}
	m := make(map[string][]byte)
	for k, v := range ksStore.All() {
		m[k] = v
	}
	return m, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||              IPFS Based Wallet Store Implementation using OrbitDB              ||
// ! ||--------------------------------------------------------------------------------||

type ipfsKsStore struct {
	controller string
	node.IPFSKVStore
}

func makeIpfsKsStore(store node.IPFSKVStore, controller string) *ipfsKsStore {
	return &ipfsKsStore{
		IPFSKVStore: store,
		controller:  controller,
	}
}

func (s *ipfsKsStore) All() map[string][]byte {
	return s.IPFSKVStore.All()
}

func (s *ipfsKsStore) Address() string {
	return s.IPFSKVStore.Address().String()
}

func (s *ipfsKsStore) DBName() string {
	return s.IPFSKVStore.DBName()
}

func (s *ipfsKsStore) Identity() string {
	return s.IPFSKVStore.Identity().ID
}

func (s *ipfsKsStore) PublicKey() []byte {
	return s.IPFSKVStore.Identity().PublicKey
}

func (s *ipfsKsStore) Type() string {
	return s.IPFSKVStore.Type()
}

func (s *ipfsKsStore) Close() error {
	return s.Close()
}

func (s *ipfsKsStore) Get(key string) ([]byte, error) {
	return s.IPFSKVStore.Get(context.Background(), key)
}

func (s *ipfsKsStore) Put(key string, value []byte) error {
	_, err := s.IPFSKVStore.Put(context.Background(), key, value)
	return err
}

func (s *ipfsKsStore) Delete(key string) error {
	_, err := s.IPFSKVStore.Delete(context.Background(), key)
	return err
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                         Helper Methods for Module Setup                        ||
// ! ||--------------------------------------------------------------------------------||
var (
	ksStore *ipfsKsStore
	inStore node.IPFSDocsStore
)

// setupKeyshareStore initializes the global keyshare store
func setupKeyshareStore() error {
	if ksStore != nil {
		return nil
	}
	snrctx := local.NewContext()
	kv, err := node.OpenKeyValueStore(context.Background(), snrctx.GlobalKsStore)
	if err != nil {
		return err
	}
	ksStore = makeIpfsKsStore(kv, snrctx.GlobalKsStore)
	return nil
}
