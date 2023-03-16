package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/x/identity/types"
)

func (k Keeper) SafeSetVerificationRelationships(ctx sdk.Context, doc *types.DidDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RelationshipKeyPrefix))
	for _, vm := range doc.VerificationMethod {
		if !k.HasVerificationRelationship(ctx, vm.Id) {
			vr := vm.ToVerificationRelationship(doc.Controller[0])
			b := k.cdc.MustMarshal(&vr)
			store.Set(types.RelationshipKey(vr.Reference), b)
			ctx.EventManager().EmitEvent(
				sdk.NewEvent("add-verification-relationship", sdk.NewAttribute("did", vm.Id), sdk.NewAttribute("controller", doc.Controller[0]), sdk.NewAttribute("address", doc.Address())),
			)
		}
	}
}

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
