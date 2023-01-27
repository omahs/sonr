package account

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/golang-jwt/jwt"
	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/x/identity/protocol/vault/account/internal/mpc"
	"github.com/sonrhq/core/x/identity/protocol/vault/account/internal/network"
	v1 "github.com/sonrhq/core/x/identity/types/vault/v1"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
	"github.com/ucan-wg/go-ucan"
)

// ErrInvalidToken indicates an access token is invalid
var ErrInvalidToken = errors.New("invalid access token")

const (
	// UCANVersion is the current version of the UCAN spec
	UCANVersion = "0.7.0"
	// UCANVersionKey is the key used in version headers for the UCAN spec
	UCANVersionKey = "ucv"
	// PrfKey denotes "Proofs" in a UCAN. Stored in JWT Claims
	PrfKey = "prf"
	// FctKey denotes "Facts" in a UCAN. Stored in JWT Claims
	FctKey = "fct"
	// AttKey denotes "Attenuations" in a UCAN. Stored in JWT Claims
	AttKey = "att"
	// CapKey indicates a resource Capability. Used in an attenuation
	CapKey = "cap"
)

// `WalletAccount` is an interface that defines the methods that a wallet account must implement.
// @property AccountConfig - The account configuration
// @property Bip32Derive - This is a method that derives a new account from a BIP32 path.
// @property GetAssertionMethod - returns the verification method for the account.
// @property {bool} IsPrimary - returns true if the account is the primary account
// @property ListConfigs - This is a list of all the configurations that are needed to sign a
// transaction.
// @property Sign - This is the function that signs a transaction.
// @property Verify - Verifies a signature
type WalletAccount interface {
	// The account configuration
	AccountConfig() *v1.AccountConfig

	// Bip32Derive derives a new account from a BIP32 path
	Bip32Derive(accName string, coinType common.CoinType) (WalletAccount, error)

	// Bytes returns the bytes of the account
	Bytes() []byte

	// Equals returns true if the account is equal to the other account
	Equals(other cryptotypes.LedgerPrivKey) bool

	// GetSignerData returns the signer data for the account
	GetSignerData() authsigning.SignerData

	// Info returns the account information
	Info() *v1.AccountInfo

	// IsPrimary returns true if the account is the primary account
	IsPrimary() bool

	// ListConfigs returns the list of all the configurations that are needed to
	// sign a transaction.
	ListConfigs() ([]*cmp.Config, error)

	// NewOriginToken creates a new UCAN token
	NewOriginToken(audienceDID string, att ucan.Attenuations, fct []ucan.Fact, notBefore, expires time.Time) (*ucan.Token, error)

	// NewAttenuatedToken creates a new UCAN token from the parent token
	NewAttenuatedToken(parent *ucan.Token, audienceDID string, att ucan.Attenuations, fct []ucan.Fact, notBefore, expires time.Time) (*ucan.Token, error)

	// PubKey returns secp256k1 public key
	PubKey() common.SNRPubKey

	// Signs a message
	Sign(bz []byte) ([]byte, error)

	// Signs a transaction
	SignTxAux(msgs ...sdk.Msg) (txtypes.AuxSignerData, error)

	// Type returns the type of the account
	Type() string

	// Verifies a signature
	Verify(bz []byte, sig []byte) (bool, error)
}

// The walletAccountImpl type is a struct that has a single field, accountConfig, which is a pointer to
// a v1.AccountConfig.
// @property accountConfig - The account configuration
type walletAccountImpl struct {
	// The account configuration
	accountConfig *v1.AccountConfig
}

// It creates a new account with the given name, address prefix, and network name
func NewAccount(accName string, coinType common.CoinType) (WalletAccount, error) {
	// The default shards that are added to the MPC wallet
	parties := party.IDSlice{"vault", "current"}
	net := network.NewOfflineNetwork(parties)
	accConf, err := mpc.Keygen(strings.ToLower(accName), "current", 1, net, coinType)
	if err != nil {
		return nil, err
	}
	return &walletAccountImpl{
		accountConfig: accConf,
	}, nil
}

