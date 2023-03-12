package v2

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/crypto/mpc"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// Account is an interface for an account in the wallet
type Account interface {
	// Address returns the address of the account based on the coin type
	Address() string

	// CoinType returns the coin type of the account
	CoinType() crypto.CoinType

	// DID returns the DID of the account
	DID() string

	// ListKeyshares returns a list of keyshares for the account
	ListKeyshares() ([]KeyShare, error)

	// Name returns the name of the account
	Name() string

	// PartyIDs returns the party IDs of the account
	PartyIDs() []crypto.PartyID

	// PubKey returns secp256k1 public key
	PubKey() *crypto.PubKey

	// Rename renames the account
	Rename(name string) error

	// Signs a message
	Sign(bz []byte) ([]byte, error)

	// Type returns the type of the account
	Type() string

	// Verifies a signature
	Verify(bz []byte, sig []byte) (bool, error)
}

type walletAccount struct {
	p     string
	files []fs.FileInfo
}

// NewWalletAccount loads an accound directory and returns a WalletAccount
func NewWalletAccount(p string) (Account, error) {
	// Check if the path is a directory
	ok, files := isDir(p)
	if !ok {
		return nil, fmt.Errorf("path %s is not a directory", p)
	}
	if !hasShare(files) {
		return nil, fmt.Errorf("directory %s does not contain any MPC shard files", p)
	}
	return &walletAccount{p: p}, nil
}

// Address returns the address of the account based on the coin type
func (wa *walletAccount) Address() string {
	return ""
}

// CoinType returns the coin type of the account
func (wa *walletAccount) CoinType() crypto.CoinType {
	parentDir := filepath.Base(filepath.Dir(wa.p))
	allCoins := crypto.AllCoinTypes()
	for _, coin := range allCoins {
		if strings.Contains(parentDir, fmt.Sprintf("%d", coin.BipPath())) {
			return coin
		}
	}
	return crypto.TestCoinType
}

// DID returns the DID of the account
func (wa *walletAccount) DID() string {
	if wa.CoinType().IsSonr() {
		return fmt.Sprintf("did:%s:%s", wa.CoinType().DidMethod(), wa.Address())
	}
	return fmt.Sprintf("did:%s:%s#%s", wa.CoinType().DidMethod(), wa.Address(), wa.Name())
}

// ListKeyshares returns a list of keyshares for the account
func (wa *walletAccount) ListKeyshares() ([]KeyShare, error) {
	var keyshares []KeyShare
	for _, f := range wa.files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".key" {
			ks, err := NewKeyshare(filepath.Join(wa.p, f.Name()))
			if err != nil {
				return nil, err
			}
			keyshares = append(keyshares, ks)
		}
	}
	return keyshares, nil
}

// Name returns the name of the account
func (wa *walletAccount) Name() string {
	return filepath.Base(wa.p)
}

// PartyIDs returns the party IDs of the account
func (wa *walletAccount) PartyIDs() []crypto.PartyID {
	var partyIDs []crypto.PartyID
	for _, f := range wa.files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".key" {
			id := strings.TrimRight(f.Name(), ".key")
			partyIDs = append(partyIDs, crypto.PartyID(id))
		}
	}
	return partyIDs
}

// PubKey returns secp256k1 public key
func (wa *walletAccount) PubKey() *crypto.PubKey {
	ks, err := NewKeyshare(wa.files[0].Name())
	if err != nil {
		return nil
	}
	skPP, ok := ks.Config().PublicPoint().(*curve.Secp256k1Point)
	if !ok {
		return nil
	}
	bz, err := skPP.MarshalBinary()
	if err != nil {
		return nil
	}
	return crypto.NewSecp256k1PubKey(bz)
}

// Rename renames the account
func (wa *walletAccount) Rename(name string) error {
	return os.Rename(wa.p, filepath.Join(filepath.Dir(wa.p), name))
}

// Signs a message using the account
func (wa *walletAccount) Sign(bz []byte) ([]byte, error) {
	var configs []*cmp.Config
	for _, f := range wa.files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".key" {
			ks, err := NewKeyshare(filepath.Join(wa.p, f.Name()))
			if err != nil {
				return nil, err
			}
			configs = append(configs, ks.Config())
		}
	}
	return mpc.SignCMP(configs, bz, wa.PartyIDs())
}

// Type returns the type of the account
func (wa *walletAccount) Type() string {
	return fmt.Sprintf("%s/ecdsa-secp256k1", wa.CoinType().Name())
}

// Verifies a signature
func (wa *walletAccount) Verify(bz []byte, sig []byte) (bool, error) {
	ks, err := NewKeyshare(wa.files[0].Name())
	if err != nil {
		return false, err
	}
	return mpc.VerifyCMP(ks.Config(), bz, sig)
}

//
// Helper functions
//

// isDir checks if the path is a directory and contains at least one MPC shard file
func isDir(p string) (bool, []fs.FileInfo) {
	fi, err := os.Stat(p)
	if err != nil {
		return false, nil
	}
	if !fi.IsDir() {
		return false, nil
	}
	// Check if the directory contains at least one MPC shard file
	files, err := ioutil.ReadDir(p)
	if err != nil {
		return false, nil
	}
	return true, files
}

// hasShare checks if the directory contains at least one MPC shard file
func hasShare(files []fs.FileInfo) bool {
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".key" {
			return true
		}
	}
	return false
}
