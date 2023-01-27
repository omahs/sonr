// Main DID Document Constructor Methods
// I.e. Document allows for Reconstruction from Storage of DID Document and Wallet
package types

import (
	fmt "fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BlankDocument creates a blank document to begin the WebAuthnProcess
func BlankDocument(idStr string) *DidDocument {
	return &DidDocument{
		ID:                   idStr,
		Context:              []string{DefaultParams().DidBaseContext, DefaultParams().DidMethodContext},
		Controller:           []string{},
		VerificationMethod:   new(VerificationMethods),
		Authentication:       new(VerificationRelationships),
		AssertionMethod:      new(VerificationRelationships),
		CapabilityInvocation: new(VerificationRelationships),
		CapabilityDelegation: new(VerificationRelationships),
		KeyAgreement:         new(VerificationRelationships),
		Service:              new(Services),
		AlsoKnownAs:          make([]string, 0),
	}
}

// BlankDocument creates a blank document to begin the WebAuthnProcess
func NewBaseDocument(akaStr string) *DidDocument {
	return &DidDocument{
		ID:                   fmt.Sprintf("did:tmp:%s", akaStr),
		Context:              []string{DefaultParams().DidBaseContext, DefaultParams().DidMethodContext},
		Controller:           []string{},
		VerificationMethod:   new(VerificationMethods),
		Authentication:       new(VerificationRelationships),
		AssertionMethod:      new(VerificationRelationships),
		CapabilityInvocation: new(VerificationRelationships),
		CapabilityDelegation: new(VerificationRelationships),
		KeyAgreement:         new(VerificationRelationships),
		Service:              new(Services),
		AlsoKnownAs: []string{
			akaStr,
		},
	}
}

// NewDocumentFromJson creates a new document from a json byte array
func NewDocumentFromJson(b []byte) (*DidDocument, error) {
	var doc DidDocument
	err := doc.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

// Address returns the address of the DID
func (d *DidDocument) Address() string {
	ptrs := strings.Split(d.ID, ":")
	return fmt.Sprintf("%s%s", ptrs[len(ptrs)-2], ptrs[len(ptrs)-1])
}

// AccAddress returns the account address of the DID
func (d *DidDocument) AccAddress() (sdk.AccAddress, error) {
	return ConvertDidToAccAddress(d.ID)
}

// CheckAccAddress checks if the provided sdk.AccAddress or string matches the DID ID
func (d *DidDocument) CheckAccAddress(t interface{}) bool {
	docAccAddr, err := d.AccAddress()
	if err != nil {
		return false
	}

	switch t.(type) {
	case sdk.AccAddress:
		return t.(sdk.AccAddress).Equals(docAccAddr)
	case string:
		addr, err := sdk.AccAddressFromBech32(t.(string))
		if err != nil {
			return false
		}
		return addr.Equals(docAccAddr)
	default:
		return false
	}
}

// SetMetadata sets the metadata of the document
func (vm *DidDocument) SetMetadata(data map[string]string) {
	vm.Metadata = MapToKeyValueList(data)
}
