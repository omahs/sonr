package accounts

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sonrhq/core/pkg/wallet"
	"github.com/sonrhq/core/pkg/wallet/accounts/internal"
	v1 "github.com/sonrhq/core/x/identity/types/vault/v1"
)

// New creates a new account with the given options.
func New(opts ...Option) (wallet.Account, error) {
	c := defaultConfig()
	for _, opt := range opts {
		opt(c)
	}
	return c.Keygen()
}

// Load loads an account from a *crypto.AccountConfig.
func Load(ac *wallet.AccountConfig) wallet.Account {
	return internal.BaseAccountFromConfig(ac)
}

// LoadFromBytes loads an account from a byte slice.
func LoadFromBytes(b []byte) (wallet.Account, error) {
	accCfg := &v1.AccountConfig{}
	if err := accCfg.Unmarshal(b); err != nil {
		return nil, fmt.Errorf("failed to unmarshal account config: %w", err)
	}
	return Load(accCfg), nil
}

// LoadFromPath loads an account from a file path.
func LoadFromPath(path string) (wallet.Account, error) {
	// Open the file at the specified path
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the JSON-encoded AccountConfig from the file
	var accountConfig v1.AccountConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&accountConfig)
	if err != nil {
		return nil, err
	}
	return Load(&accountConfig), nil
}
