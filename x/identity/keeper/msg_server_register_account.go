package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/pkg/wallet"
	"github.com/sonrhq/core/x/identity/types"
)

func (k msgServer) RegisterAccount(goCtx context.Context, msg *types.MsgRegisterAccount) (*types.MsgRegisterAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetService(ctx, msg.Origin)
	if found {
		return nil, fmt.Errorf("service already exists")
	}

	cred, err := val.VerifyCreationChallenge(msg.CredentialResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to verify challenge: %w", err)
	}

	wallChan := make(chan wallet.Wallet)
	errChan := make(chan error)
	go func() {
		wall, err := wallet.NewWallet(msg.Uuid, 1)
		if err != nil {
			errChan <- err
			return
		}
		wallChan <- wall
	}()

	select {
	case wall := <-wallChan:
		doc, err := wall.Assign(cred)
		if err != nil {
			return nil, fmt.Errorf("failed to set authentication: %w", err)
		}
		k.SetDidDocument(ctx, *doc)
		resolved, err := k.ResolveDidDocument(ctx, *doc)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve did document: %w", err)
		}
		return &types.MsgRegisterAccountResponse{
			Did:      doc.Id,
			Document: &resolved,
		}, nil
	case err := <-errChan:
		return nil, fmt.Errorf("failed to create wallet: %w", err)
	}
}
