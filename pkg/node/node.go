package node

import (
	"context"

	"berty.tech/go-orbit-db/iface"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/node/config"
	"github.com/sonrhq/core/pkg/node/internal/ipfs"
	"github.com/sonrhq/core/types/common"
)

// IPFSKVStore is an alias for a iface.KeyValueStore.
type iPFSKVStore = iface.KeyValueStore

// IPFSEventLogStore is an alias for a iface.EventLogStore.
type iPFSEventLogStore = iface.EventLogStore

// IPFSDocsStore is an alias for a iface.DocumentStore.
type iPFSDocsStore = iface.DocumentStore

// Callback is an alias for a common.NodeCallback
type Callback = config.NodeCallback

// IPFS is an alias for a common.IPFSNode.
type IPFS = config.IPFSNode

// P2P is an alias for a common.P2PNode.
type P2P = config.PeerNode

var (
	local IPFS
)

// `Node` is an interface that has three methods: `Host`, `IPFS`, and `Type`.
//
// The `Host` method returns a `Motor` interface and an error. The `IPFS` method returns a `Highway`
// interface and an error. The `Type` method returns a `Type` type.
//
// The `Motor` interface has two methods: `Start` and `Stop`. The `Start` method returns an error. The
// `Stop` method returns an error.
//
// The `Highway` interface has two methods: `Start` and
// @property Host - The motor that is hosting the node.
// @property IPFS - The IPFS node that the motor is connected to.
// @property {Type} Type - The type of node. This can be either a Motor or a Highway.
type Node interface {
	// Returning a Motor interface and an error.
	Host() P2P
	IPFS() IPFS
}

type node struct {
	host   config.PeerNode
	ipfs   config.IPFSNode
	config *config.Config
}

func (n *node) Host() P2P {
	return n.host
}

func (n *node) IPFS() IPFS {
	return n.ipfs
}

// // It creates a new host, and then creates a new node with that host
// func New(ctx context.Context, opts ...Option) (Node, error) {
// 	pctx, err := identityprotocol.NewContext(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	config := config.DefaultConfig(pctx)
// 	err = config.Apply(opts...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	i, err := ipfs.Initialize(config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &node{
// 		ipfs:   i,
// 		config: config,
// 	}, nil
// }

func init() {
	_ = StartLocalIPFS()
}

// StartLocalIPFS initializes a local IPFS node.
func StartLocalIPFS() error {
	// Start IPFS Node
	pctx, err := config.NewContext(context.Background())
	if err != nil {
		return err
	}
	config := config.DefaultConfig(pctx)
	err = config.Apply()
	if err != nil {
		return err
	}
	i, err := ipfs.Initialize(config)
	if err != nil {
		return err
	}
	local = i
	return nil
}

// NewIPFSKVStore creates a new IPFSKVStore. This requires a valid Sonr Account Public Key.
func NewIPFSStore(ctx context.Context, controller *crypto.PubKey) (IPFSStore, error) {
	if local == nil {
		return nil, common.ErrIPFSNotInitialized
	}
	name, err := bech32.ConvertAndEncode("snr", controller.Bytes())
	if err != nil {
		return nil, err
	}
	kv, err := local.LoadKeyValueStore(name)
	if err != nil {
		return nil, err
	}
	return makeIpfsStore(kv, controller), nil
}
