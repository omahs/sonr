package controller

import (
	"context"
	"errors"
	"fmt"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/wallet"
	"github.com/sonrhq/core/pkg/wallet/stores"
	"github.com/sonrhq/core/x/identity/types"
)

// `DIDControllerImpl` is a type that implements the `DIDController` interface.
// @property  - `wallet.Wallet`: This is the interface that the DID controller implements.
// @property  - `store.WalletStore`: This is the interface that the DID controller implements.
type DIDControllerImpl struct {
	store wallet.Store

	ctx context.Context
	aka string

	didDocument    *types.DidDocument
	primaryAccount wallet.Account
	authentication *types.VerificationMethod
}

// `New` creates a new DID controller instance and returns the path to the wallet store.
func New(account wallet.Account, opts ...stores.Option) (DIDController, error) {
	if account == nil {
		return nil, errors.New("account is nil")
	}
	// Create the wallet store.
	st, err := stores.New(account, opts...)
	if err != nil {
		return nil, err
	}
	docc := &DIDControllerImpl{
		ctx:            context.Background(),
		primaryAccount: account,
		store:          st,
	}
	pk, err := crypto.PubKeyFromCommon(account.PubKey())
	if err != nil {
		return nil, err
	}
	addr, err := pk.Bech32(crypto.SONRCoinType.AddrPrefix())
	if err != nil {
		return nil, err
	}
	id := types.NewSonrID(addr)
	doc := types.NewBlankDocument(id)
	vm, err := types.NewPrimaryAccountVM(pk)
	if err != nil {
		return nil, err
	}
	doc.AddAssertion(vm)
	docc.didDocument = doc
	return docc, nil
}

// Address returns the address of the DID controller.
func (d *DIDControllerImpl) Address() string {
	return d.primaryAccount.Address()
}

// ID returns the DID of the DID controller.
func (d *DIDControllerImpl) ID() string {
	return d.primaryAccount.DID()
}

// Document returns the DID document of the DID controller.
func (d *DIDControllerImpl) DidDocument() *types.DidDocument {
	return d.didDocument
}

// Creating a new account.
func (w *DIDControllerImpl) CreateAccount(name string, coinType crypto.CoinType) (*types.VerificationMethod, error) {
	acc, err := w.primaryAccount.Bip32Derive(name, coinType)
	if err != nil {
		return nil, err
	}
	// Set account in list
	err = w.store.PutAccount(acc)
	if err != nil {
		return nil, err
	}

	vm, err := w.didDocument.AddBlockchainAccount(acc.Name(), acc.CoinType(), acc.PubKey())
	if err != nil {
		return nil, err
	}
	return vm, nil
}

// Returning the account.WalletAccount and error.
func (w *DIDControllerImpl) GetAccounts(ct crypto.CoinType) ([]wallet.Account, error) {
	_, err := w.store.ListAccounts()
	if err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("account not found")
}

// ListAccounts returns the list of accounts.
func (w *DIDControllerImpl) ListAccounts() ([]wallet.Account, error) {
	return w.store.ListAccounts()
}

// Path returns the path to the wallet store.
func (w *DIDControllerImpl) Path() string {
	return w.store.Path()
}
