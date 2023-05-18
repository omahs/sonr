package types

import (
	"fmt"
	"strings"

	"github.com/sonrhq/core/internal/crypto"
)

// BlankIdentity returns a blank Identity
func BlankIdentity() *Identity {
	return &Identity{
		Id:                   "",
		Owner:                "",
		PrimaryAlias:         "",
		Authentication:       make([]string, 0),
		AssertionMethod:      make([]string, 0),
		CapabilityDelegation: make([]string, 0),
		CapabilityInvocation: make([]string, 0),
		AlsoKnownAs:          make([]string, 0),
		Metadata:             "",
	}
}

// NewSonrIdentity returns a new Identity with the given owner address and constructs
// the DID from the owner address
func NewSonrIdentity(ownerAddress string) *Identity {
	did := fmt.Sprintf("did:sonr:%s", ownerAddress)
	identity := BlankIdentity()
	identity.Id = did
	identity.Owner = ownerAddress
	return identity
}

// NewWalletIdentity takes an ownerAddress, walletAddress, and CoinType and returns a new Identity
// with the given owner address and constructs the DID from the wallet address
func NewWalletIdentity(ownerAddress, walletAddress string, coinType crypto.CoinType) *Identity {
	did := fmt.Sprintf("did:%s:%s", coinType.DidMethod(), walletAddress)
	identity := BlankIdentity()
	identity.Id = did
	identity.Owner = ownerAddress
	return identity
}

// AddAuthenticationMethod adds a VerificationMethod to the Authentication list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the Authentication list
func (id *Identity) AddAuthenticationMethod(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, auth := range id.Authentication {
		if auth == vm.Id {
			return nil, false
		}
	}
	id.Authentication = append(id.Authentication, vm.Id)
	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               "Authentication",
		VerificationMethod: vm,
		Owner:              id.Owner,
	}
	return vr, true
}

// AddAssertionMethod adds a VerificationMethod to the AssertionMethod list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the AssertionMethod list
func (id *Identity) AddAssertionMethod(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, auth := range id.AssertionMethod {
		if auth == vm.Id {
			return nil, false
		}
	}
	id.AssertionMethod = append(id.AssertionMethod, vm.Id)
	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               "AssertionMethod",
		VerificationMethod: vm,
		Owner:              id.Owner,
	}
	return vr, true
}

// AddCapabilityDelegation adds a VerificationMethod to the CapabilityDelegation list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the CapabilityDelegation list
func (id *Identity) AddCapabilityDelegation(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, auth := range id.CapabilityDelegation {
		if auth == vm.Id {
			return nil, false
		}
	}
	id.CapabilityDelegation = append(id.CapabilityDelegation, vm.Id)
	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               "CapabilityDelegation",
		VerificationMethod: vm,
		Owner:              id.Owner,
	}
	return vr, true
}

// AddCapabilityInvocation adds a VerificationMethod to the CapabilityInvocation list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the CapabilityInvocation list
func (id *Identity) AddCapabilityInvocation(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, auth := range id.CapabilityInvocation {
		if auth == vm.Id {
			return nil, false
		}
	}
	id.CapabilityInvocation = append(id.CapabilityInvocation, vm.Id)
	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               "CapabilityInvocation",
		VerificationMethod: vm,
		Owner:              id.Owner,
	}
	return vr, true
}

// AddKeyAgreement adds a VerificationMethod to the KeyAgreement list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the KeyAgreement list
func (id *Identity) AddKeyAgreement(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, auth := range id.KeyAgreement {
		if auth == vm.Id {
			return nil, false
		}
	}
	id.KeyAgreement = append(id.KeyAgreement, vm.Id)
	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               "KeyAgreement",
		VerificationMethod: vm,
		Owner:              id.Owner,
	}
	return vr, true
}

