package v1

import (
	"strings"
)

// `NewWalletConfigFromRootAccount` creates a new wallet config from a root account
func NewWalletConfigFromRootAccount(account *AccountConfig) *WalletConfig {
	return &WalletConfig{
		Address:   account.Address,
		Algorithm: "cmp",
		Accounts: map[string]*AccountConfig{
			strings.ToLower(account.Name): account,
		},
	}
}
