package controller

import (
	"context"

	"github.com/sonrhq/core/x/identity/controller/internal/store"
	"github.com/sonrhq/core/x/identity/protocol/vault/account"
	"github.com/sonrhq/core/x/identity/protocol/vault/wallet"
	"github.com/sonrhq/core/x/identity/types"
	v1 "github.com/sonrhq/core/x/identity/types/auth/v1"
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
	wallet.Wallet
	GetChallengeOptions(aka string) (*v1.ChallengeResponse, error)
	RegisterAuthenticationCredential(credentialCreationData string) (*v1.RegisterResponse, error)
	GetAssertionOptions(aka string) (*v1.AssertResponse, error)
	AuthorizeCredential(credentialRequestData string) (*v1.LoginResponse, error)
}

// `DIDControllerImpl` is a type that implements the `DIDController` interface.
// @property  - `wallet.Wallet`: This is the interface that the DID controller implements.
// @property  - `store.WalletStore`: This is the interface that the DID controller implements.
type DIDControllerImpl struct {
	wallet.Wallet
	store store.WalletStore

	ctx context.Context
	aka string

	didDocument    *types.DidDocument
	primaryAccount account.WalletAccount
}

func New(ctx context.Context, wallet wallet.Wallet) (DIDController, error) {
	docc := &DIDControllerImpl{
		Wallet: wallet,
		ctx:    ctx,
	}

	// Get the primary account.
	primaryAccount, err := wallet.PrimaryAccount()
	if err != nil {
		return nil, err
	}
	docc.primaryAccount = primaryAccount

	// Create the DID document.
	doc, err := types.NewDocument(primaryAccount.PubKey())
	if err != nil {
		return nil, err
	}
	docc.didDocument = doc
	return docc, nil
}

func (d *DIDControllerImpl) GetChallengeOptions(aka string) (*v1.ChallengeResponse, error) {
	return nil, nil
}
func (d *DIDControllerImpl) RegisterAuthenticationCredential(credentialCreationData string) (*v1.RegisterResponse, error) {
	return nil, nil
}
func (d *DIDControllerImpl) GetAssertionOptions(aka string) (*v1.AssertResponse, error) {
	return nil, nil
}
func (d *DIDControllerImpl) AuthorizeCredential(credentialRequestData string) (*v1.LoginResponse, error) {
	return nil, nil
}
