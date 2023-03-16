package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/x/identity/types"
)

func (k msgServer) DeletePublicKey(goCtx context.Context, msg *types.MsgDeletePublicKey) (*types.MsgDeletePublicKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeletePublicKeyResponse{}, nil
}
