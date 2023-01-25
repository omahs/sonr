package v1

import (
	context "context"
	"strings"

	"berty.tech/go-orbit-db/iface"
)

// `NewWalletConfigFromRootAccount` creates a new wallet config from a root account
func NewWalletConfigFromRootAccount(account *AccountConfig) *WalletConfig {
	return &WalletConfig{
		Address:   account.Address,
		Algorithm: "cmp",
		Accounts: map[string]*AccountConfig{
			strings.ToLower(account.Name): account,
		},
	}
}

// LoadKVStore is a function that loads a key value store from an address
type LoadKVStore func(address string) (iface.KeyValueStore, error)

// BackupAccounts is a function that backups accounts to a key value store
func (w *WalletConfig) BackupAccounts(l LoadKVStore) error {
	kv, err := l(w.Address)
	if err != nil {
		return err
	}

	for _, acc := range w.GetAccounts() {
		bz, err := acc.Marshal()
		if err != nil {
			return err
		}
		_, err = kv.Put(context.Background(), acc.DID(), bz)
		if err != nil {
			return err
		}
	}
	return nil
}
