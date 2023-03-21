package protocol

import (
	"testing"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/crypto/mpc"
	"github.com/tendermint/tendermint/libs/rand"
)

func TestTree(t *testing.T) {
	randUuid := rand.Str(4)

	// Call Handler for keygen
	confs, err := mpc.Keygen(crypto.PartyID(randUuid), 1, []crypto.PartyID{"vault"})
	if err != nil {
		t.Fatal(err)
	}

	var kss []KeyShare
	for _, conf := range confs {
		ksb, err := conf.MarshalBinary()
		if err != nil {
			t.Fatal(err)
		}
		ks, err := NewKeyshare(string(conf.ID), ksb, crypto.SONRCoinType, "test")
		if err != nil {
			t.Fatal(err)
		}
		kss = append(kss, ks)
	}

	for _, ks := range kss {
		t.Logf("keyshare path: %s", ks.Bip44())
		t.Logf("keyshare coin type: %s", ks.CoinType())
		t.Logf("keyshare account name: %s", ks.AccountName())
		t.Logf("keyshare key id: %s", ks.KeyID())
	}
	configProviders := []interface{}{}
	for _, ks := range kss {
		configProviders = append(configProviders, ks)
	}
}
