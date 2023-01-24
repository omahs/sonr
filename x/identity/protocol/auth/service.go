package auth

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sonrhq/core/pkg/node/config"

	"github.com/sonrhq/core/x/identity/protocol/auth/store"
	authv1 "github.com/sonrhq/core/x/identity/types/auth/v1"
)

var (
	authService *AuthService
)

type AuthService struct {
	rpName string
	rpIcon string
	cctx   client.Context
	node   config.IPFSNode
}

// It creates a new VaultService and registers it with the gRPC server
func RegisterAuthIPFSService(cctx client.Context, mux *runtime.ServeMux, node config.IPFSNode) error {
	authService = &AuthService{
		cctx:   cctx,
		rpName: "Sonr",
		rpIcon: "https://raw.githubusercontent.com/sonr-hq/sonr/master/docs/static/favicon.png",
	}
	return authv1.RegisterAuthHandlerServer(context.Background(), mux, authService)
}

func (s *AuthService) Challenge(ctx context.Context, req *authv1.ChallengeRequest) (*authv1.ChallengeResponse, error) {
	sess, err := store.NewSession(req.RpId, req.Username)
	if err != nil {
		return nil, err
	}
	return sess.GetChallengeResponse()
}

func (s *AuthService) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	sess, err := store.GetSession(s.node, req.Username, req.SessionId)
	if err != nil {
		return nil, err
	}
	return sess.RegisterCredential(req.RegistrationResponse)
}

func (s *AuthService) Assertion(ctx context.Context, req *authv1.AssertRequest) (*authv1.AssertResponse, error) {
	sess, err := store.NewSession(req.RpId, req.Identifier)
	if err != nil {
		return nil, err
	}
	return sess.GetAssertionOptions()
}

func (s *AuthService) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	sess, err := store.GetSession(s.node, req.Identifier, req.SessionId)
	if err != nil {
		return nil, err
	}
	return sess.AuthorizeCredential(req.AssertionResponse)
}
