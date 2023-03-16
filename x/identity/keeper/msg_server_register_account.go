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
		doc, err := wall.SetAuthentication(cred)
		if err != nil {
			return nil, fmt.Errorf("failed to set authentication: %w", err)
		}
		k.SetDidDocument(ctx, *doc)
	case err := <-errChan:
		fmt.Println(err)
		break
	}

	return &types.MsgRegisterAccountResponse{}, nil
}
