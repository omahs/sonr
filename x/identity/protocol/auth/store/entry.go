package store

import (
	"context"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonrhq/core/pkg/node/config"
)

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

// NewSession creates a new session with challenge to be used to register a new account
func NewSession(rpId string, aka string) (*Session, error) {
	s := defaultSession(rpId, aka)
	err := s.Apply()
	if err != nil {
		return nil, fmt.Errorf("failed to apply options to Webauthn config: %w", err)
	}
	return s, nil
}

// GetSession returns the session for the given username and session ID
func GetSession(node config.IPFSNode, username string, sessionId string) (*Session, error) {
	docs, err := node.LoadDocsStore(username)
	if err != nil {
		return nil, err
	}

	rawVal, err := docs.Get(context.Background(), sessionId, nil)
	if err != nil {
		return nil, err
	}
	return loadSessionFromMap(rawVal[0].(map[string]interface{}))
}
