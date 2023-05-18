package keeper

import (
	"crypto/rand"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/vault"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/identity/types/models"
	srvtypes "github.com/sonrhq/core/x/service/types"
)

// ChallengeLength - Length of bytes to generate for a challenge.¡¡
const ChallengeLength = 32

type WalletClaims interface {
	GetClaimableWallet() *types.ClaimableWallet
	IssueChallenge() (protocol.URLEncodedBase64, error)
	Assign(cred *srvtypes.WebauthnCredential, alias string) (Controller, error)
	Address() string
}

type walletClaims struct {
	Claims  *types.ClaimableWallet `json:"claims" yaml:"claims"`
	Creator string                 `json:"creator" yaml:"creator"`
}

// The function creates a new wallet claim with a given creator and key shares.
func NewWalletClaims(creator string, kss []models.KeyShare) (WalletClaims, error) {
	pub := kss[0].PubKey()
	keyIds := make([]string, 0)
	for _, ks := range kss {
		keyIds = append(keyIds, ks.Did())
	}

	cw := &types.ClaimableWallet{
		Creator:   creator,
		PublicKey: pub.Base64(),
		Keyshares: keyIds,
		Count:     int32(len(kss)),
		Claimed:   false,
	}

	return &walletClaims{
		Claims:  cw,
		Creator: creator,
	}, nil
}

// The function returns a WalletClaims interface that contains the claimable wallet and its creator.
func LoadClaimableWallet(cw *types.ClaimableWallet) WalletClaims {
	return &walletClaims{
		Claims:  cw,
		Creator: cw.Creator,
	}
}

// The function returns the address of the claimable wallet.
func (wc *walletClaims) Address() string {
	ptrs := strings.Split(wc.Claims.Keyshares[0], "did:sonr:")
	addr := strings.Split(ptrs[1], "#")[0]
	return addr
}

// The `GetClaimableWallet()` function is a method of the `walletClaims` struct that returns a pointer
// to the `ClaimableWallet` object stored in the struct. This allows other parts of the code to access
// the `ClaimableWallet` object and its properties.
func (wc *walletClaims) GetClaimableWallet() *types.ClaimableWallet {
	return wc.Claims
}

// This function is used to issue a challenge for the claimable wallet. It returns the public key of
// the claimable wallet as a URL-encoded base64 string, which can be used as a challenge for WebAuthn
// authentication. If the public key is empty, it returns an error.
func (wc *walletClaims) IssueChallenge() (protocol.URLEncodedBase64, error) {
	if wc.Claims.PublicKey == "" {
		return nil, fmt.Errorf("public key is empty")
	}
	return CreateChallenge()
}

// This function assigns a WebAuthn credential to a claimable wallet by creating a new DID document and
// adding the credential as an additional authentication method. It then creates a new `didController`
// instance with the new DID document and returns it as a `Controller` interface. The `alias` parameter
// is used to set the `AlsoKnownAs` field in the DID document.
func (wc *walletClaims) Assign(cred *srvtypes.WebauthnCredential, alias string) (Controller, error) {
	kss := make([]models.KeyShare, 0)
	for _, ks := range wc.Claims.Keyshares {
		ks, err := vault.GetKeyshare(ks)
		if err != nil {
			return nil, fmt.Errorf("error getting keyshare: %w", err)
		}
		kss = append(kss, ks)
	}

	acc := models.NewAccount(kss, crypto.SONRCoinType)
	err := vault.InsertAccount(acc)
	if err != nil {
		return nil, fmt.Errorf("error inserting account: %w", err)
	}
	cred.Controller = acc.Did()
	id, snrvr, _ := acc.GetIdentity(wc.Address())
	cn := &didController{
		primary:        acc,
		identity:       id,
		blockchain:     []models.Account{},
		disableIPFS:    false,
		currCredential: cred,
	}
	cn.RegisterIdentity(id, alias, uint32(wc.Claims.Id), snrvr)
	return cn, nil
}

// CreateChallenge creates a new challenge that should be signed and returned by the authenticator. The spec recommends
// using at least 16 bytes with 100 bits of entropy. We use 32 bytes.
func CreateChallenge() (challenge protocol.URLEncodedBase64, err error) {
	challenge = make([]byte, ChallengeLength)

	if _, err = rand.Read(challenge); err != nil {
		return nil, err
	}

	return challenge, nil
}

func (k Keeper) NextUnclaimedWallet(ctx sdk.Context) (*types.ClaimableWallet, protocol.URLEncodedBase64, error) {
	// Make sure more than zero unclaimed wallets exist
	if k.GetClaimableWalletCount(ctx) == 0 {
		return nil, nil, fmt.Errorf("no unclaimed wallets exist")
	}

	// Get the next unclaimed wallet
	ucws := k.GetAllClaimableWallet(ctx)
	ucw := ucws[0]
	chal, err := CreateChallenge()
	if err != nil {
		return nil, nil, fmt.Errorf("error creating challenge: %w", err)
	}
	return &ucw, chal, nil
}

func (k Keeper) AssignIdentity(ctx sdk.Context, ucw types.ClaimableWallet, cred *srvtypes.WebauthnCredential, alias string) (*types.Identity, error) {
	// Get the keyshares for the claimable wallet
	kss := make([]models.KeyShare, 0)
	for _, ks := range ucw.Keyshares {
		ks, err := vault.GetKeyshare(ks)
		if err != nil {
			return nil, fmt.Errorf("error getting keyshare: %w", err)
		}
		kss = append(kss, ks)
	}

	// Create a new account with the keyshares
	acc := models.NewAccount(kss, crypto.SONRCoinType)
	err := vault.InsertAccount(acc)
	if err != nil {
		return nil, fmt.Errorf("error inserting account: %w", err)
	}

	// Create a new DID document with the account
	cred.Controller = acc.Did()
	id, snrvr, _ := acc.GetIdentity(ucw.Address())
	cn := &didController{
		primary:        acc,
		identity:       id,
		blockchain:     []models.Account{},
		disableIPFS:    false,
		currCredential: cred,
	}
	cn.RegisterIdentity(id, alias, uint32(ucw.Id), snrvr)
	return id, nil
}
