package wallet

import (
	"context"
	"sync"

	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/x/identity/controller"
	"github.com/sonrhq/core/x/identity/protocol/vault/account"
)

type Dispatcher struct {
	n common.IPFSNode
	sync.Mutex
}

// NewDispatcher creates a new wallet dispatcher
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		// n: n,
	}
}

// CallNewWallet creates a new wallet
func (d *Dispatcher) CallNewWallet() (controller.DIDController, error) {
	// Lock the dispatcher
	d.Lock()
	defer d.Unlock()
	doneCh := make(chan controller.DIDController)
	errCh := make(chan error)

	// Create the wallet in a goroutine
	go func() {
		// The default shards that are added to the MPC wallet
		rootAcc, err := account.NewAccount("Primary", common.CoinType_CoinType_SONR)
		if err != nil {
			errCh <- err
		}
		control, err := controller.New(context.Background(), rootAcc.AccountConfig())
		if err != nil {
			errCh <- err
		}
		doneCh <- control
	}()

	// Wait for the wallet to be created
	select {
	case w := <-doneCh:
		return w, nil
	case err := <-errCh:
		return nil, err
	}
}
