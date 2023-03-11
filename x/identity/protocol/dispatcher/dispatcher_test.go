package dispatcher_test

import (
	"encoding/base64"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/wallet/stores"
	"github.com/sonrhq/core/x/identity/protocol/dispatcher"
	"github.com/stretchr/testify/assert"
)

type accTest struct {
	name     string
	coinType crypto.CoinType
}

var accTests = []accTest{
	{"Ethereum", crypto.ETHCoinType},
	{"Bitcoin", crypto.BTCCoinType},
	{"Filecoin", crypto.FILCoinType},
	{"Handshake", crypto.HNSCoinType},
	{"Litecoin", crypto.LTCCoinType},
	{"Cosmos", crypto.COSMOSCoinType},
	{"Sonr", crypto.SONRCoinType},
}

func usedNetworks() string {
	var nets []string
	for _, test := range accTests {
		nets = append(nets, test.coinType.Ticker())
	}
	return strings.Join(nets, ", ")
}

func TestNewWallet(t *testing.T) {
	t.Logf("Initialize new DID Controller...")
	startTime := time.Now()
	d := dispatcher.New()
	homeDir, err := os.UserHomeDir()
	checkErr(t, err)
	w, err := d.BuildNewDIDController("prad's iphone", stores.SetFileStorePath(homeDir, "Desktop"))
	checkErr(t, err)
	t.Logf("(%s) - Root Identifier", w.ID())
	t.Logf("Address: %s", w.Address())
	t.Logf("DONE! Time elapsed: %s\n", time.Since(startTime))

	un := usedNetworks()
	t.Logf("\nCreate %s accounts....", un)
	startTime = time.Now()
	for i, test := range accTests {
		vm, err := w.CreateAccount(test.name, test.coinType)
		checkErr(t, err)
		t.Logf("\t* [%d] #%s", i, vm.IDFragmentSuffix())
		t.Logf("\t\t↪ Address: %s", vm.BlockchainAccountId)
		t.Logf("\t\t↪ Controller: %s", vm.Controller)
		t.Logf("\t\t↪ Type: %s", vm.Type)
		t.Logf("\t\t↪ Multibase PubKey: %s", vm.PublicKeyMultibase)
	}
	t.Logf("DONE! Time elapsed: %s\n", time.Since(startTime))

	t.Logf("\nGet %s accounts....", un)
	accs, err := w.ListAccounts()
	t.Logf("DONE! Time elapsed: %s\n", time.Since(startTime))
	checkErr(t, err)
	for i, acc := range accs {
		t.Logf("\t* [%d] #%s", i, acc.DID())
		t.Logf("\t\t↪ Address: %s", acc.Address())
		t.Logf("\t\t↪ Controller: %s", acc.CoinType())
		t.Logf("\t\t↪ Type: %s", acc.Type())
		t.Logf("\t\t↪ Multibase PubKey: %s", acc.PubKey())
	}
}

func TestListAccounts(t *testing.T) {
	// wallet.GetAccountsByCoinType(walletPath string, coinType crypto.CoinType)
}

func TestDispatcherSignature(t *testing.T) {
	t.Logf("Initialize new DID Controller...")
	startTime := time.Now()
	d := dispatcher.New()

	w, err := d.BuildNewDIDController("prad's iphone")
	checkErr(t, err)
	t.Logf("(%s) - Root Identifier", w.ID())
	t.Logf("Address: %s", w.Address())
	t.Logf("DONE! Time elapsed: %s\n", time.Since(startTime))

	msg := []byte("Hello World!")
	t.Logf("\nSign msg with DID Controller....")
	startTime = time.Now()
	sig, err := w.SignWithAccount(msg, "primary")
	checkErr(t, err)
	t.Logf("\t↪ %s", base64.RawStdEncoding.EncodeToString(sig))
	t.Logf("DONE! Time elapsed: %s\n", time.Since(startTime))

	t.Logf("\nVerify msg with DID Controller....")
	startTime = time.Now()
	ok, err := w.VerifyWithAccount(msg, sig, "primary")
	checkErr(t, err)
	assert.True(t, ok)
	t.Logf("DONE! Time elapsed: %s\n", time.Since(startTime))
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
