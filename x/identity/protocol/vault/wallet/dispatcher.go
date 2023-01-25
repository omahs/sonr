package wallet

import (
	"context"
	"fmt"

	"github.com/sonrhq/core/pkg/node/config"
)

type Dispatcher struct {
	n config.IPFSNode
}

func NewDispatcher(n config.IPFSNode) *Dispatcher {
	return &Dispatcher{
		n: n,
	}
}

func (d *Dispatcher) CallNewWallet() (Wallet, error) {
	doneCh := make(chan Wallet)
	errCh := make(chan error)

	go func() {
		w, err := newWallet(d.n)
		if err != nil {
			errCh <- err
		}
		doneCh <- w
	}()

	select {
	case w := <-doneCh:
		return d.backupWallet(w)
	case err := <-errCh:
		return nil, err
	}
	// return NewWallet()
}

func (d *Dispatcher) backupWallet(w Wallet) (Wallet, error) {
	kv, err := d.n.LoadKeyValueStore(w.WalletConfig().Address)
	if err != nil {
		return nil, err
	}

	bz, err := w.WalletConfig().Marshal()
	if err != nil {
		return nil, err
	}

	op, err := kv.Put(context.Background(), w.WalletConfig().Address, bz)
	if err != nil {
		return nil, err
	}
	fmt.Println(op)
	return w, nil
}
