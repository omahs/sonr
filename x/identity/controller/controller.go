package controller

import (
	"context"
	"errors"
	"strings"

	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/x/identity/controller/internal/store"
	"github.com/sonrhq/core/x/identity/protocol/vault/account"
	"github.com/sonrhq/core/x/identity/types"
	vaultv1 "github.com/sonrhq/core/x/identity/types/vault/v1"
)

// rootWalletAccountName is the name of the root account
const rootWalletAccountName = "Primary"

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
	Document() *types.DidDocument

	// This method is used to get the challenge response from the DID controller.
	// GetChallengeOptions(aka string) (*v1.ChallengeResponse, error)

	// This is the method that will be called when the user clicks on the "Register" button.
	// RegisterAuthenticationCredential(credentialCreationData string) (*v1.RegisterResponse, error)

	// This method is used to get the options for the assertion.
	// GetAssertionOptions(aka string) (*v1.AssertResponse, error)

	// This is the method that will be called when the user clicks the "Login" button on the login page.
	// AuthorizeCredential(credentialRequestData string) (*v1.LoginResponse, error)

	// Creates a new account
	CreateAccount(name string, coinType common.CoinType) error

	// Gets an account by name
	GetAccount(name string) (account.WalletAccount, error)

	// Gets Primary account
	PrimaryAccount() (account.WalletAccount, error)

	// Gets all accounts
	ListAccounts() ([]account.WalletAccount, error)
}

// `DIDControllerImpl` is a type that implements the `DIDController` interface.
// @property  - `wallet.Wallet`: This is the interface that the DID controller implements.
// @property  - `store.WalletStore`: This is the interface that the DID controller implements.
type DIDControllerImpl struct {
	store store.WalletStore

	ctx context.Context
	aka string

	accounts       map[string]*vaultv1.AccountConfig
	didDocument    *types.DidDocument
	rootPubKey     *types.PubKey
	primaryAccount *vaultv1.AccountConfig
}

// `New` creates a new DID controller instance
func New(ctx context.Context, account *vaultv1.AccountConfig) (DIDController, error) {
	docc := &DIDControllerImpl{
		ctx:            ctx,
		primaryAccount: account,
		accounts:       make(map[string]*vaultv1.AccountConfig),
	}

	pubKey, err := account.PubKey()
	if err != nil {
		return nil, err
	}
	// Create the DID document.
	doc, err := types.NewDocument(pubKey)
	if err != nil {
		return nil, err
	}
	docc.didDocument = doc
	docc.rootPubKey = pubKey
	return docc, nil
}

// Address returns the address of the DID controller.
func (d *DIDControllerImpl) Address() string {
	addr, _ := d.primaryAccount.Address()
	return addr
}

// ID returns the DID of the DID controller.
func (d *DIDControllerImpl) ID() string {
	return d.primaryAccount.DID()
}

// Document returns the DID document of the DID controller.
func (d *DIDControllerImpl) Document() *types.DidDocument {
	return d.didDocument
}

// // This method is used to get the challenge response from the DID controller.
// func (d *DIDControllerImpl) GetChallengeOptions(aka string) (*v1.ChallengeResponse, error) {
// 	return nil, nil
// }

// // This is the method that will be called when the user clicks on the "Register" button.
// func (d *DIDControllerImpl) RegisterAuthenticationCredential(credentialCreationData string) (*v1.RegisterResponse, error) {
// 	return nil, nil
// }

// // This method is used to get the options for the assertion.
// func (d *DIDControllerImpl) GetAssertionOptions(aka string) (*v1.AssertResponse, error) {
// 	return nil, nil
// }

// // This is the method that will be called when the user clicks the "Login" button on the login page.
// func (d *DIDControllerImpl) AuthorizeCredential(credentialRequestData string) (*v1.LoginResponse, error) {
// 	return nil, nil
// }

// Creating a new account.
func (w *DIDControllerImpl) CreateAccount(name string, coinType common.CoinType) error {
	prim, err := w.PrimaryAccount()
	if err != nil {
		return err
	}
	acc, err := prim.Bip32Derive(name, coinType)
	if err != nil {
		return err
	}
	w.accounts[name] = acc.AccountConfig()
	pub, err := acc.AccountConfig().PubKey()
	if err != nil {
		return err
	}
	addr, err := pub.Bech32(acc.AccountConfig().CoinType().AddrPrefix())
	if err != nil {
		return err
	}
	err = w.didDocument.SetAssertion(pub, types.WithBlockchainAccount(addr), types.WithController(w.didDocument.Id), types.WithIDFragmentSuffix(acc.AccountConfig().Name))
	if err != nil {
		return err
	}
	return nil
}

// Returning the account.WalletAccount and error.
func (w *DIDControllerImpl) GetAccount(name string) (account.WalletAccount, error) {
	accConf, ok := w.accounts[strings.ToLower(name)]
	if !ok {
		return nil, errors.New("Account not found")
	}
	return account.NewAccountFromConfig(accConf)
}

// Returning a list of accounts.
func (w *DIDControllerImpl) ListAccounts() ([]account.WalletAccount, error) {
	accs := make([]account.WalletAccount, 0, len(w.accounts))
	for _, accConf := range w.accounts {
		acc, err := account.NewAccountFromConfig(accConf)
		if err != nil {
			return nil, err
		}
		accs = append(accs, acc)
	}
	return accs, nil
}

// Returning the primary account.
func (w *DIDControllerImpl) PrimaryAccount() (account.WalletAccount, error) {
	return account.NewAccountFromConfig(w.primaryAccount)
}
