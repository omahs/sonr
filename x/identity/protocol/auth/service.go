package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sonrhq/core/pkg/node/config"

	gocache "github.com/patrickmn/go-cache"
	"github.com/sonrhq/core/x/identity/protocol/auth/session"
	authv1 "github.com/sonrhq/core/x/identity/types/auth/v1"
)

var (
	authService *AuthService
)

// `AuthService` is a struct that contains a reference to a cache, a client context, and an IPFS node.
// @property {string} rpName - The name of the relying party.
// @property {string} rpIcon - The icon of the relying party.
// @property cache - a cache for storing the user's session
// @property cctx - The client context for the service.
// @property node - This is the IPFS node that the service will use to store and retrieve data.
type AuthService struct {
	rpName string
	rpIcon string
	cache  *gocache.Cache
	cctx   client.Context
	node   config.IPFSNode
}

// It creates a new VaultService and registers it with the gRPC server
func RegisterAuthIPFSService(cctx client.Context, mux *runtime.ServeMux, node config.IPFSNode) error {
	authService = &AuthService{
		cctx:   cctx,
		rpName: "Sonr",
		rpIcon: "https://raw.githubusercontent.com/sonrhq/core/master/docs/static/favicon.png",
		node:   node,
		cache:  gocache.New(5*time.Minute, 10*time.Minute),
	}
	return authv1.RegisterAuthHandlerServer(context.Background(), mux, authService)
}

// Creating a new session and storing it in the cache.
func (s *AuthService) Challenge(ctx context.Context, req *authv1.ChallengeRequest) (*authv1.ChallengeResponse, error) {
	sess, err := session.NewSession(req.RpId, req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to create new session: %w \nRequest Was: %s", err, req.String())
	}

	s.cache.Set(req.Username, sess, gocache.DefaultExpiration)
	return sess.GetChallengeResponse()
}

// Registering the credential.
func (s *AuthService) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	sess, err := s.findSession(req.Username)
	if err != nil {
		return nil, err
	}
	return sess.RegisterCredential(req.RegistrationResponse)
}

// Returning the options for the assertion.
func (s *AuthService) Assertion(ctx context.Context, req *authv1.AssertRequest) (*authv1.AssertResponse, error) {
	sess, err := session.NewSession(req.RpId, req.Identifier)
	if err != nil {
		return nil, err
	}
	return sess.GetAssertionOptions()
}

// A gRPC endpoint that is called by the client.
func (s *AuthService) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	sess, err := s.findSession(req.Identifier)
	if err != nil {
		return nil, err
	}
	return sess.AuthorizeCredential(req.AssertionResponse)
}

//
// Helpers
//

// findSession finds a session from the cache
func (s *AuthService) findSession(username string) (*session.Session, error) {
	raw, ok := s.cache.Get(username)
	if !ok {
		return nil, fmt.Errorf("Failed to find session")
	}
	sess, ok := raw.(*session.Session)
	if !ok {
		return nil, fmt.Errorf("Failed to parse session info")
	}
	return sess, nil
}
