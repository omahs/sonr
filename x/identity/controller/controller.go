package controller

import (
	"github.com/sonrhq/core/x/identity/protocol/vault/account"
	v1 "github.com/sonrhq/core/x/identity/types/vault/v1"
)

type DIDController interface {
	WalletConfig() *v1.WalletConfig

	// Creates a new account
	CreateAccount(name string, addrPrefix string, networkName string) error

	// Gets an account by name
	GetAccount(name string) (account.WalletAccount, error)

	// Gets Primary account
	PrimaryAccount() (account.WalletAccount, error)

	// Gets all accounts
	ListAccounts() ([]account.WalletAccount, error)
}
