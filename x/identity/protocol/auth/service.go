package auth

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sonrhq/core/pkg/node"
	"github.com/sonrhq/core/pkg/node/config"

	authv1 "github.com/sonrhq/core/x/identity/types/auth/v1"
)

type AuthService struct {
	node   config.IPFSNode
	rpName string
	rpIcon string
	cctx   client.Context
}

// It creates a new VaultService and registers it with the gRPC server
func RegisterVaultService(cctx client.Context, mux *runtime.ServeMux) error {
	node, err := node.NewIPFS(context.Background(), config.WithClientContext(cctx, true))
	if err != nil {
		return err
	}

	authService := &AuthService{
		cctx:   cctx,
		node:   node,
		rpName: "Sonr",
		rpIcon: "https://raw.githubusercontent.com/sonr-hq/sonr/master/docs/static/favicon.png",
	}
	err = authv1.RegisterAuthHandlerServer(context.Background(), mux, authService)
	return nil
}

func (s *AuthService) Authorize(ctx context.Context, req *authv1.AuthorizeRequest) (*authv1.AuthorizeResponse, error) {
	return nil, nil
}

func (s *AuthService) Challenge(ctx context.Context, req *authv1.ChallengeRequest) (*authv1.ChallengeResponse, error) {
	return nil, nil
}
