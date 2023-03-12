package v2

import "testing"

func TestNewWallet(t *testing.T) {
	w, err := NewWallet("test", 1)
	if err != nil {
		t.Fatal(err)
	}
	if w == nil {
		t.Fatal("wallet is nil")
	}

}
