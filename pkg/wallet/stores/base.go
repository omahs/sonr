package stores

import (
	"os"
	"path/filepath"

	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/pkg/wallet"
	"github.com/sonrhq/core/pkg/wallet/stores/internal"
)

const (
	DEFAULT_WALLET_PATH = "_SNR_WALLET_"
)

// NewWalletStore returns a new WalletStore
func New(acc wallet.Account, opts ...Option) (wallet.Store, error) {
	userHomeDir, _ := os.UserHomeDir()
	testWalOut := filepath.Join(userHomeDir, DEFAULT_WALLET_PATH)
	cfg := &storeConfig{
		acc:  acc,
		path: testWalOut,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	return cfg.Apply()
}

// storeConfig is the configuration for the store
type storeConfig struct {
	acc      wallet.Account
	path     string
	ipfsNode common.IPFSNode
	isIPFS   bool
}

// Apply applies the configuration to the store
func (cfg *storeConfig) Apply() (wallet.Store, error) {
	if cfg.isIPFS {
		return internal.NewIPFSStore(cfg.ipfsNode, cfg.acc.Config())
	}
	err := os.MkdirAll(cfg.path, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return internal.NewFileStore(cfg.path, cfg.acc.Config())
}

// Option is a function that configures the store
type Option func(*storeConfig)

// SetFileStorePath sets the base path to use a file store. You can provide a list of paths to append, The default wallet name will be appended to the path.
func SetFileStorePath(path ...string) Option {
	return func(cfg *storeConfig) {
		cfg.path = filepath.Join(path...)
		cfg.path = filepath.Join(cfg.path, DEFAULT_WALLET_PATH)
		cfg.isIPFS = false
	}
}

// SetIPFSStore sets the store to use an IPFS store
func SetIPFSStore(ipfsNode common.IPFSNode) Option {
	return func(cfg *storeConfig) {
		cfg.ipfsNode = ipfsNode
		cfg.isIPFS = true
	}
}

// ResetWalletStore resets the wallet store
func ResetWalletStore(path ...string) error {
	testWalOut := filepath.Join(path...)
	testWalOut = filepath.Join(testWalOut, DEFAULT_WALLET_PATH)
	return os.RemoveAll(testWalOut)
}