// > This function creates a new wallet account from the given account configuration
func NewAccountFromConfig(accConf *v1.AccountConfig) (WalletAccount, error) {
	return &walletAccountImpl{
		accountConfig: accConf,
	}, nil
}

// It returns the account configuration.
func (w *walletAccountImpl) AccountConfig() *v1.AccountConfig {
	return w.accountConfig
}

// Deriving a new account from a BIP32 path.
func (w *walletAccountImpl) Bip32Derive(accName string, coinType common.CoinType) (WalletAccount, error) {
	if !w.IsPrimary() {
		return nil, errors.New("cannot derive from non-primary account")
	}
	oldConfs, err := w.ListConfigs()
	if err != nil {
		return nil, err
	}
	c, err := oldConfs[0].DeriveBIP32(uint32(coinType.Index()))
	if err != nil {
		return nil, err
	}
	conf, err := v1.NewDerivedAccountConfig(accName, coinType, c)
	if err != nil {
		return nil, err
	}
	return NewAccountFromConfig(conf)
}

// Bytes returns the bytes of the account.
func (w *walletAccountImpl) Bytes() []byte {
	return w.PubKey().Bytes()
}

// Equals returns true if the account is equal to the other account.
func (w *walletAccountImpl) Equals(other cryptotypes.LedgerPrivKey) bool {
	return bytes.Equal(w.Bytes(), other.Bytes())
}

// Returning the signer data for the account.
func (w *walletAccountImpl) GetSignerData() authsigning.SignerData {
	return authsigning.SignerData{
		Address:       w.accountConfig.DID(),
		ChainID:       "sonr",
		AccountNumber: 0,
		Sequence:      0,
		PubKey:        w.PubKey(),
	}
}

// Returning the account information.
func (w *walletAccountImpl) Info() *v1.AccountInfo {
	addr, _ := w.accountConfig.Address()
	return &v1.AccountInfo{
		Label:   w.accountConfig.Name,
		Address: addr,
		Index:   w.accountConfig.CoinType().PathComponent(),
		Network: w.accountConfig.CoinType().Name(),
	}
}

// It returns true if the account is the primary account.
func (w *walletAccountImpl) IsPrimary() bool {
	return w.accountConfig.Name == "primary"
}

// Returning the list of all the configurations that are needed to sign a transaction.
func (w *walletAccountImpl) ListConfigs() ([]*cmp.Config, error) {
	return v1.DeserializeConfigList(w.accountConfig.Shares)
}

// NewOriginToken returns a new origin token for the account.
func (w *walletAccountImpl) NewOriginToken(audienceDID string, att ucan.Attenuations, fct []ucan.Fact, notBefore, expires time.Time) (*ucan.Token, error) {
	return newToken(w.accountConfig, audienceDID, nil, att, fct, notBefore, expires)
}

// NewAttenuatedToken returns a new attenuated token for the account.
func (w *walletAccountImpl) NewAttenuatedToken(parent *ucan.Token, audienceDID string, att ucan.Attenuations, fct []ucan.Fact, notBefore, expires time.Time) (*ucan.Token, error) {
	if !parent.Attenuations.Contains(att) {
		return nil, fmt.Errorf("scope of ucan attenuations must be less than it's parent")
	}
	return newToken(w.accountConfig, audienceDID, append(parent.Proofs, ucan.Proof(parent.Raw)), att, fct, notBefore, expires)
}

// Returning the secp256k1 public key.
func (w *walletAccountImpl) PubKey() common.SNRPubKey {
	pbKey, _ := w.accountConfig.PubKey()
	return pbKey
}

// Signing a transaction.
func (w *walletAccountImpl) Sign(bz []byte) ([]byte, error) {
	return signWithAccount(w.accountConfig, bz)
}

// Signing a transaction.
func (w *walletAccountImpl) SignTxAux(msgs ...sdk.Msg) (txtypes.AuxSignerData, error) {
	txBody, err := buildTx(msgs...)
	if err != nil {
		return txtypes.AuxSignerData{}, err
	}
	doc, sig, err := signTxDocDirectAux(w, txBody)
	if err != nil {
		return txtypes.AuxSignerData{}, err
	}
	addr, _ := w.accountConfig.Address()
	return txtypes.AuxSignerData{
		Address: addr,
		Mode:    signing.SignMode_SIGN_MODE_DIRECT_AUX,
		SignDoc: doc,
		Sig:     sig,
	}, nil
}

