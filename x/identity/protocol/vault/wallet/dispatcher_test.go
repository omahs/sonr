package wallet

import (
	"testing"

	"github.com/sonrhq/core/pkg/common"
)

func TestDispatcher(t *testing.T) {
	d := NewDispatcher()
	w, err := d.CallNewWallet()
	checkErr(t, err)
	t.Log(w.Address())
	err = w.CreateAccount("Ethereum", common.CoinType_CoinType_ETHEREUM)
	checkErr(t, err)
	err = w.CreateAccount("Bitcoin", common.CoinType_CoinType_BITCOIN)
	checkErr(t, err)
	accs, err := w.ListAccounts()
	checkErr(t, err)
	for _, acc := range accs {
		t.Log(acc.AccountConfig())
	}
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

// snr16nzrp4x3sachmraq34uzr9tpzpp5tegcjam80z
// snr1qd3q2qfrax99264gcwts8jhentkttv7cgnl23k44u0w2j5n74cyqyxukmmh
