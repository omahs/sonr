package wallet

import (
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
	if err := w.WalletConfig().BackupAccounts(d.n.LoadKeyValueStore); err != nil {
		return nil, err
	}
	return w, nil
}
