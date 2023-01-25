package v1

import (
	"errors"
	"time"

	"github.com/sonrhq/core/x/identity/types"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// It takes a `cmp.Config` and returns a `common.WalletShare` that can be used to create a wallet
func NewShareConfig(network string, conf *cmp.Config) *ShareConfig {
	skPP, ok := conf.PublicPoint().(*curve.Secp256k1Point)
	if !ok {
		return nil
	}
	pub, err := types.PubKeyFromCurvePoint(skPP)
	if err != nil {
		return nil
	}

	confBz, err := conf.MarshalBinary()
	if err != nil {
		return nil
	}
	return &ShareConfig{
		SelfId:     string(conf.ID),
		Network:    network,
		CreatedAt:  time.Now().Unix(),
		ConfigData: confBz,
		PublicKey:  pub.Raw(),
	}
}

// Unmarshalling the config data and returning the config.
func (s *ShareConfig) GetCMPConfig() (*cmp.Config, error) {
	conf := &cmp.Config{}
	if err := conf.UnmarshalBinary(s.ConfigData); err != nil {
		return nil, err
	}
	return conf, nil
}

// Converting the public key from the ShareConfig to a secp256k1.PubKey.
func (s *ShareConfig) GetCryptoPubKey() (*types.PubKey, error) {
	return types.NewPubKey(s.PublicKey, types.KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019), nil
}

// A method that returns the party ID of the share config.
func (s *ShareConfig) PartyID() party.ID {
	return party.ID(s.SelfId)
}

// Getting the public point from the first share.
func (a *ShareConfig) PublicPoint() (*curve.Secp256k1Point, error) {
	conf, err := a.GetCMPConfig()
	if err != nil {
		return nil, err
	}
	skPP, ok := conf.PublicPoint().(*curve.Secp256k1Point)
	if !ok {
		return nil, errors.New("invalid public point")
	}
	return skPP, nil
}
