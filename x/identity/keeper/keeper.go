package keeper

import (
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sonrhq/core/x/identity/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
		groupKeeper   types.GroupKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper, groupKeeper types.GroupKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}
	k := &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		accountKeeper: accountKeeper, bankKeeper: bankKeeper, groupKeeper: groupKeeper,
	}
	return k
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams()
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                          DIDDocument Keeper Functions                          ||
// ! ||--------------------------------------------------------------------------------||

// CheckAlsoKnownAs checks if an alias is already used
func (k Keeper) CheckAlsoKnownAs(ctx sdk.Context, alias string) error {
	_, found := k.GetDidDocumentByAlsoKnownAs(ctx, alias)
	if found {
		return status.Error(codes.AlreadyExists, "Alias already exists")
	}
	return nil
}

// SetDidDocument set a specific didDocument in the store from its index
func (k Keeper) SetDidDocument(ctx sdk.Context, didDocument types.DidDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrimaryIdentityPrefix))

	ptrs := strings.Split(didDocument.Id, ":")
	addr := ptrs[len(ptrs)-1]
	didDocument.Owner = addr

	b := k.cdc.MustMarshal(&didDocument)
	store.Set(types.DidDocumentKey(
		didDocument.Id,
	), b)
}

// GetDidDocument returns a didDocument from its index
func (k Keeper) GetDidDocument(
	ctx sdk.Context,
	did string,
) (val types.DidDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrimaryIdentityPrefix))
	b := store.Get(types.DidDocumentKey(
		did,
	))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetDidDocumentByAlsoKnownAs returns a didDocument from its index
func (k Keeper) GetDidDocumentByAlsoKnownAs(
	ctx sdk.Context,
	alias string,
) (val types.DidDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrimaryIdentityPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var doc types.DidDocument
		k.cdc.MustUnmarshal(iterator.Value(), &doc)
		if doc.AlsoKnownAs[0] == alias {
			val = doc
			found = true
		}
	}
	return val, found
}

// GetDidDocumentByOwner iterates over all didDocuments and returns the first one that matches the address
func (k Keeper) GetDidDocumentByOwner(
	ctx sdk.Context,
	addr string,
) (val types.DidDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrimaryIdentityPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var doc types.DidDocument
		k.cdc.MustUnmarshal(iterator.Value(), &doc)
		if doc.Owner == addr {
			val = doc
			found = true
		}
	}
	return val, found
}

// GetAllDidDocument returns all didDocument
func (k Keeper) GetAllPrimaryIdentities(ctx sdk.Context) (list []types.DidDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrimaryIdentityPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DidDocument
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                 Relationships - Authentication Keeper Functions                ||
// ! ||--------------------------------------------------------------------------------||
// HasAuthentication checks if the element exists in the store
func (k Keeper) HasAuthentication(ctx sdk.Context, reference string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthenticationKeyPrefix))
	return store.Has(types.RelationshipKey(reference))
}

// SetAuthentication set a specific Service in the store from its index
func (k Keeper) SetAuthentication(ctx sdk.Context, VerificationRelationship types.VerificationRelationship) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthenticationKeyPrefix))
	b := k.cdc.MustMarshal(&VerificationRelationship)
	store.Set(types.RelationshipKey(VerificationRelationship.Reference), b)
}

// GetAuthentication returns a Service from its index
func (k Keeper) GetAuthentication(ctx sdk.Context, reference string) (val types.VerificationRelationship, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthenticationKeyPrefix))

	b := store.Get(types.RelationshipKey(reference))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllAuthentication returns all Relationship
func (k Keeper) GetAllAuthentication(ctx sdk.Context) (list []types.VerificationRelationship) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthenticationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VerificationRelationship
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                   Relationships - Assertion Keeper Functions                   ||
// ! ||--------------------------------------------------------------------------------||
// HasAssertion checks if the element exists in the store
func (k Keeper) HasAssertion(ctx sdk.Context, reference string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssertionKeyPrefix))
	return store.Has(types.RelationshipKey(reference))
}

