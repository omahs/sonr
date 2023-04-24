package controller

import (
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/x/identity/keeper"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/identity/types/models"
)

type WalletClaims interface {
	GetClaimableWallet() *types.ClaimableWallet
	ListKeyshares() ([]models.KeyShare, error)
	IssueChallenge() (protocol.URLEncodedBase64, error)
}

type walletClaims struct {
	claims  *types.ClaimableWallet
	creator string
}

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
		claims: cw,
		creator: creator,
	}, nil
}

func LoadClaimableWallet(cw *types.ClaimableWallet) WalletClaims {
	return &walletClaims{
		claims: cw,
		creator: cw.Creator,
	}
}

func (wc *walletClaims) GetClaimableWallet() *types.ClaimableWallet {
	return wc.claims
}

func (wc *walletClaims) ListKeyshares() ([]models.KeyShare, error) {
	return keeper.GetKeysharesFromClaims(wc.claims)
}

func (wc *walletClaims) IssueChallenge() (protocol.URLEncodedBase64, error) {
	if wc.claims.PublicKey == "" {
		return nil, fmt.Errorf("public key is empty")
	}

	// Convert PublicKey to []byte
	pub, err := crypto.Base58Decode(wc.claims.PublicKey)
	if err != nil {
		return nil, err
	}
	return pub, nil
}
