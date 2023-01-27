package controller

import (
	"github.com/sonrhq/core/x/identity/protocol/vault/store"
	"github.com/sonrhq/core/x/identity/protocol/vault/wallet"
)

type DIDController interface {
	wallet.Wallet
	store.WalletStore
}
