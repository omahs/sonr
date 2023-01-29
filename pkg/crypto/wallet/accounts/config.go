package accounts

import (
	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/crypto/wallet"
	"github.com/sonrhq/core/pkg/crypto/wallet/accounts/internal"
	"github.com/sonrhq/core/pkg/crypto/wallet/accounts/internal/mpc"
	"github.com/sonrhq/core/pkg/crypto/wallet/accounts/internal/network"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
)

var (
	// Default Account Name is the default name of the account.
	kDefaultAccName = "primary"

	// Default Threshold is the default number of required signatures to authorize a transaction.
	kDefaultThreshold = 1

	// Default CoinType is the default coin type of the account.
	kDefaultCoinType = common.CoinType_CoinType_SONR

	// Default Self ID is the default ID of the account that is used to sign transactions.
	kDefaultSelfID = "current"

	// Default Group is the default list of all parties that can sign transactions.
	kDefaultGroup = []string{"vault"}
)

// config is the configuration of an account used for MPC Interactions.
type config struct {
	// AccName is the name of the account.
	AccName string

	// Threshold is the number of required signatures to authorize a transaction.
	Threshold int

	// Self ID is the ID of the account that is used to sign transactions.
	ID string

	// Group is the list of all parties that can sign transactions.
	Peers []string

	// CoinType is the coin type of the account.
	CoinType common.CoinType
}

// defaultConfig returns the default configuration of an account.
func defaultConfig() *config {
	return &config{
		AccName:   kDefaultAccName,
		Threshold: kDefaultThreshold,
		ID:        kDefaultSelfID,
		Peers:     kDefaultGroup,
		CoinType:  kDefaultCoinType,
	}
}

// Keygen calls the Keygen function with the set values.
func (c *config) Keygen() (wallet.Account, error) {
	net := c.Network()
	accConf, rootCnf, err := mpc.Keygen(c.AccName, c.SelfID(), c.Threshold, net, c.CoinType)
	if err != nil {
		return nil, err
	}
	return internal.BaseAccountFromConfig(accConf, rootCnf), nil
}

// Network returns the network of the account.
func (c *config) Network() crypto.Network {
	return network.NewOfflineNetwork(c.PartyIDs())
}

// Keygen calls the Keygen function with the set values.
func (c *config) SelfID() party.ID {
	return party.ID(c.ID)
}

// PartyIDs returns the list of all parties that can sign transactions.
func (c *config) PartyIDs() party.IDSlice {
	pids := make(party.IDSlice, len(c.Peers))
	for i, p := range c.Peers {
		pids[i] = party.ID(p)
	}
	pids = append(pids, c.SelfID())
	return pids
}

// Option is a function that configures an account.
type Option func(*config)

// WithName sets the name of the account.
func WithName(name string) Option {
	return func(c *config) {
		c.AccName = name
	}
}

// WithThreshold sets the number of required signatures to authorize a transaction.
func WithThreshold(threshold int) Option {
	return func(c *config) {
		c.Threshold = threshold
	}
}

// WithSelfID sets the ID of the account that is used to sign transactions.
func WithSelfID(selfID string) Option {
	return func(c *config) {
		c.ID = selfID
	}
}

// WithPeers sets the list of all parties that can sign transactions.
func WithPeers(peers ...string) Option {
	return func(c *config) {
		c.Peers = peers
	}
}

// WithCoinType sets the coin type of the account.
func WithCoinType(coinType common.CoinType) Option {
	return func(c *config) {
		c.CoinType = coinType
	}
}
