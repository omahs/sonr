package store

import (
	"context"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
	"github.com/sonrhq/core/x/identity/types"
)

// `Session` is a struct that contains a `string` (`ID`), a `string` (`RPID`), a
// `common.WebauthnCredential` (`WebauthnCredential`), a `types.DidDocument` (`DidDoc`), a
// `webauthn.SessionData` (`Data`), and a `string` (`AlsoKnownAs`).
// @property {string} ID - The session ID.
// @property {string} RPID - The Relying Party ID. This is the domain of the relying party.
// @property WebauthnCredential - This is the credential that was created by the user.
// @property DidDoc - The DID Document of the user.
// @property Data - This is the data that is returned from the webauthn.Create() function.
// @property {string} AlsoKnownAs - The user's username.
type Session struct {
	// Session ID
	ID   string
	rpid string
	aka  string

	// Relying Party ID
	webauthn *webauthn.WebAuthn

	// User Data
	didDoc     *types.DidDocument
	data       *webauthn.SessionData
	isExisting bool
}

// Option is a function that configures a session
type Option func(*webauthn.Config)

// WithRPIcon sets the RPIcon
func WithRPIcon(icon string) Option {
	return func(s *webauthn.Config) {
		s.RPIcon = icon
	}
}

// WithRPOrigins sets the RPOrigins
func WithRPOrigins(origins []string) Option {
	return func(s *webauthn.Config) {
		s.RPOrigins = origins
	}
}

// WithTimeout sets the Timeout
func WithTimeout(timeout int) Option {
	return func(s *webauthn.Config) {
		s.Timeout = timeout
	}
}

// WithAttestionPreference sets the AttestionPreference
func WithAttestionPreference(pref protocol.ConveyancePreference) Option {
	return func(s *webauthn.Config) {
		s.AttestationPreference = pref
	}
}

// WithAuthenticatorSelect sets the AuthenticatorSelect
func WithAuthenticatorSelect(selectAuth protocol.AuthenticatorSelection) Option {
	return func(s *webauthn.Config) {
		s.AuthenticatorSelection = selectAuth
	}
}

// Apply applies the options to the session
func (s *Session) Apply(opts ...Option) error {
	c := &webauthn.Config{
		RPID:                   s.rpid,
		RPDisplayName:          s.aka,
		RPIcon:                 defaultRpIcon,
		RPOrigins:              defaultRpOrigins,
		Timeout:                defaultTimeout,
		AttestationPreference:  defaultAttestationPreference,
		AuthenticatorSelection: defaultAuthSelect,
	}
	for _, opt := range opts {
		opt(c)
	}
	wauth, err := webauthn.New(c)
	if err != nil {
		return err
	}
	s.webauthn = wauth
	return nil
}

// Sync puts the session for the given username and session ID in Orbit DB
func (s *Session) Sync() error {
	docs, err := node.LoadDocsStore(s.aka)
	if err != nil {
		return err
	}
	_, err = docs.Put(context.Background(), s.ToMap())
	if err != nil {
		return err
	}
	return nil
}

// ToMap converts the session to a map
func (s *Session) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"_id":      s.ID,
		"Webauthn": s.webauthn,
		"DidDoc":   s.didDoc,
		"Data":     s.data,
	}
}

// defaultSession returns a default session
func defaultSession(rpid string, aka string) *Session {
	id := uuid.New().String()[:8]
	return &Session{
		ID:         id,
		isExisting: false,
		didDoc:     types.NewBaseDocument(aka, id),
		rpid:       rpid,
		aka:        aka,
	}
}

// loadSessionFromMap creates a session from a map
func loadSessionFromMap(m map[string]interface{}) (*Session, error) {
	s := &Session{
		ID:       m["_id"].(string),
		webauthn: m["Webauthn"].(*webauthn.WebAuthn),
		didDoc:   m["DidDoc"].(*types.DidDocument),
		data:     m["Data"].(*webauthn.SessionData),
	}
	return s, nil
}
