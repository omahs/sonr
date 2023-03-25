package resolver

import (
	"context"

	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/pkg/node"
)

type KVStoreItem interface {
	Bytes() []byte
	Did() string
}

type BasicStoreItem struct {
	did  string
	data []byte
}

func (i *BasicStoreItem) Bytes() []byte {
	return i.data
}

func (i *BasicStoreItem) Did() string {
	return i.did
}

func NewBasicStoreItem(did string, data []byte) *BasicStoreItem {
	return &BasicStoreItem{
		did:  did,
		data: data,
	}
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                          Global Resolver Store Methods                         ||
// ! ||--------------------------------------------------------------------------------||

// InsertKeyShare inserts a record into the IPFS store for the given controller
func InsertKeyShare(i KVStoreItem) error {
	err := setupKeyshareStore()
	if err != nil {
		return err
	}

	return ksStore.Put(i.Did(), i.Bytes())
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
func ListKeyShares() ([]KVStoreItem, error) {
	err := setupKeyshareStore()
	if err != nil {
		return nil, err
	}
	m := make([]KVStoreItem, 0)
	for k, v := range ksStore.All() {
		m = append(m, NewBasicStoreItem(k, v))
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
