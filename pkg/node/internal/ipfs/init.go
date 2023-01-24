package ipfs

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	nodeconfig "github.com/sonrhq/core/pkg/node/config"
	orbitdb "berty.tech/go-orbit-db"
	"berty.tech/go-orbit-db/iface"
	files "github.com/ipfs/go-ipfs-files"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/kubo/config"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/coreapi"
	klibp2p "github.com/ipfs/kubo/core/node/libp2p"
	"github.com/ipfs/kubo/plugin/loader"
	"github.com/ipfs/kubo/repo/fsrepo"
	snrConfig "github.com/sonrhq/core/pkg/node/config"
)

// Initialize creates a new local IPFS node
func Initialize(ctx context.Context, c *snrConfig.Config) (snrConfig.IPFSNode, error) {
	// Apply the options
	n := defaultNode(ctx, c)
	err := n.initialize()
	if err != nil {
		return nil, err
	}
	// Connect to the bootstrap nodes
	err = n.Connect(n.config.BootstrapMultiaddrs...)
	if err != nil {
		return nil, err
	}
	db, err := orbitdb.NewOrbitDB(ctx, n.CoreAPI(), &orbitdb.NewOrbitDBOptions{})
	if err != nil {
		return nil, err
	}
	n.orbitDb = db
	return n, nil
}

// Miscellanenous
var loadPluginsOnce sync.Once

// TopicMessageHandler is a function that handles a message received on a topic
type TopicMessageHandler func(topic string, msg icore.PubSubMessage) error

//
// Node Setup and Initialization
//

// defaultNode creates a new node with default options
func defaultNode(ctx context.Context, cnfg *snrConfig.Config) *localIpfs {
	return &localIpfs{
		ctx:    ctx,
		config: cnfg,
	}
}

// It's creating a new node and returning the coreAPI and the node itself.
func (c *localIpfs) initialize() error {
	// Spawn a local peer using a temporary path, for testing purposes
	var onceErr error
	loadPluginsOnce.Do(func() {
		onceErr = setupPlugins("")
	})
	if onceErr != nil {
		return onceErr
	}

	// Create a Temporary Repo
	repoPath, err := createHomeRepo()
	if err != nil {
		return fmt.Errorf("error creating temporary repo: %s", err)
	}

	node, err := createNode(c.ctx, repoPath)
	if err != nil {
		return err
	}

	api, err := coreapi.NewCoreAPI(node)
	if err != nil {
		return err
	}

	// Set the node and repoPath
	c.node = node
	c.repoPath = repoPath
	c.api = api
	return nil
}

// It loads plugins from the `externalPluginsPath` directory and injects them into the application
func setupPlugins(externalPluginsPath string) error {
	// Load any external plugins if available on externalPluginsPath
	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))
	if err != nil {
		return fmt.Errorf("error loading plugins: %s", err)
	}

	// Load preloaded and external plugins
	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	return nil
}

// It creates a temporary directory, initializes a new IPFS repo in that directory, and returns the
// path to the repo
func createHomeRepo() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home dir: %s", err)
	}
	repoPath := filepath.Join(homeDir, ".sonr", "ipfs")
	// Create a config with default options and a 2048 bit key
	cfg, err := config.Init(io.Discard, 2048)
	if err != nil {
		return "", err
	}
	// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#ipfs-filestore
	cfg.Experimental.FilestoreEnabled = true
	// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#ipfs-urlstore
	cfg.Experimental.UrlstoreEnabled = true
	// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#ipfs-p2p
	cfg.Experimental.Libp2pStreamMounting = true
	// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#p2p-http-proxy
	cfg.Experimental.P2pHttpProxy = true

	// Create the repo with the config
	err = fsrepo.Init(repoPath, cfg)
	if err != nil {
		return "", fmt.Errorf("failed to init ephemeral node: %s", err)
	}
	return repoPath, nil
}

// Creates an IPFS node and returns its coreAPI
func createNode(ctx context.Context, repoPath string) (*core.IpfsNode, error) {
	// Open the repo
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}

	err = repo.SetConfigKey("Pubsub.Enabled", true)
	if err != nil {
		return nil, err
	}
	err = repo.SetConfigKey("Pubsub.Router", "gossipsub")
	if err != nil {
		return nil, err
	}

	// Construct the node
	nodeOptions := &core.BuildCfg{
		Online:  true,
		Routing: klibp2p.DHTServerOption, // This option sets the node to be a full DHT node (both fetching and storing DHT Records)
		// Routing: libp2p.DHTClientOption, // This option sets the node to be a client DHT node (only fetching records)
		Repo: repo,
		ExtraOpts: map[string]bool{
			"pubsub": true,
			"ipnsps": true,
		},
	}

	// Create the node
	return core.NewNode(ctx, nodeOptions)
}

// It takes a path to a file or directory, and returns a UnixFS node
// It takes a path to a file or directory, and returns a UnixFS node
func getUnixfsNode(path string) (files.Node, error) {
	st, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	f, err := files.NewSerialFile(path, false, st)
	if err != nil {
		return nil, err
	}

	return f, nil
}

//
// Helper functions
//

// fetchDocsAddress fetches the address of the document store for a given username
func fetchDocsAddress(orb iface.OrbitDB, username string) (string, error) {
	addr, err := orb.DetermineAddress(context.Background(), username, nodeconfig.DB_DOCUMENT_STORE.String(), nil)
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}

// fetchEventLogAddress fetches the address of the event log for a given username
func fetchEventLogAddress(orb iface.OrbitDB, username string) (string, error) {
	addr, err := orb.DetermineAddress(context.Background(), username, nodeconfig.DB_EVENT_LOG_STORE.String(), nil)
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}

// fetchKeyValueAddress fetches the address of the key value store for a given username
func fetchKeyValueAddress(orb iface.OrbitDB, username string) (string, error) {
	addr, err := orb.DetermineAddress(context.Background(), username, nodeconfig.DB_KEY_VALUE_STORE.String(), nil)
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}
