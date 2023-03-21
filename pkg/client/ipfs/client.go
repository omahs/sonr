package ipfs

import (
	"bytes"
	"context"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonrhq/core/types/crypto"
)

const (
	// IPFS API address
	defaultIPFSApiAddr = "http://localhost:5001"
)

// Client is a wrapper around the IPFS shell client.
type Client struct {
	shell *shell.Shell
}

func NewClient(ipfsApiAddr string) *Client {
	return &Client{
		shell: shell.NewShell(ipfsApiAddr),
	}
}

// AddDirectory adds a local directory to IPFS.
func (c *Client) AddDirectory(dirPath string) (string, error) {
	return c.shell.AddDir(dirPath)
}


// ImportSecpk251PubKey imports a secpk251 public key with the given name and key bytes.
func (c *Client) ImportSecpk251PubKey(ctx context.Context, pubKey *crypto.PubKey, options ...shell.KeyImportOpt) (string, error) {
	addr, err := bech32.ConvertAndEncode("snr", pubKey.Bytes())
	if err != nil {
		return "", err
	}
	keyReader := bytes.NewReader(pubKey.Bytes())
	err = c.shell.KeyImport(ctx, addr, keyReader, options...)
	if err != nil {
		return "", err
	}

	return addr, nil
}


// TraverseLocalDirectory traverses the given local directory and performs the
// specified operation on each file.
func (c *Client) TraverseLocalDirectory(dirPath string, operation func(filePath string, info os.FileInfo) error) error {
	return filepath.Walk(dirPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		return operation(filePath, info)
	})
}
