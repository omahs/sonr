package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/x/identity/types"
)

// SetService set a specific Service in the store from its index
func (k Keeper) SetService(ctx sdk.Context, DomainRecord types.Service) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ServiceKeyPrefix))
	b := k.cdc.MustMarshal(&DomainRecord)
	store.Set(types.ServiceKey(
		DomainRecord.Id,
	), b)
}

// GetDomainRecord returns a DomainRecord from its index
func (k Keeper) GetService(
	ctx sdk.Context,
	id string,
) (val types.Service, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ServiceKeyPrefix))

	b := store.Get(types.ServiceKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDomainRecord removes a DomainRecord from the store
func (k Keeper) RemoveDomainRecord(
	ctx sdk.Context,
	id string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ServiceKeyPrefix))
	store.Delete(types.ServiceKey(
		id,
	))
}

// GetAllServices returns all Services
func (k Keeper) GetAllServices(ctx sdk.Context) (list []types.Service) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ServiceKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Service
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}
