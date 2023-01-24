package vault

import (
	"context"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sonrhq/core/pkg/node/config"
	"github.com/sonrhq/core/x/identity/protocol/vault/store"
	v1 "github.com/sonrhq/core/x/identity/types/vault/v1"
)

// Default Variables
var (
	defaultRpOrigins = []string{
		"https://auth.sonr.io",
		"https://sonr.id",
		"https://sandbox.sonr.network",
		"localhost:3000",
	}
	vaultService *VaultService
)

// `VaultService` is a type that implements the `v1.VaultServer` interface, and has a field called
// `highway` of type `*HighwayNode`.
// @property  - `v1.VaultServer`: This is the interface that the Vault service implements.
// @property highway - This is the HighwayNode that the VaultService is running on.
type VaultService struct {
	bank   *store.VaultBank
	node   config.IPFSNode
	rpName string
	rpIcon string
	cctx   client.Context
	cache  *gocache.Cache
}

// It creates a new VaultService and registers it with the gRPC server
func RegisterVaultIPFSService(cctx client.Context, mux *runtime.ServeMux, node config.IPFSNode) error {

	cache := gocache.New(time.Minute*2, time.Minute*10)
	vaultService = &VaultService{
		cctx:   configureClientCtx(cctx),
		bank:   store.InitBank(cctx, node, cache),
		node:   node,
		rpName: "Sonr",
		rpIcon: "https://raw.githubusercontent.com/sonr-hq/sonr/master/docs/static/favicon.png",
		cache:  cache,
	}
	return v1.RegisterVaultHandlerServer(context.Background(), mux, vaultService)
}

// Register registers a new keypair and returns the public key.
func (v *VaultService) NewWallet(ctx context.Context, req *v1.NewWalletRequest) (*v1.NewWalletResponse, error) {
	// Get Session
	didDoc, wallet, err := v.bank.FinishRegistration(req.SessionId, req.CredentialResponse)
	if err != nil {
		return nil, err
	}

	// Return response
	return &v1.NewWalletResponse{
		Success:     true,
		Address:     wallet.WalletConfig().Address,
		DidDocument: didDoc,
	}, nil
}

// CreateAccount derives a new key from the private key and returns the public key.
func (v *VaultService) CreateAccount(ctx context.Context, req *v1.CreateAccountRequest) (*v1.CreateAccountResponse, error) {
	return nil, fmt.Errorf("Method is unimplemented")
}

// ListAccounts lists all the accounts derived from the private key.
func (v *VaultService) ListAccounts(ctx context.Context, req *v1.ListAccountsRequest) (*v1.ListAccountsResponse, error) {
	return nil, fmt.Errorf("Method is unimplemented")
}

// DeleteAccount deletes the account with the given address.
func (v *VaultService) DeleteAccount(ctx context.Context, req *v1.DeleteAccountRequest) (*v1.DeleteAccountResponse, error) {
	return nil, fmt.Errorf("Method is unimplemented")
}

// Refresh refreshes the keypair and returns the public key.
func (v *VaultService) Refresh(ctx context.Context, req *v1.RefreshRequest) (*v1.RefreshResponse, error) {
	return nil, fmt.Errorf("Method is unimplemented")
}

// Sign signs the data with the private key and returns the signature.
func (v *VaultService) SignTransaction(ctx context.Context, req *v1.SignTransactionRequest) (*v1.SignTransactionResponse, error) {
	return nil, fmt.Errorf("Method is unimplemented")
}

//
// Helper Functions
//

// configures the client context with the given parameters
func configureClientCtx(cctx client.Context) client.Context {
	ctx := cctx.WithFromName("alice")
	acc := ctx.GetFromAddress()
	ctx = ctx.WithFeePayerAddress(acc)
	ctx = ctx.WithFromAddress(acc)
	return ctx
}
