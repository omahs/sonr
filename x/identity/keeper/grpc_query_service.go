package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/x/identity/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Service(c context.Context, req *types.QueryGetServiceRequest) (*types.QueryGetServiceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	//Gets `did:snr:addr` from `did:snr:addr#svc`
	val, found := k.GetService(
		ctx,
		req.Origin,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}
	chal, err := val.IssueChallenge()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryGetServiceResponse{Service: val, Challenge: chal.String()}, nil
}

func (k Keeper) ServiceAll(goCtx context.Context, req *types.QueryAllServiceRequest) (*types.QueryAllServiceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryAllServiceResponse{}, nil
}
