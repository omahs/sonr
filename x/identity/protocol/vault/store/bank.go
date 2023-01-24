package store

import (
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sonrhq/core/pkg/node/config"
	"github.com/sonrhq/core/x/identity/protocol/vault/wallet"
	"github.com/sonrhq/core/x/identity/types"
)

type VaultBank struct {
	// The IPFS node that the vault is running on
	node config.IPFSNode

	// The wallet that the vault is using
	cache *gocache.Cache

	// The TxBuilder
	cctx client.Context
}

// Creates a new Vault
func InitBank(cctx client.Context, node config.IPFSNode, cache *gocache.Cache) *VaultBank {
	return &VaultBank{
		node:  node,
		cache: cache,
		cctx:  cctx,
	}
}

func (v *VaultBank) StartRegistration(entry *Session) (string, string, error) {
	optsJson, err := entry.BeginRegistration()
	if err != nil {
		return "", "", err
	}
	v.putEntryIntoCache(entry)
	return optsJson, entry.ID, nil
}

func (v *VaultBank) FinishRegistration(sessionId string, credsJson string) (*types.DidDocument, wallet.Wallet, error) {
	// Get Session
	entry, err := v.getEntryFromCache(sessionId)
	if err != nil {
		return nil, nil, err
	}
	didDoc, err := entry.FinishRegistration(credsJson)
	if err != nil {
		return nil, nil, err
	}
	// Create a new offline wallet
	wallet, err := wallet.NewWallet(v.cctx)
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("Failed to create new offline wallet using MPC: %s", err))
	}

	primAcc, err := wallet.PrimaryAccount()
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("Failed to get primary account: %s", err))
	}
	didDoc.AddAssertion(primAcc.GetAssertionMethod())
	return didDoc, wallet, nil
}

func (v *VaultBank) StartLogin(entry *Session) (string, string, error) {
	optsJson, err := entry.BeginLogin()
	if err != nil {
		return "", "", err
	}
	v.putEntryIntoCache(entry)
	return optsJson, entry.ID, nil
}

func (v *VaultBank) FinishLogin(sessionId string, credsJson string) (bool, error) {
	// Get Session
	entry, err := v.getEntryFromCache(sessionId)
	if err != nil {
		return false, err
	}
	didDoc, err := entry.FinishLogin(credsJson)
	if err != nil {
		return false, err
	}
	return didDoc, nil
}

func (v *VaultBank) getEntryFromCache(id string) (*Session, error) {
	val, ok := v.cache.Get(id)
	if !ok {
		return nil, errors.New("Failed to find entry for ID")
	}
	e, ok := val.(*Session)
	if !ok {
		return nil, errors.New("Invalid type for session entry")
	}
	return e, nil
}

func (v *VaultBank) putEntryIntoCache(entry *Session) error {
	if entry == nil {
		return errors.New("Entry cannot be nil to put into cache")
	}
	return v.cache.Add(entry.ID, entry, -1)
}
