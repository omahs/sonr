package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	idtypes "github.com/sonrhq/core/x/identity/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// NewWebauthnCreationOptions returns the webauthn creation options.
func (p Params) NewWebauthnCreationOptions(s *ServiceRecord, uuid string, challenge protocol.URLEncodedBase64) (protocol.CredentialCreation, error) {
	// Build the credential creation options.
	opts := protocol.PublicKeyCredentialCreationOptions{
		// Generated Challenge.
		Challenge: challenge,

		// Service resulting properties.
		User: s.GetUserEntity(uuid),

		// Preconfigured parameters.
		Parameters: []protocol.CredentialParameter{
			{
				Type:      protocol.PublicKeyCredentialType,
				Algorithm: webauthncose.AlgES256,
			},
		},
		RelyingParty: protocol.RelyingPartyEntity{
			CredentialEntity: protocol.CredentialEntity{
				Name: s.Name,
			},
			ID: s.Origin,
		},
		Timeout: int(60000),
		AuthenticatorSelection: protocol.AuthenticatorSelection{
			AuthenticatorAttachment: protocol.Platform,
			ResidentKey:             protocol.ResidentKeyRequirementRequired,
			UserVerification:        protocol.VerificationRequired,
		},
		Attestation: protocol.PreferDirectAttestation,
	}
	return protocol.CredentialCreation{Response: opts}, nil
}

// NewWebauthnAssertionOptions returns the webauthn assertion options.
func (p Params) NewWebauthnAssertionOptions(s *ServiceRecord, challenge protocol.URLEncodedBase64, auth []*idtypes.VerificationMethod) (protocol.CredentialAssertion, error) {
	allowedCreds := []protocol.CredentialDescriptor{}
	for _, vm := range auth {
		allowedCreds = append(allowedCreds, protocol.CredentialDescriptor{
			CredentialID: vm.WebauthnCredentialID(),
			Type:         "public-key",
		})
	}

	// Build the credential assertion options.
	opts := protocol.PublicKeyCredentialRequestOptions{
		// Generated Challenge.
		Challenge:        challenge,
		RelyingPartyID:   s.Origin,
		UserVerification: protocol.VerificationRequired,

		// Preconfigured parameters.
		Timeout: int(60000),
		AllowedCredentials: allowedCreds,
	}
	return protocol.CredentialAssertion{Response: opts}, nil
}
