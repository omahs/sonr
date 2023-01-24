package store

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/google/uuid"
	"github.com/sonrhq/core/x/identity/types"
)

// Default Variables
var (
	// Default Origins
	defaultRpOrigins = []string{
		"https://auth.sonr.io",
		"https://sonr.id",
		"https://sandbox.sonr.network",
		"http://localhost:3000",
	}

	// Default Icon to display
	defaultRpIcon = "https://raw.githubusercontent.com/sonr-hq/sonr/master/docs/static/favicon.png"

	// Default name to display
	defaultRpName = "Sonr"

	// defaultAttestionPreference
	defaultAttestationPreference = protocol.PreferDirectAttestation

	// defaultAuthSelect
	defaultAuthSelect = protocol.AuthenticatorSelection{
		AuthenticatorAttachment: protocol.AuthenticatorAttachment("platform"),
	}

	// defaultTimeout
	defaultTimeout = 60000
)

// defaultSession returns a default session
func defaultSession() *Session {
	return &Session{
		ID:                  uuid.New().String()[:8],
		RPDisplayName:       defaultRpName,
		RPIcon:              defaultRpIcon,
		RPOrigins:           defaultRpOrigins,
		Timeout:             defaultTimeout,
		AttestionPreference: defaultAttestationPreference,
		AuthenticatorSelect: defaultAuthSelect,
		isExisting:          false,
	}
}

// requiredOptions returns configured options
func requiredOptions(rpid, aka string) []Option {
	return []Option{
		WithRPID(rpid),
		WithAlsoKnownAs(aka),
	}
}

// Option is a function that configures a session
type Option func(*Session)

// WithRPID sets the RPID
func WithRPID(id string) Option {
	return func(s *Session) {
		s.RPID = id
	}
}

// WithAlsoKnownAs sets the RPDisplayName
func WithAlsoKnownAs(name string) Option {
	return func(s *Session) {
		s.RPDisplayName = name
		s.alsoKnownAs = name
	}
}

// WithDIDDoc sets the DIDDoc
func WithDIDDoc(doc *types.DidDocument) Option {
	return func(s *Session) {
		s.didDoc = doc
		s.isExisting = true
	}
}

// WithRPIcon sets the RPIcon
func WithRPIcon(icon string) Option {
	return func(s *Session) {
		s.RPIcon = icon
	}
}

// WithRPOrigins sets the RPOrigins
func WithRPOrigins(origins []string) Option {
	return func(s *Session) {
		s.RPOrigins = origins
	}
}

// WithTimeout sets the Timeout
func WithTimeout(timeout int) Option {
	return func(s *Session) {
		s.Timeout = timeout
	}
}

// WithAttestionPreference sets the AttestionPreference
func WithAttestionPreference(pref protocol.ConveyancePreference) Option {
	return func(s *Session) {
		s.AttestionPreference = pref
	}
}

// WithAuthenticatorSelect sets the AuthenticatorSelect
func WithAuthenticatorSelect(selectAuth protocol.AuthenticatorSelection) Option {
	return func(s *Session) {
		s.AuthenticatorSelect = selectAuth
	}
}
