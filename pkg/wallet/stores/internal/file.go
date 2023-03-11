package internal

import (
	"context"
	"time"

	"github.com/sonrhq/core/pkg/wallet"
	"github.com/sonrhq/core/pkg/wallet/accounts"
	vaultv1 "github.com/sonrhq/core/x/identity/types/vault/v1"
	"github.com/ucan-wg/go-ucan"
)

type FileStore struct {
	sonrAcc  *vaultv1.AccountConfig
	basePath string
}

func NewFileStore(p string, accCfg *vaultv1.AccountConfig) (wallet.Store, error) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	ds := &FileStore{
		sonrAcc:  accCfg,
		basePath: p,
	}
	acc, err := accounts.Load(accCfg)
	if err != nil {
		return nil, err
	}
	err = ds.PutAccount(acc)
	if err != nil {
		return nil, err
	}
	return ds, nil
}

func (ds *FileStore) ListAccounts() ([]wallet.Account, error) {
	return ListAccounts()
}

func (ds *FileStore) PutAccount(w wallet.Account) error {
	accPath, err := getAccountPath(ds.basePath, w.CoinType())
	if err != nil {
		return err
	}

	WriteAccountConfig(accPath, ds.basePath, w.Config())

	return nil
}

// JWKClaims returns the JWKClaims for the store to be signed by the identity
func (ds *FileStore) JWKClaims(acc wallet.Account) (string, error) {
	caps := ucan.NewNestedCapabilities("DELEGATOR", "AUTHENTICATOR", "CREATE", "READ", "UPDATE")
	att := ucan.Attenuations{
		{Cap: caps.Cap("AUTHENTICATOR"), Rsc: ucan.NewStringLengthResource("mpc/acc", "*")},
		{Cap: caps.Cap("SUPER_USER"), Rsc: ucan.NewStringLengthResource("mpc/acc", "b5:world_bank_population:*")},
	}
	zero := time.Time{}
	origin, err := acc.NewOriginToken(acc.PubKey().DID(), att, nil, zero, zero)
	if err != nil {
		return "", err
	}
	return origin, nil
}

// VerifyJWKClaims verifies the JWKClaims for the store
func (ds *FileStore) VerifyJWKClaims(claims string, acc wallet.Account) error {
	p := exampleParser()
	_, err := p.ParseAndVerify(context.Background(), claims)
	if err != nil {
		return err
	}
	return nil
}
