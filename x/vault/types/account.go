package types

import (
	"crypto/sha1"
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	_ "github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/sonrhq/core/internal/mpc"
	"github.com/sonrhq/core/pkg/crypto"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
	"golang.org/x/crypto/pbkdf2"
	"lukechampine.com/blake3"
)

// Account is an interface for an account in the wallet
type Account interface {
	// Address returns the address of the account based on the coin type
	Address() string

	// CoinType returns the coin type of the account
	CoinType() crypto.CoinType

	// DeriveAccount returns a new account with the same keyshares but a new coin type
	DeriveAccount(ct crypto.CoinType, idx int, name string) (Account, error)

	// DID returns the DID of the account
	Did() string

	// GenerateSecretKey generates a secret phrase from a fragment and the account owner's DID
	GenerateSecretKey(fragment string) ([]byte, error)

	// GetAuthInfo creates an AuthInfo for a transaction
	GetAuthInfo(gas sdk.Coins) (*txtypes.AuthInfo, error)

	// ListKeyShares returns the list of keyshares of the account as a list of string dids
	ListKeyShares() []string

	// MapKeyShare maps a function to all the keyshares of the account
	MapKeyShare(f func(ks KeyShare) KeyShare)

	// PubKey returns secp256k1 public key
	PubKey() *crypto.PubKey

	// Signs a message
	Sign(bz []byte) ([]byte, error)

	// Signs a cosmos transaction
	SignCosmosTx(msgs ...sdk.Msg) ([]byte, error)

	// ToProto returns the proto representation of the account
	ToProto() *AccountInfo

	// Type returns the type of the account
	Type() string

	// Verifies a signature
	Verify(bz []byte, sig []byte) (bool, error)
}

type account struct {
	kss []KeyShare

	n    uint64
	p    string
	ct   crypto.CoinType
	cont string
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                     General                                    ||
// ! ||--------------------------------------------------------------------------------||

// NewAccount creates a new account
func NewAccount(kss []KeyShare, ct crypto.CoinType) Account {
	return &account{kss: kss, n: 0, p: "", ct: ct}
}

// Address returns the address of the account based on the coin type
func (a *account) Address() string {
	return a.CoinType().FormatAddress(a.PubKey())
}

// CoinType returns the coin type of the account
func (a *account) CoinType() crypto.CoinType {
	return a.ct
}

// DeriveAccount returns a new account with the same keyshares but a new coin type
func (a *account) DeriveAccount(ct crypto.CoinType, idx int, name string) (Account, error) {
	newAccCh := make(chan Account)
	errCh := make(chan error)

	go func() {
		var newKss []KeyShare
		for _, oldks := range a.kss {
			ks, err := oldks.DeriveBip44(ct, idx, name)
			if err != nil {
				errCh <- err
				return
			}
			newKss = append(newKss, ks)
		}
		newAccCh <- NewAccount(newKss, ct)
	}()

	// Create the new models.Account and map the keyshares to the resolver
	select {
	case newAcc := <-newAccCh:
		return newAcc, nil
	case err := <-errCh:
		return nil, err
	}
}

// GenerateSecretKey generates a new secret phrase of 32 bytes
func (a *account) GenerateSecretKey(fragment string) ([]byte, error) {
	sig, err := a.Sign([]byte(fragment))
	if err != nil {
		return nil, err
	}
	hash := blake3.Sum256(sig)
	pwd := pbkdf2.Key(hash[:], []byte(a.Address()), 10, 128, sha1.New)
	hashDerivKey := blake3.Sum256(pwd)
	return hashDerivKey[:], nil
}

// ListKeyShares returns the list of keyshares of the account as a list of string dids
func (a *account) ListKeyShares() []string {
	var ks []string
	for _, k := range a.kss {
		ks = append(ks, k.Did())
	}
	return ks
}

// MapKeyShare maps a function to all the keyshares of the account
func (acc *account) MapKeyShare(f func(ks KeyShare) KeyShare) {
	for i, ks := range acc.kss {
		acc.kss[i] = f(ks)
	}
}

// PubKey returns secp256k1 public key
func (wa *account) PubKey() *crypto.PubKey {
	return wa.kss[0].PubKey()
}

// Signs a message using the account
func (wa *account) Sign(bz []byte) ([]byte, error) {
	var configs []*cmp.Config
	for _, ks := range wa.kss {
		configs = append(configs, ks.Config())
	}
	return mpc.SignCMP(configs, bz)
}

// Signs a cosmos transaction
func (wa *account) SignCosmosTx(msgs ...sdk.Msg) ([]byte, error) {
	if !wa.CoinType().IsCosmos() && !wa.CoinType().IsSonr() {
		return nil, fmt.Errorf("coin type %s not supported for cosmos tx signing", wa.CoinType())
	}
	return SignAnyTransactions(wa, msgs...)
}

// ToProto returns the proto representation of the account
func (wa *account) ToProto() *AccountInfo {
	return &AccountInfo{
		Address:   wa.Address(),
		CoinType:  wa.CoinType().String(),
		Did:       wa.Did(),
		PublicKey: wa.PubKey().Base64(),
		Type:      wa.Type(),
	}
}

// Verifies a signature using first unlocked keyshare
func (wa *account) Verify(bz []byte, sig []byte) (bool, error) {
	// Find first unlocked keyshare
	var uks KeyShare
	for _, ks := range wa.kss {
		if ks.IsEncrypted() {
			continue
		}
		uks = ks
		break
	}
	if uks == nil {
		return false, fmt.Errorf("no unlocked keyshares")
	}
	return mpc.VerifyCMP(uks.Config(), bz, sig)
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                  Sonr Specific                                 ||
// ! ||--------------------------------------------------------------------------------||

// DID returns the DID of the account
func (wa *account) Did() string {
	return fmt.Sprintf("did:%s:%s", wa.CoinType().DidMethod(), wa.Address())
}

// Type returns the type of the account
func (wa *account) Type() string {
	return fmt.Sprintf("%s/ecdsa-secp256k1", wa.CoinType().Name())
}

//
// ! ||--------------------------------------------------------------------------------||
// ! ||                              Cosmos specific methods                           ||
// ! ||--------------------------------------------------------------------------------||
//

// GetAuthInfo creates an AuthInfo instance for this account with the specified gas amount.
func (wa *account) GetAuthInfo(gas sdk.Coins) (*txtypes.AuthInfo, error) {
	// Build signerInfo parameters
	seckpPubKey, err := wa.PubKey().Secp256k1()
	if err != nil {
		return nil, err
	}
	anyPubKey, err := codectypes.NewAnyWithValue(seckpPubKey)
	if err != nil {
		return nil, err
	}

	// Create AuthInfo
	authInfo := txtypes.AuthInfo{
		SignerInfos: []*txtypes.SignerInfo{
			{
				PublicKey: anyPubKey,
				ModeInfo: &txtypes.ModeInfo{
					Sum: &txtypes.ModeInfo_Single_{
						Single: &txtypes.ModeInfo_Single{
							Mode: 1,
						},
					},
				},
				Sequence: 0,
			},
		},
		Fee: &txtypes.Fee{
			Amount:   gas,
			GasLimit: uint64(300000),
		},
	}
	return &authInfo, nil
}

// // Lock encrypts all user-facing keyshares
// func (wa *account) Lock(c servicetypes.Credential) error {
// 	for _, ks := range wa.kss {
// 		if err := ks.Encrypt(c); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// // Unlock decrypts all user-facing keyshares
// func (wa *account) Unlock(c servicetypes.Credential) error {
// 	for _, ks := range wa.kss {
// 		if err := ks.Decrypt(c); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
