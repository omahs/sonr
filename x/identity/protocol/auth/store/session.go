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

	// Webauthn Config
	RPID                string
	RPDisplayName       string
	RPIcon              string
	RPOrigins           []string
	Timeout             int
	AttestionPreference protocol.ConveyancePreference
	AuthenticatorSelect protocol.AuthenticatorSelection

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

// Apply applies the options to the session
func (s *Session) Apply(opts ...Option) error {
	for _, opt := range opts {
		opt(s)
	}
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
		"_id":                 s.ID,
		"RPID":                s.RPID,
		"RPDisplayName":       s.RPDisplayName,
		"RPIcon":              s.RPIcon,
		"RPOrigins":           s.RPOrigins,
		"Timeout":             s.Timeout,
		"AttestionPreference": s.AttestionPreference,
		"AuthenticatorSelect": s.AuthenticatorSelect,
		"DidDoc":              s.didDoc,
		"Data":                s.data,
		"AlsoKnownAs":         s.alsoKnownAs,
		"IsExisting":          s.isExisting,
	}
}

// Validating the session.
func (s *Session) Validate() error {
	if s.ID == "" {
		return fmt.Errorf("Session ID is empty")
	}
	if s.isExisting && s.didDoc == nil {
		return fmt.Errorf("Session is existing but didDoc is nil")
	}
	if !strings.EqualFold(s.alsoKnownAs, s.RPDisplayName) {
		return fmt.Errorf("Session RPDisplayName is not equal to alsoKnownAs")
	}
	if s.RPID == "" {
		return fmt.Errorf("Session RPID is empty")
	}
	if s.RPDisplayName == "" {
		return fmt.Errorf("Session RPDisplayName is empty")
	}
	if s.RPIcon == "" {
		return fmt.Errorf("Session RPIcon is empty")
	}
	if len(s.RPOrigins) == 0 {
		return fmt.Errorf("Session RPOrigins is empty")
	}
	if s.Timeout == 0 {
		return fmt.Errorf("Session Timeout is empty")
	}
	if s.AttestionPreference == "" {
		return fmt.Errorf("Session AttestionPreference is empty")
	}
	if s.AuthenticatorSelect.AuthenticatorAttachment == "" {
		return fmt.Errorf("Session AuthenticatorSelect is empty")
	}
	return nil
}

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

// loadSessionFromMap creates a session from a map
func loadSessionFromMap(m map[string]interface{}) (*Session, error) {
	s := &Session{
		ID:                  m["_id"].(string),
		RPID:                m["RPID"].(string),
		RPDisplayName:       m["RPDisplayName"].(string),
		RPIcon:              m["RPIcon"].(string),
		RPOrigins:           m["RPOrigins"].([]string),
		Timeout:             m["Timeout"].(int),
		AttestionPreference: m["AttestionPreference"].(protocol.ConveyancePreference),
		AuthenticatorSelect: m["AuthenticatorSelect"].(protocol.AuthenticatorSelection),
		alsoKnownAs:         m["AlsoKnownAs"].(string),
		isExisting:          m["IsExisting"].(bool),
		didDoc:              m["DidDoc"].(*types.DidDocument),
		data:                m["Data"].(webauthn.SessionData),
	}
	err := s.Validate()
	if err != nil {
		return nil, err
	}
	return s, nil
}

// requiredOptions returns configured options
func requiredOptions(rpid, aka string) []Option {
	return []Option{
		WithRPID(rpid),
		WithAlsoKnownAs(aka),
	}
}
