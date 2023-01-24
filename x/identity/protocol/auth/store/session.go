package store

import (
	"fmt"
	"strings"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
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

// NewSession creates a new session with challenge to be used to register a new account
func NewSession(rpId string, aka string) (*Session, error) {
	s := defaultSession()
	err := s.Apply(requiredOptions(rpId, aka)...)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Apply applies the options to the session
func (s *Session) Apply(opts ...Option) error {
	for _, opt := range opts {
		opt(s)
	}
	if !s.isExisting {
		s.didDoc = types.NewBaseDocument(s.alsoKnownAs, s.ID)
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
	if s.ID == "" {
		return fmt.Errorf("Session ID is empty")
	}
	return nil
}
