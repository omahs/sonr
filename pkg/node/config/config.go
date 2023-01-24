package config

import (
	"github.com/cosmos/cosmos-sdk/client"
	p2pcrypto "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/shengdoushi/base58"
	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
)

// StoreType is the type of a store
type StoreType string

const (
	// DB_EVENT_LOG_STORE is a store that stores events
	DB_EVENT_LOG_STORE StoreType = "eventlog"

	// DB_KEY_VALUE_STORE is a store that stores key-value pairs
	DB_KEY_VALUE_STORE StoreType = "keyvalue"

	// DB_DOCUMENT_STORE is a store that stores documents
	DB_DOCUMENT_STORE StoreType = "docstore"
)

// A method of the StoreType type.
func (st StoreType) String() string {
	return string(st)
}

// Default configuration
var (
	// defaultBootstrapMultiaddrs is the default list of bootstrap nodes
	defaultBootstrapMultiaddrs = []string{
		// IPFS Bootstrapper nodes.
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		// "/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		// "/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		// "/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",

		// IPFS Cluster Pinning nodes
		// "/ip4/138.201.67.219/tcp/4001/p2p/QmUd6zHcbkbcs7SMxwLs48qZVX3vpcM8errYS7xEczwRMA",
		// "/ip4/138.201.67.219/udp/4001/quic/p2p/QmUd6zHcbkbcs7SMxwLs48qZVX3vpcM8errYS7xEczwRMA",
		// "/ip4/138.201.67.220/tcp/4001/p2p/QmNSYxZAiJHeLdkBg38roksAR9So7Y5eojks1yjEcUtZ7i",
		// "/ip4/138.201.67.220/udp/4001/quic/p2p/QmNSYxZAiJHeLdkBg38roksAR9So7Y5eojks1yjEcUtZ7i",
		// "/ip4/138.201.68.74/tcp/4001/p2p/QmdnXwLrC8p1ueiq2Qya8joNvk3TVVDAut7PrikmZwubtR",
		// "/ip4/138.201.68.74/udp/4001/quic/p2p/QmdnXwLrC8p1ueiq2Qya8joNvk3TVVDAut7PrikmZwubtR",
		// "/ip4/94.130.135.167/tcp/4001/p2p/QmUEMvxS2e7iDrereVYc5SWPauXPyNwxcy9BXZrC1QTcHE",
		// "/ip4/94.130.135.167/udp/4001/quic/p2p/QmUEMvxS2e7iDrereVYc5SWPauXPyNwxcy9BXZrC1QTcHE",

		// You can add more nodes here, for example, another IPFS node you might have running locally, mine was:
		// "/ip4/127.0.0.1/tcp/4010/p2p/QmZp2fhDLxjYue2RiUvLwT9MWdnbDxam32qYFnGmxZDh5L",
		// "/ip4/127.0.0.1/udp/4010/quic/p2p/QmZp2fhDLxjYue2RiUvLwT9MWdnbDxam32qYFnGmxZDh5L",
	}

	// defaultCallback is the default callback for the motor
	defaultCallback = common.DefaultCallback()

	// defaultRendezvousString is the default rendezvous string for the motor
	defaultRendezvousString = "sonr"

	// Remote API address
	defaultAPIAddr = "/ip4/198.199.78.62/tcp/9094"
)

// Config is the configuration for the node
type Config struct {
	CCtx client.Context

	// BootstrapMultiaddrs is the list of bootstrap nodes
	BootstrapMultiaddrs []string

	// Callback is the callback for the motor
	Callback common.NodeCallback

	// RendezvousString is the rendezvous string for the motor
	RendezvousString string

	// GroupIDs is the list of peer ids for the node
	GroupIDs []party.ID

	// SelfPartyID is the party id for the node
	SelfPartyID party.ID

	// PeerType is the type of peer
	PeerType common.PeerType

	// RemoteIPFSURL is the remote IPFS URL
	RemoteIPFSURL string

	// EncryptionKey is the encryption key for the node
	EncryptionKey p2pcrypto.PrivKey

	// EncryptionPrivKeyPath is the encryption key for the node
	EncryptionPrivKeyPath string

	// EncryptionPubKeyPath is the encryption key for the node
	EncryptionPubKeyPath string
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		BootstrapMultiaddrs: defaultBootstrapMultiaddrs,
		Callback:            defaultCallback,
		RendezvousString:    defaultRendezvousString,
		RemoteIPFSURL:       defaultAPIAddr,
		PeerType:            common.PeerType_HIGHWAY,
		SelfPartyID:         party.ID("current"),
	}
}

