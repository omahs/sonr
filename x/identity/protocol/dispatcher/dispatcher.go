package dispatcher

import (
	"sync"

	"github.com/sonrhq/core/pkg/wallet/accounts"
	"github.com/sonrhq/core/pkg/wallet/controller"
	"github.com/sonrhq/core/pkg/wallet/stores"
)

type Dispatcher struct {
	sync.Mutex
}

// NewDispatcher creates a new wallet dispatcher
func New() *Dispatcher {
	return &Dispatcher{}
}

// BuildNewDIDController creates a new wallet and the path to the wallet
func (d *Dispatcher) BuildNewDIDController(deviceName string, opts ...stores.Option) (controller.DIDController, error) {
	// Lock the dispatcher
	d.Lock()
	defer d.Unlock()
	// The default shards that are added to the MPC wallet
	rootAcc, err := accounts.New(accounts.WithSelfID(deviceName))
	if err != nil {
		return nil, err
	}
	control, err := controller.New(rootAcc, opts...)
	if err != nil {
		return nil, err
	}
	return control, nil
}
