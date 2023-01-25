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

// snr16nzrp4x3sachmraq34uzr9tpzpp5tegcjam80z
// snr1qd3q2qfrax99264gcwts8jhentkttv7cgnl23k44u0w2j5n74cyqyxukmmh
