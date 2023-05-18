package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonrhq/core/x/service/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ListServiceRecords(goCtx context.Context, req *types.ListServiceRecordsRequest) (*types.ListServiceRecordsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var serviceRecords []types.ServiceRecord
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	serviceRecordStore := prefix.NewStore(store, types.KeyPrefix(types.ServiceRecordKeyPrefix))

	pageRes, err := query.Paginate(serviceRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var serviceRecord types.ServiceRecord
		if err := k.cdc.Unmarshal(value, &serviceRecord); err != nil {
			return err
		}

		serviceRecords = append(serviceRecords, serviceRecord)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.ListServiceRecordsResponse{ServiceRecord: serviceRecords, Pagination: pageRes}, nil
}

func (k Keeper) ServiceRecord(goCtx context.Context, req *types.QueryServiceRecordRequest) (*types.QueryServiceRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetServiceRecord(
		ctx,
		req.Origin,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryServiceRecordResponse{ServiceRecord: val}, nil
}

func (k Keeper) ServiceRelationship(goCtx context.Context, req *types.QueryGetServiceRelationshipRequest) (*types.QueryGetServiceRelationshipResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	serviceRelationships, found := k.GetServiceRelationship(ctx, req.Origin)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetServiceRelationshipResponse{ServiceRelationships: serviceRelationships}, nil
}

func (k Keeper) ServiceAttestion(goCtx context.Context, req *types.QueryGetServiceAttestionRequest) (*types.QueryGetServiceAttestionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryGetServiceAttestionResponse{}, nil
}

func (k Keeper) ServiceAssertion(goCtx context.Context, req *types.QueryGetServiceAssertionRequest) (*types.QueryGetServiceAssertionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryGetServiceAssertionResponse{}, nil
}
