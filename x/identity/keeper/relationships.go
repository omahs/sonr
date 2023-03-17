package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/x/identity/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HasVerificationRelationship checks if the element exists in the store
func (k Keeper) HasVerificationRelationship(ctx sdk.Context, reference string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RelationshipKeyPrefix))
	return store.Has(types.RelationshipKey(reference))
}

// SetVerificationRelationship set a specific Service in the store from its index
func (k Keeper) SetVerificationRelationship(ctx sdk.Context, VerificationRelationship types.VerificationRelationship) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RelationshipKeyPrefix))
	b := k.cdc.MustMarshal(&VerificationRelationship)
	store.Set(types.RelationshipKey(VerificationRelationship.Reference), b)
}

// GetVerificationRelationship returns a Service from its index
func (k Keeper) GetVerificationRelationship(ctx sdk.Context, reference string) (val types.VerificationRelationship, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RelationshipKeyPrefix))

	b := store.Get(types.RelationshipKey(reference))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllVerificationRelationships returns all Relationship
func (k Keeper) GetAllVerificationRelationships(ctx sdk.Context) (list []types.VerificationRelationship) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RelationshipKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VerificationRelationship
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) ResolveDidDocument(ctx sdk.Context, doc types.DidDocument) (types.ResolvedDidDocument, error) {
	resolvedDidDocument := doc.ToResolved()

	vrs := []types.VerificationRelationship{}
	for _, relationship := range doc.VerificationMethod {
		vr, ok := k.GetVerificationRelationship(ctx, relationship.Id)
		if !ok {
			return types.ResolvedDidDocument{}, status.Error(codes.NotFound, fmt.Sprintf("verification relationship %s not found", relationship.Id))
		}
		vrs = append(vrs, vr)
	}

	resolvedDidDocument.AddVerificationRelationship(vrs)
	return *resolvedDidDocument, nil
}

func (k Keeper) GetVerificationRelationshipsFromList(ctx sdk.Context, addrs ...string) ([]types.VerificationRelationship, error) {
	vrs := make([]types.VerificationRelationship, 0, len(addrs))

	for _, addr := range addrs {
		if vr, found := k.GetVerificationRelationship(sdk.UnwrapSDKContext(ctx), addr); found {
			vrs = append(vrs, vr)
		} else {
			return nil, status.Error(codes.NotFound, "not found")
		}
	}

	return vrs, nil
}
