package gateway

import (
	"github.com/kataras/go-sessions/v3"
	"github.com/valyala/fasthttp"
)

// Authenticator represents the interface for session-based authentication.
type Authenticator interface {
	// StartSession starts a new session for a user. It returns the session ID.
	StartSession(ctx *fasthttp.RequestCtx, values ...SessionValue) (string, error)

	// EndSession ends the specified session. It returns an error if the session does not exist.
	EndSession(ctx *fasthttp.RequestCtx, sessionID string) error

	// IsValidSessionID checks if the specified session ID is valid (i.e., it corresponds to an active session).
	IsValidSessionID(ctx *fasthttp.RequestCtx, sessionID string) bool

	// GetSession retrieves the session information for a session ID. It returns an error if the session does not exist.
	GetSession(ctx *fasthttp.RequestCtx, sessionID string) (*Session, error)
}

func NewAuthenticator() Authenticator {
	return &authenticator{
		Name:   "sonr",
		Days:   7,
		Secret: "secret",
	}
}

// authenticator implements the Authenticator interface.
type authenticator struct {
	Name   string
	Days   int
	Secret string
}

// StartSession starts a new session for a user. It returns the session ID.
func (a *authenticator) StartSession(ctx *fasthttp.RequestCtx, values ...SessionValue) (string, error) {
	sess := sessions.StartFasthttp(ctx)
	session := defaultSession()
	for _, value := range values {
		value(session)
	}
	session.Save(sess)
	sessionID := sess.ID()
	return sessionID, nil
}

// EndSession ends the specified session. It returns an error if the session does not exist.
func (a *authenticator) EndSession(ctx *fasthttp.RequestCtx, sessionID string) error {
	sess := sessions.StartFasthttp(ctx)
	sess.Destroy()
	return nil
}

// IsValidSessionID checks if the specified session ID is valid (i.e., it corresponds to an active session).
func (a *authenticator) IsValidSessionID(ctx *fasthttp.RequestCtx, sessionID string) bool {
	sess := sessions.StartFasthttp(ctx)
	return sess.ID() == sessionID
}

// GetSession retrieves the session information for a session ID. It returns an error if the session does not exist.
func (a *authenticator) GetSession(ctx *fasthttp.RequestCtx, sessionID string) (*Session, error) {
	sess := sessions.StartFasthttp(ctx)
	return LoadSession(sess), nil
}

