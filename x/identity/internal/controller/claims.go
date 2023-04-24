package controller

import (
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/x/identity/internal/vault"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/identity/types/models"
	srvtypes "github.com/sonrhq/core/x/service/types"
)

type WalletClaims interface {
	GetClaimableWallet() *types.ClaimableWallet
	ListKeyshares() ([]models.KeyShare, error)
	IssueChallenge() (protocol.URLEncodedBase64, error)
	Assign(cred *srvtypes.WebauthnCredential, alias string) (Controller, error)
}

type walletClaims struct {
	Claims  *types.ClaimableWallet `json:"claims" yaml:"claims"`
	Creator string                 `json:"creator" yaml:"creator"`
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
		Claims:  cw,
		Creator: creator,
	}, nil
}

func LoadClaimableWallet(cw *types.ClaimableWallet) WalletClaims {
	return &walletClaims{
		Claims:  cw,
		Creator: cw.Creator,
	}
}

func (wc *walletClaims) GetClaimableWallet() *types.ClaimableWallet {
	return wc.Claims
}

func (wc *walletClaims) ListKeyshares() ([]models.KeyShare, error) {
	return vault.GetKeysharesFromClaims(wc.Claims)
}

func (wc *walletClaims) IssueChallenge() (protocol.URLEncodedBase64, error) {
	if wc.Claims.PublicKey == "" {
		return nil, fmt.Errorf("public key is empty")
	}
	return protocol.URLEncodedBase64(wc.Claims.PublicKey), nil
}

func (wc *walletClaims) Assign(cred *srvtypes.WebauthnCredential, alias string) (Controller, error) {
	kss := make([]models.KeyShare, 0)
	for _, ks := range wc.Claims.Keyshares {
		ks, err := vault.GetKeyshare(ks)
		if err != nil {
			return nil, err
		}
		kss = append(kss, ks)
	}

	acc := models.NewAccount(kss, crypto.SONRCoinType)
	doc := acc.DidDocument()
	credential := srvtypes.NewCredential(cred, doc.Id)
	vm := credential.ToVerificationMethod()
	_, err := doc.LinkAdditionalAuthenticationMethod(vm)
	if err != nil {
		return nil, err
	}

	cn := &didController{
		primary:    acc,
		primaryDoc: doc,
		blockchain: []models.Account{},
		txHash: "",
		disableIPFS: false,
		currCredential: cred,
	}
	resp, err := cn.CreatePrimaryIdentity(doc, acc, alias)
	if err != nil {
		return nil, err
	}
	cn.txHash = resp.TxResponse.TxHash
	return cn, nil
}
