// Utility functions for DID Authentication - https://w3c.github.io/did-core/#authentication
// I.e. Verification Material for Webauthn Credentials or KeyPrints. These are used to unlock the Controller Wallet.
package types

import (
	fmt "fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonrhq/core/pkg/common"
)

// AuthenticationCount returns the number of Assertion Methods
func (vm *DidDocument) AuthenticationCount() int {
	return vm.Authentication.Count()
}

// FindAuthenticationMethod finds a VerificationMethod by its ID
func (d *DidDocument) FindAuthenticationMethod(id string) *VerificationMethod {
	return d.Authentication.FindByID(id)
}

// FindAuthenticationMethodByFragment finds a VerificationMethod by its fragment
func (d *DidDocument) FindAuthenticationMethodByFragment(fragment string) *VerificationMethod {
	return d.Authentication.FindByFragment(fragment)[0]
}

// AddAuthenticationMethod adds a VerificationMethod as AuthenticationMethod
// If the controller is not set, it will be set to the document's ID
func (d *DidDocument) AddAuthentication(v *VerificationMethod) {
	if v.Controller == "" {
		v.Controller = d.ID
	}
	d.VerificationMethod.Add(v)
	d.Authentication.Add(v)
}

// AddWebauthnCredential adds a Webauthn Credential as AuthenticationMethod
func (d *DidDocument) AddWebauthnCredential(cred *common.WebauthnCredential) error {
	label := fmt.Sprintf("webauthn-%v", d.AuthenticationCount()+1)
	vm, err := NewWebAuthnVM(cred, WithIDFragmentSuffix(label))
	if err != nil {
		return err
	}
	d.VerificationMethod.Add(vm)
	d.Authentication.Add(vm)
	return nil
}

// AllowedWebauthnCredentials returns a list of CredentialDescriptors for Webauthn Credentials
func (d *DidDocument) AllowedWebauthnCredentials() []protocol.CredentialDescriptor {
	allowList := make([]protocol.CredentialDescriptor, 0)
	creds := d.WebAuthnCredentials()
	for _, cred := range creds {
		allowList = append(allowList, cred.Descriptor())
	}
	return allowList
}