// Apply applies the options to the configuration
func (c *Config) Apply(opts ...Option) error {
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return err
		}
	}
	return nil
}

// GetCapabilityDelegation returns the capability delegation
func (c *Config) GetCapabilityDelegation() *types.VerificationMethod {
	_, pubKey, err := c.LoadEncKeys()
	if err != nil {
		return nil
	}
	return &types.VerificationMethod{
		ID:                 types.ConvertAccAddressToDid(c.CCtx.FromAddress),
		Type:               types.KeyType_KeyType_ED25519_VERIFICATION_KEY_2018,
		Controller:         types.ConvertAccAddressToDid(c.CCtx.FromAddress),
		PublicKeyMultibase: base58.Encode(pubKey[:], base58.BitcoinAlphabet),
	}
}

// LoadEncKeys loads the encryption keys
func (c *Config) LoadEncKeys() (*[32]byte, *[32]byte, error) {
	return loadBoxKeys(c.CCtx)
}

// IsLocal returns true if the node is local
func (c *Config) IsLocal() bool {
	return !c.IsMotor()
}

// IsMotor returns true if the node is a motor
func (c *Config) IsMotor() bool {
	return c.PeerType == common.PeerType_MOTOR
}

// Option is a function that configures a Node
type Option func(*Config) error

// AddBootstrappers adds additional nodes to start initial connections with
func AddBootstrappers(bootstrappers []string) Option {
	return func(c *Config) error {
		c.BootstrapMultiaddrs = append(c.BootstrapMultiaddrs, bootstrappers...)
		return nil
	}
}

// SetRemoteIPFS sets the remote IPFS URL
func SetRemoteIPFS(remoteIPFSURL string) Option {
	return func(c *Config) error {
		c.RemoteIPFSURL = remoteIPFSURL
		return nil
	}
}

// WithGroupIds sets the peer ids for the node
func WithGroupIds(partyIds ...party.ID) Option {
	return func(c *Config) error {
		if len(partyIds) > 0 {
			c.GroupIDs = partyIds
		}
		return nil
	}
}

// WithNodeCallback sets the callback for the motor
func WithNodeCallback(callback common.NodeCallback) Option {
	return func(c *Config) error {
		c.Callback = callback
		return nil
	}
}

// WithPartyId sets the party id for the node. This is to be replaced by the User defined label for the device
func WithPartyId(partyId string) Option {
	return func(c *Config) error {
		c.SelfPartyID = party.ID(partyId)
		return nil
	}
}

// WithPeerType sets the type of peer
func WithPeerType(peerType common.PeerType) Option {
	return func(c *Config) error {
		c.PeerType = peerType
		return nil
	}
}

// WithRemoteIPFSURL sets the remote IPFS URL
func WithRemoteIPFSURL(remoteIPFSURL string) Option {
	return func(c *Config) error {
		c.RemoteIPFSURL = remoteIPFSURL
		return nil
	}
}

// WithEncryptionKeyPath sets the encryption private key for the node from a file
func WithClientContext(cctx client.Context, generate bool) Option {
	return func(c *Config) error {
		c.CCtx = cctx
		if hasKeys(cctx) {
			c.EncryptionPrivKeyPath = kEncPrivKeyPath(cctx)
			c.EncryptionPubKeyPath = kEncPubKeyPath(cctx)
		}
		if generate {
			err := generateBoxKeys(cctx)
			if err != nil {
				return err
			}
			c.EncryptionPrivKeyPath = kEncPrivKeyPath(cctx)
			c.EncryptionPubKeyPath = kEncPubKeyPath(cctx)
		}
		return nil
	}
}
