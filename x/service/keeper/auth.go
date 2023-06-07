package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/local"
	identitytypes "github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/service/types"
	"github.com/sonrhq/core/x/vault"
)

// This function is a method of the `Keeper` struct and is used to register a new user identity. It takes a context and a `RegisterUserRequest` as input and returns a `RegisterUserResponse` and an error. The function first retrieves the service record associated with the request
// origin and checks if the desired alias is available. It then retrieves the claimable wallet associated with the request UCW ID and verifies the creation challenge using the service's `VerifyCreationChallenge` method. If the challenge is verified, the function assigns an identity
// to the user using the `AssignIdentity` method of the identity keeper and returns the assigned identity and its ID in the response. If any error occurs during the process, the function returns an error.
func (k Keeper) RegisterUser(goCtx context.Context, req *types.RegisterUserRequest) (*types.RegisterUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	didDocChan := make(chan *identitytypes.DIDDocument)
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
	account, err := k.vaultKeeper.AssignVault(ctx, req.UcwId, credential)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Identity could not be assigned")
	}

	// Start a go routine to create the identity
	go func(acc vault.Account) {
		// Create btc, eth default accounts
		btcAcc, err := acc.DeriveAccount(crypto.BTCCoinType, 0, "BTC#1")
		if err != nil {
			k.Logger(ctx).Error("(Gateway/service) - error deriving BTC account", err)
			return
		}

		ethAcc, err := acc.DeriveAccount(crypto.ETHCoinType, 0, "ETH#1")
		if err != nil {
			k.Logger(ctx).Error("(Gateway/service) - error deriving ETH account", err)
			return
		}

		// Create DID Document
		didDoc, err := k.identityKeeper.AssignIdentity(credential.ToVerificationMethod(), account, req.Alias, btcAcc, ethAcc)
		if err != nil {
			k.Logger(ctx).Error("(Gateway/service) - error assigning identity", err)
			return
		}
		didDocChan <- didDoc

		// Sign and broadcast identity registration message
		bz, err := acc.SignCosmosTx(identitytypes.NewMsgRegisterIdentity(acc.Address(), didDoc))
		if err != nil {
			k.Logger(ctx).Error("(Gateway/service) - error signing identity registration message", err)
			return
		}
		_, err = local.Context().BroadcastTx(bz)
		if err != nil {
			k.Logger(ctx).Error("(Gateway/service) - error broadcasting identity registration message", err)
			return
		}
	}(account)

	// Return the identity
	didDoc := <-didDocChan
	return &types.RegisterUserResponse{
		Did:      didDoc.Id,
		Identity: didDoc,
		Alias:    req.Alias,
		Jwt:      "",
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
	did, err := k.identityKeeper.ResolveIdentityByPrimaryAlias(ctx, req.Alias)
	if err != nil {
		k.Logger(ctx).Error("(Gateway/service) - error getting identity")
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "identity not found")
	}
	err = service.VerifyAssertionChallenge(req.Assertion, did.ListAuthenticationVerificationMethods()...)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "challenge verification failed")
	}
	return &types.AuthenticateUserResponse{
		Did:      did.Id,
		Identity: &did,
		Alias:    req.Alias,
		Jwt:      "",
	}, nil
}