// SetAssertion set a specific Service in the store from its index
func (k Keeper) SetAssertion(ctx sdk.Context, VerificationRelationship types.VerificationRelationship) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssertionKeyPrefix))
	b := k.cdc.MustMarshal(&VerificationRelationship)
	store.Set(types.RelationshipKey(VerificationRelationship.Reference), b)
}

// GetAssertion returns a Service from its index
func (k Keeper) GetAssertion(ctx sdk.Context, reference string) (val types.VerificationRelationship, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssertionKeyPrefix))

	b := store.Get(types.RelationshipKey(reference))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllAssertion returns all Relationship
func (k Keeper) GetAllAssertion(ctx sdk.Context) (list []types.VerificationRelationship) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssertionKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VerificationRelationship
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

// ! ||--------------------------------------------------------------------------------||
// ! ||             Relationships - Capability Delegation Keeper Functions             ||
// ! ||--------------------------------------------------------------------------------||

// HasCapabilityDelegation checks if the capability delegation relationship exists in the store
func (k Keeper) HasCapabilityDelegation(ctx sdk.Context, reference string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CapabilityDelegationKeyPrefix))
	return store.Has(types.RelationshipKey(reference))
}

// SetCapabilityDelegation sets a specific capability delegation relationship in the store from its reference
func (k Keeper) SetCapabilityDelegation(ctx sdk.Context, delegation types.VerificationRelationship) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CapabilityDelegationKeyPrefix))
	b := k.cdc.MustMarshal(&delegation)
	store.Set(types.RelationshipKey(delegation.Reference), b)
}

// GetCapabilityDelegation returns a capability delegation relationship from its reference
func (k Keeper) GetCapabilityDelegation(ctx sdk.Context, reference string) (delegation types.VerificationRelationship, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CapabilityDelegationKeyPrefix))

	b := store.Get(types.RelationshipKey(reference))
	if b == nil {
		return delegation, false
	}

	k.cdc.MustUnmarshal(b, &delegation)
	return delegation, true
}

// GetAllCapabilityDelegations returns all capability delegation relationships
func (k Keeper) GetAllCapabilityDelegations(ctx sdk.Context) (list []types.VerificationRelationship) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CapabilityDelegationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var delegation types.VerificationRelationship
		k.cdc.MustUnmarshal(iterator.Value(), &delegation)
		list = append(list, delegation)
	}
	return
}

// ! ||--------------------------------------------------------------------------------||
// ! ||             Relationships - Capability Invocation Keeper Functions             ||
// ! ||--------------------------------------------------------------------------------||

// HasCapabilityInvocation checks if the capability invocation relationship exists in the store
func (k Keeper) HasCapabilityInvocation(ctx sdk.Context, reference string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CapabilityInvocationKeyPrefix))
	return store.Has(types.RelationshipKey(reference))
}

// SetCapabilityInvocation sets a specific capability invocation relationship in the store from its reference
func (k Keeper) SetCapabilityInvocation(ctx sdk.Context, invocation types.VerificationRelationship) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CapabilityInvocationKeyPrefix))
	b := k.cdc.MustMarshal(&invocation)
	store.Set(types.RelationshipKey(invocation.Reference), b)
}

// GetCapabilityInvocation returns a capability invocation relationship from its reference
func (k Keeper) GetCapabilityInvocation(ctx sdk.Context, reference string) (invocation types.VerificationRelationship, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CapabilityInvocationKeyPrefix))

	b := store.Get(types.RelationshipKey(reference))
	if b == nil {
		return invocation, false
	}

	k.cdc.MustUnmarshal(b, &invocation)
	return invocation, true
}

