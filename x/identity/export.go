package identity

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonrhq/core/internal/crypto/mpc"
	"github.com/sonrhq/core/internal/vault"
	"github.com/sonrhq/core/x/identity/keeper"
	"github.com/sonrhq/core/x/identity/types"
	servicetypes "github.com/sonrhq/core/x/service/types"
)

// Controller is the identity controller
type Controller = keeper.Controller

// WalletClaims is the wallet claims interface
type WalletClaims = keeper.WalletClaims

// NewController creates a new identity controller
type ControllerOption = keeper.ControllerOption

// NewController creates a new identity controller
func NewController(options ...ControllerOption) (Controller, error) {
	return keeper.NewController(options...)
}

// LoadController loads an identity controller
func LoadController(doc *types.Identity) (Controller, error) {
	return keeper.LoadController(doc)
}

// LoadControllerWithDid loads an identity controller with a did document
func LoadControllerWithDid(doc *types.DidDocument) (Controller, error) {
	return keeper.LoadControllerWithDid(doc)
}

// The function WithUsername sets the username option for a controller.
func WithUsername(username string) ControllerOption {
	return func(o *keeper.ControllerOptions) {
		o.Username = username
	}
}

// The function WithConfigHandlers sets the OnConfigGenerated field of a controller.Options struct to a
// list of handlers.
func WithConfigHandlers(handlers ...mpc.OnConfigGenerated) ControllerOption {
	return func(o *keeper.ControllerOptions) {
		o.OnConfigGenerated = handlers
	}
}

// The function sets a Webauthn credential as an option for a controller.
func WithWebauthnCredential(cred *servicetypes.WebauthnCredential) ControllerOption {
	return func(o *keeper.ControllerOptions) {
		o.WebauthnCredential = cred
	}
}

// The function returns a ControllerOption that disables IPFS in the options of a controller.
func WithIPFSDisabled() ControllerOption {
	return func(o *keeper.ControllerOptions) {
		o.DisableIPFS = true
	}
}

// WithBroadcastTx returns a ControllerOption that enables broadcasting of transactions in the options of a controller.
func WithBroadcastTx() ControllerOption {
	return func(o *keeper.ControllerOptions) {
		o.BroadcastTx = true
	}
}

// LoadClaimableWallet loads a claimable wallet
func LoadClaimableWallet(cw *types.ClaimableWallet) WalletClaims {
	return keeper.LoadClaimableWallet(cw)
}

// StoreCredential stores a credential
func StoreCredential(cred servicetypes.Credential) error {
	return vault.StoreCredential(cred)
}

// FetchCredential fetches a credential
func FetchCredential(keyId string) (servicetypes.Credential, error) {
	return vault.FetchCredential(keyId)
}

// FetchCredentials fetches all credentials from a DidDocument
func FetchCredentials(doc *types.DidDocument) ([]servicetypes.Credential, error) {
	var creds []servicetypes.Credential
	for _, vm := range doc.ListAuthenticationMethods() {
		c, err := vault.FetchCredential(vm.Id)
		if err != nil {
			return nil, err
		}
		creds = append(creds, c)
	}
	return creds, nil
}

// FetchWebauthnCredentialDescriptors fetches all webauthn credential descriptors from a DidDocument
func FetchWebauthnCredentialDescriptors(doc *types.DidDocument) ([]protocol.CredentialDescriptor, error) {
	var creds []servicetypes.Credential
	for _, vm := range doc.ListAuthenticationMethods() {
		c, err := vault.FetchCredential(vm.Id)
		if err != nil {
			return nil, err
		}
		creds = append(creds, c)
	}
	var descriptors []protocol.CredentialDescriptor
	for _, cred := range creds {
		descriptors = append(descriptors, cred.CredentialDescriptor())
	}
	return descriptors, nil
}
