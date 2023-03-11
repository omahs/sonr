package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sonrhq/core/pkg/wallet"
	"github.com/sonrhq/core/pkg/wallet/accounts"
	vaultv1 "github.com/sonrhq/core/x/identity/types/vault/v1"
)

// Define the BIP44 prefix
const bip44Prefix = "m/44'/"

// Define the BIP32 path constants
const (
	purpose  = 44
	coinType = 703
)

// FindFilePathWithPrefix finds the full path of a file starting with a parent directory named "m/44'".
// Returns an error if the parent directory doesn't exist.
func FindFilePathWithPrefix(rootDir string, fileName string) (string, error) {
	prefix := "m/44'"
	dir := filepath.Join(rootDir, prefix)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return "", fmt.Errorf("directory %s doesn't exist", dir)
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Base(path) == fileName {
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	return filepath.Join(dir, fileName), nil
}

// WriteAccountConfig writes an AccountConfig to the specified file path in the
// BIP32 file system.
func WriteAccountConfig(filePath string, basePath string, accountConfig *vaultv1.AccountConfig) error {
	// Create or truncate the file at the specified path
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Marshal the AccountConfig to JSON and write to the file
	jsonBytes, err := json.MarshalIndent(accountConfig, "", "    ")
	if err != nil {
		return err
	}
	_, err = file.Write(jsonBytes)
	if err != nil {
		return err
	}

	return nil
}

// ReadAccountConfig reads an AccountConfig from the specified file path in the
// BIP32 file system.
func ReadAccountConfig(path string) (*vaultv1.AccountConfig, error) {
	// Open the file at the specified path
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the JSON-encoded AccountConfig from the file
	var accountConfig vaultv1.AccountConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&accountConfig)
	if err != nil {
		return nil, err
	}

	return &accountConfig, nil
}

// ListAccounts returns a list of all accounts in the BIP32 file system.
func ListAccounts() ([]wallet.Account, error) {
	var accs []wallet.Account

	files, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) != ".json" {
			continue
		}

		f, err := os.Open(file.Name())
		if err != nil {
			return nil, err
		}
		defer f.Close()

		var account vaultv1.AccountConfig
		err = json.NewDecoder(f).Decode(&account)
		if err != nil {
			return nil, err
		}
		acc, err := accounts.Load(&account)
		if err != nil {
			return nil, err
		}
		accs = append(accs, acc)
	}

	return accs, nil
}
