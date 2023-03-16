package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/x/identity/types"
)

func (k msgServer) ImportPublicKey(goCtx context.Context, msg *types.MsgImportPublicKey) (*types.MsgImportPublicKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgImportPublicKeyResponse{}, nil
}
