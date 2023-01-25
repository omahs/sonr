package wallet

import (
	"context"
	"testing"

	"github.com/sonrhq/core/pkg/node"
)

func TestDispatcher(t *testing.T) {
	n, err := node.NewIPFS(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	d := NewDispatcher(n)
	w, err := d.CallNewWallet()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(w.WalletConfig().Address)
}