// SetPrimaryAlias sets the PrimaryAlias of the DID Document to the given alias and appends the alias to the AlsoKnownAs list
// Returns false if the alias is already the AlsoKnownAs list.
func (id *Identity) SetPrimaryAlias(alias string) bool {
	for _, aka := range id.AlsoKnownAs {
		if aka == alias {
			id.PrimaryAlias = alias
			return false
		}
	}
	id.AlsoKnownAs = append(id.AlsoKnownAs, alias)
	id.PrimaryAlias = alias
	return true
}

// ! ||--------------------------------------------------------------------------------||
// ! ||              Primary Identities are DIDDocuments for Sonr Accounts             ||
// ! ||--------------------------------------------------------------------------------||
// NewPrimaryIdentity creates a new DID Document for a primary identity with the given controller and coin type. Returns nil if the controller isnt a sonr account.
func NewPrimaryIdentity(did string, pubKey *crypto.PubKey, cred *VerificationMethod) *DidDocument {
	did, addr := crypto.SONRCoinType.FormatDID(pubKey)
	vm := &VerificationMethod{
		Id:                  did,
		Type:                pubKey.Type(),
		BlockchainAccountId: addr,
	}
	doc := NewBlankDocument(did)
	doc.AssertionMethod = append(doc.AssertionMethod, vm.Id)
	doc.VerificationMethod = append(doc.VerificationMethod, vm)
	if cred != nil {
		doc.VerificationMethod = append(doc.VerificationMethod, cred)
		doc.Authentication = append(doc.Authentication, cred.Id)
	}
	return doc
}

func (d *DidDocument) AddBlockchainIdentity(blockchainIdentity *DidDocument) {
	d.CapabilityDelegation = append(d.CapabilityDelegation, blockchainIdentity.Id)
}

func (d *DidDocument) SetResolvableDomain(resolvableDomain string) {
	d.AlsoKnownAs = append(d.AlsoKnownAs, resolvableDomain)
}

func (d *DidDocument) ListBlockchainIdentities() []string {
	return d.CapabilityDelegation
}

// LinkAdditionalAuthenticationMethod sets the AuthenticationMethod of the DID Document to a PubKey and configured with the given options
func (d *DidDocument) LinkAdditionalAuthenticationMethod(vm *VerificationMethod) *VerificationMethod {
	d.VerificationMethod = append(d.VerificationMethod, vm)
	d.Authentication = append(d.Authentication, vm.Id)
	d.Controller = append(d.Controller, vm.Id)
	return vm
}

// AllowedWebauthnCredentials returns a list of CredentialDescriptors for Webauthn Credentials
func (d *DidDocument) ListCredentialVerificationMethods() []*VerificationMethod {
	allowList := make([]*VerificationMethod, 0)
	credIdList := []string{}
	for _, vm := range d.Authentication {
		credIdList = append(credIdList, vm)
	}

	for _, id := range credIdList {
		vm, _ := d.GetAuthenticationMethod(id)
		allowList = append(allowList, vm)
	}
	return allowList
}

// KnownCredentials returns

// ! ||--------------------------------------------------------------------------------||
// ! ||             Blockchain Identities are intended for Wallet Accounts             ||
// ! ||--------------------------------------------------------------------------------||

// NewBlockchainIdentity creates a new DID Document for a blockchain identity with the given controller and coin type. Returns nil if the controller isnt a sonr account.
func NewBlockchainIdentity(controller string, coinType crypto.CoinType, pubKey *crypto.PubKey) *DidDocument {
	did, addr := coinType.FormatDID(pubKey)
	vm := &VerificationMethod{
		Id:                  did,
		Type:                pubKey.Type(),
		Controller:          controller,
		BlockchainAccountId: addr,
	}
	doc := NewBlankDocument(did)
	doc.Controller = append(doc.Controller, controller)
	doc.VerificationMethod = append(doc.VerificationMethod, vm)
	return doc
}

func ConvertAccAddressToDid(accAddress string) string {
	return strings.ToLower("did:sonr:" + accAddress)
}
