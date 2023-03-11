package wallet

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sonrhq/core/pkg/crypto"
)

const (
	DEFAULT_WALLET_PATH = "_SNR_WALLET_"
)

// GetAccountsByCoinType returns all the accounts for the specified coin type in the
// given wallet path.
func GetAccountsByCoinType(basePath string, coinType crypto.CoinType) ([]string, error) {
	// Validate the coin type
	if coinType < 0 || coinType > crypto.XRPCoinType {
		return nil, fmt.Errorf("invalid coin type")
	}

	// Generate the path to the directory containing the accounts for the specified coin type
	coinTypeStr := fmt.Sprintf("%d'", coinType)
	coinTypePath := filepath.Join(basePath, DEFAULT_WALLET_PATH, coinTypeStr)

	// Check if the coin type directory exists
	if _, err := os.Stat(coinTypePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("coin type directory does not exist")
	}

	// List all the account files in the coin type directory
	accountFiles, err := ioutil.ReadDir(coinTypePath)
	if err != nil {
		return nil, err
	}

	// Extract the account names from the file names
	var accounts []string
	for _, accountFile := range accountFiles {
		accountName := accountFile.Name()
		accountName = accountName[:len(accountName)-5] // Remove ".json" extension
		accounts = append(accounts, accountName)
	}
	return accounts, nil
}
