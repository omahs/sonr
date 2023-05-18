package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/sonrhq/core/x/service/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) RegisterServiceRecord(goCtx context.Context, msg *types.MsgRegisterServiceRecord) (*types.MsgRegisterServiceRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetServiceRecord(
		ctx,
		msg.Record.Id,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Id already set")
	}
	k.SetServiceRecord(
		ctx,
		*msg.Record,
	)
	return &types.MsgRegisterServiceRecordResponse{}, nil
}

func (k msgServer) UpdateServiceRecord(goCtx context.Context, msg *types.MsgUpdateServiceRecord) (*types.MsgUpdateServiceRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetServiceRecord(
		ctx,
		msg.Record.Id,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Id not set")
	}

	// Checks if the the msg Controller is the same as the current owner
	if msg.Controller != valFound.Controller {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.SetServiceRecord(ctx, *msg.Record)
	return &types.MsgUpdateServiceRecordResponse{}, nil
}

func (k msgServer) BurnServiceRecord(goCtx context.Context, msg *types.MsgBurnServiceRecord) (*types.MsgBurnServiceRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetServiceRecord(
		ctx,
		msg.Id,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Id not set")
	}

	// Checks if the the msg Controller is the same as the current owner
	if msg.Controller != valFound.Controller {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveServiceRecord(
		ctx,
		msg.Id,
	)

	return &types.MsgBurnServiceRecordResponse{}, nil
}

func (k msgServer) RegisterUserEntity(goCtx context.Context, msg *types.MsgRegisterUserEntity) (*types.MsgRegisterUserEntityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRegisterUserEntityResponse{}, nil
}

func (k msgServer) AuthenticateUserEntity(goCtx context.Context, msg *types.MsgAuthenticateUserEntity) (*types.MsgAuthenticateUserEntityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAuthenticateUserEntityResponse{}, nil
}
