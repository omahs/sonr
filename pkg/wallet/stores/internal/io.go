package internal

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

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

// WriteAccountConfig writes an AccountConfig to the specified file path in the
// BIP32 file system.
func WriteAccountConfig(path string, accountConfig *vaultv1.AccountConfig) error {
	// Create the parent directories for the file path
	err := os.MkdirAll(filepath.Join(bip44Prefix, strconv.Itoa(int(coinType))), os.ModePerm)
	if err != nil {
		return err
	}

	// Create or truncate the file at the specified path
	file, err := os.Create(filepath.Join(bip44Prefix, strconv.Itoa(int(coinType)), path))
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
