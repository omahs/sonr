package store

import (
	"context"
	"fmt"
	"strings"

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
	ID string

	// Relying Party ID
	RPID     string
	webauthn *webauthn.WebAuthn
	config   *webauthn.Config

	// User Data
	didDoc      *types.DidDocument
	data        webauthn.SessionData
	alsoKnownAs string
	isExisting  bool
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
		s.config.RPDisplayName = name
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
		s.config.RPIcon = icon
	}
}

// WithRPOrigins sets the RPOrigins
func WithRPOrigins(origins []string) Option {
	return func(s *Session) {
		s.config.RPOrigins = origins
	}
}

// WithTimeout sets the Timeout
func WithTimeout(timeout int) Option {
	return func(s *Session) {
		s.config.Timeout = timeout
	}
}

// WithAttestionPreference sets the AttestionPreference
func WithAttestionPreference(pref protocol.ConveyancePreference) Option {
	return func(s *Session) {
		s.config.AttestationPreference = pref
	}
}

// WithAuthenticatorSelect sets the AuthenticatorSelect
func WithAuthenticatorSelect(selectAuth protocol.AuthenticatorSelection) Option {
	return func(s *Session) {
		s.config.AuthenticatorSelection = selectAuth
	}
}

// Apply applies the options to the session
func (s *Session) Apply(opts ...Option) error {
	for _, opt := range opts {
		opt(s)
	}
	if !s.isExisting {
		s.didDoc = types.NewBaseDocument(s.alsoKnownAs, s.ID)
	}
	wauth, err := webauthn.New(s.config)
	if err != nil {
		return err
	}
	s.webauthn = wauth
	return s.Validate()
}

// Sync puts the session for the given username and session ID in Orbit DB
func (s *Session) Sync() error {
	docs, err := node.LoadDocsStore(s.GetUsername())
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
		"_id":         s.ID,
		"RPID":        s.RPID,
		"Config":      s.config,
		"Webauthn":    s.webauthn,
		"DidDoc":      s.didDoc,
		"Data":        s.data,
		"AlsoKnownAs": s.alsoKnownAs,
		"IsExisting":  s.isExisting,
	}
}

// Validating the session.
func (s *Session) Validate() error {
	if s.ID == "" {
		return fmt.Errorf("Session ID is empty")
	}
	if s.didDoc == nil {
		return fmt.Errorf("Session didDoc is nil")
	}
	if !strings.EqualFold(s.alsoKnownAs, s.webauthn.Config.RPDisplayName) {
		return fmt.Errorf("Session RPDisplayName is not equal to alsoKnownAs")
	}
	if s.RPID == "" {
		return fmt.Errorf("Session RPID is empty")
	}
	if s.webauthn == nil {
		return fmt.Errorf("Session webauthn is nil")
	}
	if s.config == nil {
		return fmt.Errorf("Session config is nil")
	}
	return nil
}

// defaultSession returns a default session
func defaultSession() *Session {
	return &Session{
		ID:          uuid.New().String()[:8],
		isExisting:  false,
		alsoKnownAs: defaultRpName,
		config: &webauthn.Config{
			RPID:                   "localhost",
			RPDisplayName:          defaultRpName,
			RPIcon:                 defaultRpIcon,
			RPOrigins:              defaultRpOrigins,
			Timeout:                defaultTimeout,
			AttestationPreference:  defaultAttestationPreference,
			AuthenticatorSelection: defaultAuthSelect,
		},
	}
}

// loadSessionFromMap creates a session from a map
func loadSessionFromMap(m map[string]interface{}) (*Session, error) {
	s := &Session{
		ID:          m["_id"].(string),
		RPID:        m["RPID"].(string),
		webauthn:    m["Webauthn"].(*webauthn.WebAuthn),
		config:      m["Config"].(*webauthn.Config),
		alsoKnownAs: m["AlsoKnownAs"].(string),
		isExisting:  m["IsExisting"].(bool),
		didDoc:      m["DidDoc"].(*types.DidDocument),
		data:        m["Data"].(webauthn.SessionData),
	}
	err := s.Validate()
	if err != nil {
		return nil, err
	}
	return s, nil
}
