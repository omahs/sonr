package protocol

import (
	"context"

	connect_go "github.com/bufbuild/connect-go"
	"github.com/sonrhq/core/pkg/node"
	v1 "github.com/sonrhq/core/types/highway/v1"
	highwayv1connect "github.com/sonrhq/core/types/highway/v1/highwayv1connect"
)

type Protocol struct {
	Authentication highwayv1connect.AuthenticationHandler
	Accounts 	 highwayv1connect.MpcHandler
	Vault 		 highwayv1connect.VaultHandler
	Node  		 node.IPFS
}

func (p *Protocol) Keygen(ctx context.Context, req  *connect_go.Request[v1.KeygenRequest]) (*connect_go.Response[v1.KeygenResponse], error) {

	return nil, nil
}



func (p *Protocol) LoginStart(context.Context, *connect_go.Request[v1.LoginStartRequest]) (*connect_go.Response[v1.LoginStartResponse], error) {
	return nil, nil
}

func (p *Protocol) LoginFinish(context.Context, *connect_go.Request[v1.LoginFinishRequest]) (*connect_go.Response[v1.LoginFinishResponse], error) {
	return nil, nil
}

func (p *Protocol) CreateAccount(context.Context, *connect_go.Request[v1.CreateAccountRequest]) (*connect_go.Response[v1.CreateAccountResponse], error) {
	return nil, nil
}

func (p *Protocol) GetAccount(context.Context, *connect_go.Request[v1.GetAccountRequest]) (*connect_go.Response[v1.GetAccountResponse], error) {
	return nil, nil
}
