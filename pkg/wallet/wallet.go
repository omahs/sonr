package v2

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/crypto/mpc"
	"github.com/sonrhq/core/types/common"
)

type Wallet interface {
	// Get the wallet's controller
	Controller() string

	// CreateAccount creates a new account for the given coin type
	CreateAccount(coin crypto.CoinType) (Account, error)

	// ListAllocatedCoins returns a list of coins that this currently has accounts for
	ListCoins() ([]crypto.CoinType, error)

	// ListAccounts returns a list of accounts for the given coin type
	ListAccounts() (map[crypto.CoinType][]Account, error)

	// ListAccountsForCoin returns a list of accounts for the given coin type
	ListAccountsForCoin(coin crypto.CoinType) ([]Account, error)

	// GetAccount returns the account for the given coin type and account name
	GetAccount(coin crypto.CoinType, name string) (Account, error)

	// GetAccountByAddress returns the account for the given address and parses the coin type from the address
	GetAccountByAddress(address string) (Account, error)

	// GetAccountByDID returns the account for the given DID and parses the coin type from the DID
	GetAccountByDID(did string) (Account, error)

	// GetAccountByPublicKey returns the account for the given public key and parses the coin type from the public key
	GetAccountByPublicKey(key string) (Account, error)
}

type wallet struct {
	currentId string
	threshold int
	path      string

	info *common.WalletInfo

	fileStore *FileStore
}

func NewWallet(currentId string, threshold int) (Wallet, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(homeDir, "Desktop", "_SONR_WALLET_")
	fs, err := NewFileStore(path)
	if err != nil {
		return nil, err
	}
	w := &wallet{
		currentId: currentId,
		threshold: threshold,
		info:      &common.WalletInfo{},
		fileStore: fs,
	}

	// Call Handler for keygen
	confs, err := mpc.Keygen(crypto.PartyID(currentId), threshold, []crypto.PartyID{"default", "vault"})
	if err != nil {
		return nil, err
	}

	_, err = w.fileStore.WriteCmpConfigs(crypto.SONRCoinType, confs)
	if err != nil {
		return nil, err
	}
	return w, nil
}

// Controller returns the controller of the wallet as did string
func (w *wallet) Controller() string {
	accs, err := w.fileStore.ListAccountsForToken(crypto.SONRCoinType)
	if err != nil {
		return ""
	}
	if len(accs) == 0 {
		return ""
	}
	return accs[0].DID()
}

// CreateAccount creates a new account for the given coin type
func (w *wallet) CreateAccount(coin crypto.CoinType) (Account, error) {
	// Call Handler for keygen
	confs, err := mpc.Keygen(crypto.PartyID(w.currentId), w.threshold, []crypto.PartyID{"default", "vault"})
	if err != nil {
		return nil, err
	}
	acc, err := w.fileStore.WriteCmpConfigs(coin, confs)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

// ListCoins returns a list of coins that this currently has accounts for
func (w *wallet) ListCoins() ([]crypto.CoinType, error) {
	accs, err := w.fileStore.ListAccounts()
	if err != nil {
		return nil, err
	}
	var coins []crypto.CoinType
	for coin := range accs {
		coins = append(coins, coin)
	}
	return coins, nil
}

// ListAccounts returns a list of accounts for the given coin type
func (w *wallet) ListAccounts() (map[crypto.CoinType][]Account, error) {
	return w.fileStore.ListAccounts()
}

// ListAccountsForCoin returns a list of accounts for the given coin type
func (w *wallet) ListAccountsForCoin(coin crypto.CoinType) ([]Account, error) {
	return w.fileStore.ListAccountsForToken(coin)
}

// GetAccount returns the account for the given coin type and account name
func (w *wallet) GetAccount(coin crypto.CoinType, name string) (Account, error) {
	accsList, err := w.fileStore.ListAccountsForToken(coin)
	if err != nil {
		return nil, err
	}
	for _, acc := range accsList {
		if acc.Name() == name {
			return acc, nil
		}
	}
	return nil, fmt.Errorf("account %s not found", name)
}

// GetAccountByAddress returns the account for the given address and parses the coin type from the address
func (w *wallet) GetAccountByAddress(address string) (Account, error) {
	coin := findCoinTypeFromAddress(address)
	return w.GetAccount(coin, address)
}

// GetAccountByDID returns the account for the given DID and parses the coin type from the DID
func (w *wallet) GetAccountByDID(did string) (Account, error) {
	addr, coin, _ := parseBlockchainAccountFromDID(did)
	return w.GetAccount(coin, addr)
}

// GetAccountByPublicKey returns the account for the given public key and parses the coin type from the public key
func (w *wallet) GetAccountByPublicKey(key string) (Account, error) {
	coinAccs, err := w.ListAccounts()
	if err != nil {
		return nil, err
	}
	pk, err := crypto.PubKeyFromBytes([]byte(key))
	if err != nil {
		return nil, err
	}

	for _, accs := range coinAccs {
		for _, acc := range accs {
			if acc.PubKey().Equals(pk) {
				return acc, nil
			}
		}
	}
	return nil, fmt.Errorf("account for public key %s not found", key)
}

// RenameAccount renames the account for the given coin type and account name
func (w *wallet) RenameAccount(coin crypto.CoinType, name, newName string) error {
	acc, err := w.GetAccount(coin, name)
	if err != nil {
		return err
	}
	return acc.Rename(newName)
}

// findCoinTypeFromAddress returns the CoinType for the given address
func findCoinTypeFromAddress(addr string) crypto.CoinType {
	for _, ct := range crypto.AllCoinTypes() {
		if strings.Contains(addr, ct.AddrPrefix()) {
			return ct
		}
	}
	return crypto.TestCoinType
}

// parseBlockchainAccountFromDID returns the blockchain account for the given DID returns as Address, CoinType, Account Name
func parseBlockchainAccountFromDID(did string) (string, crypto.CoinType, string) {
	// Split the DID into its constituent parts
	parts := strings.Split(did, ":")

	// If the DID is for a Sonr account, there will only be two parts: "did" and the address
	if len(parts) == 3 && parts[1] == "sonr" {
		return parts[2], crypto.CoinTypeFromDidMethod(parts[1]), ""
	}

	// If the DID is for a non-Sonr account, there will be three parts: "did", the coin type, and the address
	if len(parts) == 4 && parts[1] != "sonr" {
		// Split the account identifier into its constituent parts
		accountParts := strings.Split(parts[3], "#")
		if len(accountParts) == 1 {
			return parts[2], crypto.CoinTypeFromDidMethod(parts[1]), ""
		}
		return parts[2], crypto.CoinTypeFromDidMethod(parts[1]), accountParts[1]
	}

	// If the DID is not in the expected format, return empty strings
	return "", crypto.TestCoinType, ""
}
