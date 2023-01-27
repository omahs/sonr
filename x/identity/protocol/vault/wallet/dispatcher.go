package wallet

import (
	"sync"

	"github.com/sonrhq/core/pkg/common"
)

type Dispatcher struct {
	n common.IPFSNode
	sync.Mutex
}

// NewDispatcher creates a new wallet dispatcher
func NewDispatcher(n common.IPFSNode) *Dispatcher {
	return &Dispatcher{
		n: n,
	}
}

// CallNewWallet creates a new wallet
func (d *Dispatcher) CallNewWallet() (Wallet, error) {
	// Lock the dispatcher
	d.Lock()
	defer d.Unlock()
	doneCh := make(chan Wallet)
	errCh := make(chan error)

	// Create the wallet in a goroutine
	go func() {
		w, err := newWallet(d.n)
		if err != nil {
			errCh <- err
		}
		doneCh <- w
	}()

	// Wait for the wallet to be created
	select {
	case w := <-doneCh:
		if err := w.WalletConfig().BackupAccounts(d.n.LoadKeyValueStore); err != nil {
			return nil, err
		}
		return w, nil
	case err := <-errCh:
		return nil, err
	}
}