// GetAllCapabilityInvocations returns all capability invocation relationships
func (k Keeper) GetAllCapabilityInvocations(ctx sdk.Context) (list []types.VerificationRelationship) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CapabilityInvocationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var invocation types.VerificationRelationship
		k.cdc.MustUnmarshal(iterator.Value(), &invocation)
		list = append(list, invocation)
	}
	return
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                 Relationships - Key Agreement Keeper Functions                 ||
// ! ||--------------------------------------------------------------------------------||
// HasKeyAgreement checks if the key agreement relationship exists in the store
func (k Keeper) HasKeyAgreement(ctx sdk.Context, reference string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeyAgreementKeyPrefix))
	return store.Has(types.RelationshipKey(reference))
}

// SetKeyAgreement sets a specific key agreement relationship in the store from its reference
func (k Keeper) SetKeyAgreement(ctx sdk.Context, agreement types.VerificationRelationship) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeyAgreementKeyPrefix))
	b := k.cdc.MustMarshal(&agreement)
	store.Set(types.RelationshipKey(agreement.Reference), b)
}

// GetKeyAgreement returns a key agreement relationship from its reference
func (k Keeper) GetKeyAgreement(ctx sdk.Context, reference string) (agreement types.VerificationRelationship, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeyAgreementKeyPrefix))

	b := store.Get(types.RelationshipKey(reference))
	if b == nil {
		return agreement, false
	}

	k.cdc.MustUnmarshal(b, &agreement)
	return agreement, true
}

// GetAllKeyAgreements returns all key agreement relationships
func (k Keeper) GetAllKeyAgreements(ctx sdk.Context) (list []types.VerificationRelationship) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeyAgreementKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var agreement types.VerificationRelationship
		k.cdc.MustUnmarshal(iterator.Value(), &agreement)
		list = append(list, agreement)
	}
	return
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                  Wallet Claims                                 ||
// ! ||--------------------------------------------------------------------------------||

// GetClaimableWalletCount get the total number of claimableWallet
func (k Keeper) GetClaimableWalletCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ClaimableWalletCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetClaimableWalletCount set the total number of claimableWallet
func (k Keeper) SetClaimableWalletCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ClaimableWalletCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendClaimableWallet appends a claimableWallet in the store with a new id and update the count
func (k Keeper) AppendClaimableWallet(
	ctx sdk.Context,
	claimableWallet types.ClaimableWallet,
) uint64 {
	// Create the claimableWallet
	count := k.GetClaimableWalletCount(ctx)

	// Set the ID of the appended value
	claimableWallet.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimableWalletKey))
	appendedValue := k.cdc.MustMarshal(&claimableWallet)
	store.Set(GetClaimableWalletIDBytes(claimableWallet.Id), appendedValue)

	// Update claimableWallet count
	k.SetClaimableWalletCount(ctx, count+1)

	return count
}

// SetClaimableWallet set a specific claimableWallet in the store
func (k Keeper) SetClaimableWallet(ctx sdk.Context, claimableWallet types.ClaimableWallet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimableWalletKey))
	b := k.cdc.MustMarshal(&claimableWallet)
	store.Set(GetClaimableWalletIDBytes(claimableWallet.Id), b)
}

// GetClaimableWallet returns a claimableWallet from its id
func (k Keeper) GetClaimableWallet(ctx sdk.Context, id uint64) (val types.ClaimableWallet, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimableWalletKey))
	b := store.Get(GetClaimableWalletIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveClaimableWallet removes a claimableWallet from the store
func (k Keeper) RemoveClaimableWallet(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimableWalletKey))
	store.Delete(GetClaimableWalletIDBytes(id))
}

// GetAllClaimableWallet returns all claimableWallet
func (k Keeper) GetAllClaimableWallet(ctx sdk.Context) (list []types.ClaimableWallet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimableWalletKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClaimableWallet
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetClaimableWalletIDBytes returns the byte representation of the ID
func GetClaimableWalletIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetClaimableWalletIDFromBytes returns ID in uint64 format from a byte array
func GetClaimableWalletIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
