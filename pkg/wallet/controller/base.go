package controller

import (
	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/wallet"
	"github.com/sonrhq/core/x/identity/types"
)

// `DIDController` is a type that is both a `wallet.Wallet` and a `store.WalletStore`.
// @property GetChallengeResponse - This method is used to get the challenge response from the DID
// controller.
// @property RegisterAuthenticationCredential - This is the method that will be called when the user
// clicks on the "Register" button.
// @property GetAssertionOptions - This method is used to get the options for the assertion.
// @property AuthorizeCredential - This is the method that will be called when the user clicks the
// "Login" button on the login page.
type DIDController interface {
	// Address
	Address() string

	// DID
	ID() string

	// DID Document
	DidDocument() *types.DidDocument

	// Creates a new account and returns the address of the account.
	CreateAccount(name string, coinType crypto.CoinType) (*types.VerificationMethod, error)

	// Gets an account by name
	GetAccounts(ct crypto.CoinType) ([]wallet.Account, error)

	// Gets all accounts
	ListAccounts() ([]wallet.Account, error)

	// Path returns the path of the wallet store
	Path() string

	// Sign a message with the primary account
	SignWithAccount(message []byte, account string) ([]byte, error)

	// Verify a message with the primary account
	VerifyWithAccount(message []byte, signature []byte, account string) (bool, error)
}
