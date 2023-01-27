package store

import (
	"github.com/sonrhq/core/pkg/node/config"
	v1 "github.com/sonrhq/core/x/identity/types/vault/v1"
	vaultv1 "github.com/sonrhq/core/x/identity/types/vault/v1"
)

type WalletStore interface {
	// GetShare returns a *cmp.Config for the given name
	GetShare(name string) (*v1.ShareConfig, error)

	// PutShare stores the given *cmp.Config under the given name
	SetShare(share *v1.ShareConfig) error

	// JWKClaims returns the JWKClaims for the store to be signed by the identity
	JWKClaims() (string, error)

	// VerifyJWKClaims verifies the JWKClaims for the store
	VerifyJWKClaims(claims string) error
}

// NewWalletStore returns a new WalletStore
func NewWalletStore(cnfg *vaultv1.AccountConfig, opts ...Option) (WalletStore, error) {
	cfg := &storeConfig{
		accCfg: cnfg,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg.Apply()
}

// storeConfig is the configuration for the store
type storeConfig struct {
	accCfg   *vaultv1.AccountConfig
	path     string
	ipfsNode config.IPFSNode
	isFile   bool
	isIPFS   bool
}

// Apply applies the configuration to the store
func (cfg *storeConfig) Apply() (WalletStore, error) {
	if cfg.isFile {
		return newFileStore(cfg.path, cfg.accCfg)
	}
	if cfg.isIPFS {
		return newIPFSStore(cfg.ipfsNode, cfg.accCfg)
	}
	return newMemoryStore(cfg.accCfg)
}

// Option is a function that configures the store
type Option func(*storeConfig)

// SetFileStore sets the store to use a file store
func SetFileStore(path string) Option {
	return func(cfg *storeConfig) {
		cfg.path = path
		cfg.isFile = true
		cfg.isIPFS = false
	}
}

// SetIPFSStore sets the store to use an IPFS store
func SetIPFSStore(ipfsNode config.IPFSNode) Option {
	return func(cfg *storeConfig) {
		cfg.ipfsNode = ipfsNode
		cfg.isFile = false
		cfg.isIPFS = true
	}
}
