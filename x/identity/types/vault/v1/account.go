package v1

import (
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	types "github.com/sonrhq/core/x/identity/types"

	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// It takes a name, index, address prefix, and a slice of shares, and returns an account config
func NewAccountConfigFromShares(name string, index uint32, addrPrefix string, shares []*ShareConfig) (*AccountConfig, error) {
	addr, err := bech32.ConvertAndEncode(addrPrefix, shares[0].PublicKey)
	if err != nil {
		return nil, err
	}
	return &AccountConfig{
		Name:         strings.ToLower(name),
		Index:        index,
		Address:      addr,
		PublicKey:    shares[0].PublicKey,
		Shares:       shares,
		Bech32Prefix: addrPrefix,
		CreatedAt:    time.Now().Unix(),
	}, nil
}

// Creating a map of party.ID to cmp.Config.
func (a *AccountConfig) GetConfigMap() map[party.ID]*cmp.Config {
	configMap := make(map[party.ID]*cmp.Config)
	for _, s := range a.Shares {
		conf, err := s.GetCMPConfig()
		if err != nil {
			continue
		}
		configMap[s.PartyID()] = conf
	}
	return configMap
}

// Creating a slice of party.ID from the shares.
func (a *AccountConfig) PartyIDs() []party.ID {
	ids := make([]party.ID, 0, len(a.Shares))
	for _, share := range a.Shares {
		ids = append(ids, party.ID(share.SelfId))
	}
	return ids
}

// Getting the public point from the first share.
func (a *AccountConfig) PublicPoint() (curve.Point, error) {
	return a.Shares[0].PublicPoint()
}

// GetCryptoPubKey returns the public key of the first share.
func (a *AccountConfig) GetCryptoPubKey() (*types.PubKey, error) {
	return a.Shares[0].GetCryptoPubKey()
}
