package v1

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
	tmcrypto "github.com/tendermint/tendermint/crypto"
	"golang.org/x/crypto/ripemd160"
)

// It takes a name, index, address prefix, and a slice of shares, and returns an account config
func NewDerivedAccountConfig(name string, coinType crypto.CoinType, share *cmp.Config) (*AccountConfig, error) {
	pub, err := ExtractPubKeyFromConfig(share)
	if err != nil {
		return nil, err
	}
	shareConfigs, err := SerializeConfigList(share)
	if err != nil {
		return nil, err
	}
	return &AccountConfig{
		Name:          strings.ToLower(name),
		Multibase:     pub.Multibase(),
		Shares:        shareConfigs,
		CoinTypeIndex: int32(coinType.BipPath()),
		CreatedAt:     time.Now().Unix(),
		PublicKey:     pub.Raw(),
	}, nil
}

// It takes a name, index, address prefix, and a slice of shares, and returns an account config
func NewAccountConfigFromShares(name string, coinType crypto.CoinType, shares []*cmp.Config) (*AccountConfig, error) {
	pub, err := ExtractPubKeyFromConfig(shares[0])
	if err != nil {
		return nil, err
	}
	shareConfigs, err := SerializeConfigList(shares...)
	if err != nil {
		return nil, err
	}
	return &AccountConfig{
		Name:          strings.ToLower(name),
		Multibase:     pub.Multibase(),
		Shares:        shareConfigs,
		CoinTypeIndex: int32(coinType.BipPath()),
		CreatedAt:     time.Now().Unix(),
		PublicKey:     pub.Raw(),
	}, nil
}

// Returning the address of the account.
func (a *AccountConfig) Address() tmcrypto.Address {
	sha := sha256.Sum256(a.PublicKey)
	hasherRIPEMD160 := ripemd160.New()
	hasherRIPEMD160.Write(sha[:]) // does not error
	return tmcrypto.Address(hasherRIPEMD160.Sum(nil))
}

// Returning the coin type of the account.
func (a *AccountConfig) CoinType() crypto.CoinType {
	return crypto.CoinTypeFromBipPath(a.CoinTypeIndex)
}

// GetConfigAtIndex returns the config at the given index.
func (a *AccountConfig) GetConfigAtIndex(index int) (*cmp.Config, error) {
	if index >= len(a.Shares) {
		return nil, fmt.Errorf("index %d out of range", index)
	}
	share := a.Shares[index]
	conf := cmp.EmptyConfig(curve.Secp256k1{})
	if err := conf.UnmarshalBinary(share); err != nil {
		return nil, err
	}
	return conf, nil
}

// Creating a slice of party.Id from the shares.
func (a *AccountConfig) PartyIDs() []party.ID {
	ids := make([]party.ID, 0, len(a.Shares))
	for i := range a.Shares {
		share, err := a.GetConfigAtIndex(i)
		if err != nil {
			panic(err)
		}
		ids = append(ids, share.ID)
	}
	return ids
}

// Getting the public point from the first share.
func (a *AccountConfig) PublicPoint() (curve.Point, error) {
	share, err := a.GetConfigAtIndex(0)
	if err != nil {
		return nil, err
	}
	return share.PublicPoint(), nil
}

// PubKey returns the public key of the first share.
func (a *AccountConfig) PubKey() (*crypto.PubKey, error) {
	return crypto.NewSecp256k1PubKey(a.PublicKey), nil
}

// SerializeConfigList returns a slice of bytes representing the list of configs.
func SerializeConfigList(configs ...*cmp.Config) ([][]byte, error) {
	bz := make([][]byte, 0, len(configs))
	for _, config := range configs {
		b, err := config.MarshalBinary()
		if err != nil {
			return nil, err
		}
		bz = append(bz, b)
	}
	return bz, nil
}

// DeserializeConfigList returns a slice of configs from a slice of bytes.
func DeserializeConfigList(bz [][]byte) ([]*cmp.Config, error) {
	configs := make([]*cmp.Config, 0, len(bz))
	for _, b := range bz {
		config := cmp.EmptyConfig(curve.Secp256k1{})
		if err := config.UnmarshalBinary(b); err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}
	return configs, nil
}

// ExtractPubKeyFromConfig takes a `cmp.Config` and returns the public key
func ExtractPubKeyFromConfig(conf *cmp.Config) (*crypto.PubKey, error) {
	skPP, ok := conf.PublicPoint().(*curve.Secp256k1Point)
	if !ok {
		return nil, errors.New("invalid public point")
	}
	bz, err := skPP.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return crypto.NewSecp256k1PubKey(bz), nil
}
