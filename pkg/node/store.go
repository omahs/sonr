package node

import (
	"context"

	"github.com/sonrhq/core/x/identity/types"
)

type ServiceOperation string

type IPFSStore interface {
	// Address returns the full address of the store.
	Address() string

	// All returns all the keys and values of the store.
	All() map[string][]byte

	// DBName returns the name of the store.
	DBName() string

	// Identity returns the identity of the store.
	Identity() string

	// PublicKey returns the base64 public key of the store.
	PublicKey() []byte

	// Type returns the type of the store.
	Type() string

	// Close closes the store.
	Close() error

	// Get returns the value of the given key.
	Get(key string) ([]byte, error)

	// Put adds a new value to the store.
	Put(key string, value []byte) error

	// Delete removes a value from the store.
	Delete(key string) error

	// Service returns the DID service of the store.
	Service() *types.Service
}

type ipfsStore struct {
	address  string
	dbName   string
	identity string
	pubKey   []byte
	store    iPFSKVStore
	record   *types.Service
}

func makeIpfsStore(store iPFSKVStore, controller string) IPFSStore {
	return &ipfsStore{
		address:  store.Address().String(),
		dbName:   store.DBName(),
		identity: store.Identity().ID,
		pubKey:   store.Identity().PublicKey,
		store:    store,
		record:   types.NewIPFSStoreService(store.Address().String(), types.NewSonrID(controller)),
	}
}

func (s *ipfsStore) All() map[string][]byte {
	return s.store.All()
}

func (s *ipfsStore) Address() string {
	return s.address
}

func (s *ipfsStore) DBName() string {
	return s.dbName
}

func (s *ipfsStore) Identity() string {
	return s.identity
}

func (s *ipfsStore) PublicKey() []byte {
	return s.pubKey
}

func (s *ipfsStore) Type() string {
	return s.store.Type()
}

func (s *ipfsStore) Close() error {
	return s.store.Close()
}

func (s *ipfsStore) Get(key string) ([]byte, error) {
	return s.store.Get(context.Background(), key)
}

func (s *ipfsStore) Put(key string, value []byte) error {
	_, err := s.store.Put(context.Background(), key, value)
	return err
}

func (s *ipfsStore) Delete(key string) error {
	_, err := s.store.Delete(context.Background(), key)
	return err
}

func (s *ipfsStore) Service() *types.Service {
	return s.record
}
