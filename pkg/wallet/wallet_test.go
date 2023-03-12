package v2

import (
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

	var tests = []struct {
		coinType crypto.CoinType
	}{
		{crypto.BTCCoinType},
		{crypto.ETHCoinType},
		{crypto.LTCCoinType},
		{crypto.SOLCoinType},
	}

	for _, tt := range tests {
		_, err := w.CreateAccount(tt.coinType)
		if err != nil {
			t.Fatal(err)
		}

		accs, err := w.ListAccountsForCoin(tt.coinType)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("<%s>", tt.coinType.Ticker())
		for i, acc := range accs {
			t.Logf("\t[%d] - %s", i, acc.Name())
			t.Logf("\t\t↪ Address: %s", acc.Address())
			t.Logf("\tt PubKey: %s", acc.PubKey())
			t.Logf("\t\t↪ Path: %s", acc.Path())
		}
	}
}
