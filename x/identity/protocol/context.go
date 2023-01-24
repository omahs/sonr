package protocol

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/crypto/nacl/box"
)

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

	// defaultRendezvousString is the default rendezvous string for the motor
	defaultRendezvousString = "sonr"
)

// `Context` is a struct that contains the information needed to run the `go-ipfs` node.
// @property {string} HomeDir - The home directory of the user running the application.
// @property {string} RepoPath - The path to the IPFS repo.
// @property {string} NodeRESTUri - The REST endpoint of the node.
// @property {string} NodeGRPCUri - The GRPC endpoint of the node.
// @property {string} NodeFaucetUri - The URI of the faucet service.
// @property {string} Rendevouz - The rendevouz point for the swarm.
// @property {[]string} BsMultiaddrs - The bootstrap multiaddrs.
// @property encPubKey - The public key of the encryption key pair.
// @property encPrivKey - The private key used to encrypt the data.
type Context struct {
	HomeDir       string
	RepoPath      string
	NodeRESTUri   string
	NodeGRPCUri   string
	NodeFaucetUri string
	Rendevouz     string
	BsMultiaddrs  []string

	encPubKey  *[32]byte
	encPrivKey *[32]byte
}

// It creates a new context object, initializes the encryption keys, and returns the context object
func NewContext(homeDir string) (Context, error) {
	ctx := Context{
		HomeDir:       homeDir,
		RepoPath:      filepath.Join(homeDir, ".sonr", "ipfs"),
		NodeRESTUri:   "http://api.sonr.network",
		NodeGRPCUri:   "grpc.sonr.network",
		NodeFaucetUri: "http://faucet.sonr.network",
		Rendevouz:     defaultRendezvousString,
		BsMultiaddrs:  defaultBootstrapMultiaddrs,
	}
	err := ctx.initEncKeys()
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

// Write encrypts a message using the box algorithm
// This encrypts msg and appends the result to the nonce.
func (b Context) EncryptMessage(msg []byte, peerPk []byte) []byte {
	return box.Seal(nil, msg, b.findNonce(peerPk), b.encPubKey, b.encPrivKey)
}

// The recipient can decrypt the message using their private key and the
// sender's public key. When you decrypt, you must use the same nonce you
// used to encrypt the message. One way to achieve this is to store the
// nonce alongside the encrypted message. Above, we stored the nonce in the
// first 24 bytes of the encrypted text.
func (b Context) DecryptMessage(encMsg []byte, peerPk []byte) ([]byte, bool) {
	return box.Open(nil, encMsg, b.findNonce(peerPk), b.encPubKey, b.encPrivKey)
}

// Nonce returns the nonce for the box
func (b Context) findNonce(peerPk []byte) *[24]byte {
	var nonce [24]byte
	copy(nonce[:], peerPk[:24])
	return &nonce
}

func (c Context) initEncKeys() error {
	if !hasKeys(c) {
		err := generateBoxKeys(c)
		if err != nil {
			return err
		}
	}
	pk, pb, err := loadBoxKeys(c)
	if err != nil {
		return err
	}
	c.encPrivKey = pk
	c.encPubKey = pb
	return nil
}

func kEncPrivKeyPath(cctx Context) string {
	return filepath.Join(cctx.HomeDir, ".sonr", "highway", "encryption_key")
}

func kEncPubKeyPath(cctx Context) string {
	return filepath.Join(cctx.HomeDir, ".sonr", "highway", "encryption_key.pub")
}

func hasEncryptionKey(cctx Context) bool {
	_, err := os.Stat(kEncPrivKeyPath(cctx))
	return err == nil
}

func hasEncryptionPubKey(cctx Context) bool {
	_, err := os.Stat(kEncPubKeyPath(cctx))
	return err == nil
}

func hasKeys(cctx Context) bool {
	return hasEncryptionKey(cctx) && hasEncryptionPubKey(cctx)
}

func generateBoxKeys(cctx Context) error {
	pub, priv, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(kEncPrivKeyPath(cctx)), 0755)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(kEncPubKeyPath(cctx)), 0755)
	if err != nil {
		return err
	}
	err = os.WriteFile(kEncPrivKeyPath(cctx), priv[:], 0600)
	if err != nil {
		return err
	}
	err = os.WriteFile(kEncPubKeyPath(cctx), pub[:], 0600)
	if err != nil {
		return err
	}
	return nil
}

func loadBoxKeys(cctx Context) (*[32]byte, *[32]byte, error) {
	if !hasKeys(cctx) {
		return nil, nil, fmt.Errorf("no keys found")
	}
	priv, err := os.ReadFile(kEncPrivKeyPath(cctx))
	if err != nil {
		return nil, nil, err
	}
	pub, err := os.ReadFile(kEncPubKeyPath(cctx))
	if err != nil {
		return nil, nil, err
	}
	var privKey [32]byte
	var pubKey [32]byte
	copy(privKey[:], priv)
	copy(pubKey[:], pub)
	return &privKey, &pubKey, nil
}
