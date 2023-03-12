package v2

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/sonrhq/core/pkg/crypto"
)

func TestNewWallet(t *testing.T) {
	w, err := NewWallet("test", 1)
	if err != nil {
		t.Fatal(err)
	}
	if w == nil {
		t.Fatal("wallet is nil")
	}
}

func TestCreateAccount(t *testing.T) {
	w, err := NewWallet("test", 1)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range defaultCoinTestsSet() {
		_, err := w.CreateAccount(tt.coinType)
		if err != nil {
			t.Fatal(err)
		}

		accs, err := w.ListAccountsForCoin(tt.coinType)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("<%s> %s (%d)", tt.coinType.Ticker(), tt.coinType.Name(), tt.coinType.BipPath())
		for i, acc := range accs {
			t.Logf("- [%d] %s", i, acc.Name())
			t.Logf(" \t↪ Address: %s", acc.Address())
			t.Logf(" \t↪ PubKey: %s", acc.PubKey().Multibase())
		}
		t.Logf("")
	}
}

func TestGetAccount(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(homeDir, "Desktop", "_SONR_WALLET_")

	w, err := LoadWallet(path)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range defaultAccountTestsSet() {
		acc, err := w.GetAccount(tt.coinType, tt.index)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("<%s> %s (%d)", tt.coinType.Ticker(), tt.coinType.Name(), tt.coinType.BipPath())
		t.Logf("- [%d] %s", tt.index, acc.DID())
		t.Logf(" \t↪ Address: %s", acc.Address())
		t.Logf(" \t↪ PubKey: %s", acc.PubKey().Multibase())
		t.Logf("")
	}
}

func TestSignWithAccount(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(homeDir, "Desktop", "_SONR_WALLET_")

	w, err := LoadWallet(path)
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range defaultAccountTestsSet() {
		acc, err := w.GetAccount(tt.coinType, tt.index)
		if err != nil {
			t.Fatal(err)
		}
		msg := []byte(fmt.Sprintf("Hello %s!", tt.coinType.Name()))
		t.Logf("- [%d] %s - %s", tt.index, acc.Name(), acc.Address())
		t.Logf(" \t↪ Message: %s", string(msg))
		sig, err := acc.Sign(msg)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf(" \t↪ Signature: %s", base64.StdEncoding.EncodeToString(sig))

		ok, err := acc.Verify(msg, sig)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("   Verify => %v", ok)
		t.Logf("")
	}
}

func defaultCoinTestsSet() []struct {
	coinType crypto.CoinType
} {
	return []struct {
		coinType crypto.CoinType
	}{
		{crypto.BTCCoinType},
		{crypto.ETHCoinType},
		{crypto.SONRCoinType},
		{crypto.FILCoinType},
		{crypto.DOGECoinType},
	}
}

func defaultAccountTestsSet() []struct {
	coinType crypto.CoinType
	index    int
} {
	return []struct {
		coinType crypto.CoinType
		index    int
	}{
		{crypto.BTCCoinType, 0},
		{crypto.ETHCoinType, 0},
		{crypto.SONRCoinType, 0},
		{crypto.FILCoinType, 0},
		{crypto.DOGECoinType, 0},
	}
}
