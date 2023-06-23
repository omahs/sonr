package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	identitytypes "github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/service/types"
)

// This function is a method of the `Keeper` struct and is used to register a new user identity. It takes a context and a `RegisterUserRequest` as input and returns a `RegisterUserResponse` and an error. The function first retrieves the service record associated with the request
// origin and checks if the desired alias is available. It then retrieves the claimable wallet associated with the request UCW ID and verifies the creation challenge using the service's `VerifyCreationChallenge` method. If the challenge is verified, the function assigns an identity
// to the user using the `AssignIdentity` method of the identity keeper and returns the assigned identity and its ID in the response. If any error occurs during the process, the function returns an error.
func (k Keeper) RegisterUser(goCtx context.Context, req *types.RegisterUserRequest) (*types.RegisterUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	service, ok := k.GetServiceRecord(ctx, req.Origin)
	if !ok {
		k.Logger(ctx).Error("(Gateway/service) - error getting service record")
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "service record not found")
	}

	err := k.identityKeeper.CheckAlsoKnownAs(ctx, req.Alias)
	if err != nil {
		k.Logger(ctx).Error("(Gateway/service) - error checking alias")
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "alias already taken")
	}

	credential, err := service.VerifyCreationChallenge(req.Attestation, req.Challenge)
	if err != nil && credential == nil {
		k.Logger(ctx).Debug("(Gateway/service) - error verifying challenge")
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "challenge verification failed")
	}

	// Assign identity to user entity
	account, ks, err := k.vaultKeeper.AssignVault(ctx, req.UcwId, credential)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Identity could not be assigned")
	}
	snr := account[0]
	eth := account[1]
	btc := account[2]

	// Create DID Document and broadcast tx
	credential.Controller = snr.Did()
	authVm := credential.ToVerificationMethod()
	didDoc := identitytypes.NewDIDDocument(snr, authVm, req.Alias)
	didDoc.LinkCapabilityInvocationFromVaultAccount(eth, btc)
	txMsg := identitytypes.NewMsgRegisterIdentity(snr.Address(), didDoc)
	go k.identityKeeper.SignAndBroadcastCosmosTx(snr, txMsg)
	return &types.RegisterUserResponse{
		Did:      didDoc.Id,
		Identity: didDoc,
		Alias:    req.Alias,
		WebauthnCredential: credential,
		VaultKeyshare: 	 ks,
	}, nil
}


// The `AuthenticateUser` function is a method of the `Keeper` struct and is used to authenticate a user. It takes a context and an `AuthenticateUserRequest` as input and returns an `AuthenticateUserResponse` and an error. However, in the given code, the function is not implemented
// and returns an error message indicating that it is not implemented.
func (k Keeper) AuthenticateUser(goCtx context.Context, req *types.AuthenticateUserRequest) (*types.AuthenticateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	service, ok := k.GetServiceRecord(ctx, req.Origin)
	if !ok {
		k.Logger(ctx).Error("(Gateway/service) - error getting service record")
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "service record not found")
	}
	did, ok := k.identityKeeper.GetIdentityByPrimaryAlias(ctx, req.Alias)
	if !ok {
		k.Logger(ctx).Error("(Gateway/service) - error getting identity")
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "identity not found")
	}
	cred, err := service.VerifyAssertionChallenge(req.Assertion, did.ListAuthenticationVerificationMethods()...)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "challenge verification failed")
	}

	acc, err := k.vaultKeeper.UnlockVault(ctx, &did, cred)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "challenge verification failed")
	}
	return &types.AuthenticateUserResponse{
		Did:      acc.Did(),
		Identity: &did,
		Alias:    req.Alias,
	}, nil
}