// Type returns the type of the account.
func (w *walletAccountImpl) Type() string {
	return "ECDSA-SECP256K1"
}

// Verifying a signature.
func (w *walletAccountImpl) Verify(bz []byte, sig []byte) (bool, error) {
	pubKey, err := w.accountConfig.PubKey()
	if err != nil {
		return false, err
	}
	return pubKey.VerifySignature(bz, sig), nil
}

// It takes a list of messages, creates a transaction body, and marshals it
func buildTx(msgs ...sdk.Msg) ([]byte, error) {
	// func BuildTx(w *crypto.MPCWallet, msgs ...sdk.Msg) (*txtypes.TxBody, error) {
	// Create Any for each message
	anyMsgs, err := txtypes.SetMsgs(msgs)
	if err != nil {
		return nil, err
	}

	// Create TXRaw and Marshal
	txBody := &txtypes.TxBody{
		Messages: anyMsgs,
		Memo:     "xP;",
	}
	return txBody.Marshal()
}

// It creates a new token, signs it, and returns it
func newToken(a *v1.AccountConfig, audienceDID string, prf []ucan.Proof, att ucan.Attenuations, fct []ucan.Fact, nbf, exp time.Time) (*ucan.Token, error) {
	t := jwt.New(jwt.SigningMethodES256)

	t.Header[UCANVersionKey] = UCANVersion

	var (
		nbfUnix int64
		expUnix int64
	)

	if !nbf.IsZero() {
		nbfUnix = nbf.Unix()
	}
	if !exp.IsZero() {
		expUnix = exp.Unix()
	}
	pub, err := a.PubKey()
	if err != nil {
		return nil, err
	}

	// set our claims
	t.Claims = &ucan.Claims{
		StandardClaims: &jwt.StandardClaims{
			Issuer:    pub.DID(),
			Audience:  audienceDID,
			NotBefore: nbfUnix,
			// set the expire time
			// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
			ExpiresAt: expUnix,
		},
		Attenuations: att,
		Facts:        fct,
		Proofs:       prf,
	}
	sig, err := t.SigningString()
	if err != nil {
		return nil, err
	}
	raw, err := signWithAccount(a, []byte(sig))
	if err != nil {
		return nil, err
	}
	return &ucan.Token{
		Raw:          string(raw),
		Attenuations: att,
		Facts:        fct,
		Proofs:       prf,
	}, nil
}

// Signing a transaction.
func signTxDocDirectAux(w *walletAccountImpl, txBody []byte) (*txtypes.SignDocDirectAux, []byte, error) {
	// Build the public key.
	pk, err := codectypes.NewAnyWithValue(w.PubKey())
	if err != nil {
		return nil, nil, err
	}

	// Build the sign doc.
	doc := &txtypes.SignDocDirectAux{
		ChainId:   "sonr",
		PublicKey: pk,
		BodyBytes: txBody,
	}

	// Marshal the document.
	bz, err := doc.Marshal()
	if err != nil {
		return nil, nil, err
	}

	// Sign the document.
	sig, err := w.Sign(bz)
	if err != nil {
		return nil, nil, err
	}

	// Return the document and the signature.
	return doc, sig, nil
}

// signWithAccount signs a message with the given account configuration
func signWithAccount(a *v1.AccountConfig, msg []byte) ([]byte, error) {
	net := network.NewOfflineNetwork(a.PartyIDs())
	configs, err := v1.DeserializeConfigList(a.Shares)
	if err != nil {
		return nil, err
	}

	doneChan := make(chan []byte, 1)
	var wg sync.WaitGroup
	for _, c := range configs {
		wg.Add(1)
		go func(conf *cmp.Config) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			sig, err := mpc.CmpSign(conf, msg, net.Ls(), net, &wg, pl)
			if err != nil {
				return
			}
			if conf.ID == "current" {
				doneChan <- sig
			}
		}(c)
	}
	wg.Wait()
	return <-doneChan, nil
}
