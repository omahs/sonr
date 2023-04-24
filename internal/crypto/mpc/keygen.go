package mpc

import (
	"sync"

	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/crypto/mpc/algorithm"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// Keygen Generates a new ECDSA private key shared among all the given participants.
func Keygen(current crypto.PartyID, option ...KeygenOption) ([]*cmp.Config, error) {
	var mtx sync.Mutex
	var wg sync.WaitGroup
	opts := defaultKeygenOpts(current)
	opts.Apply(option...)

	net := opts.getOfflineNetwork()
	confs := make([]*cmp.Config, 0)
	for _, id := range net.Ls() {
		wg.Add(1)
		go func(id party.ID, network crypto.Network) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			mtx.Lock()

			conf, err := algorithm.CmpKeygen(id, net.Ls(), network, opts.Threshold, &wg, pl)
			opts.handleRoutineErr(err)
			opts.handleConfigGeneration(conf)
			confs = append(confs, conf)

			mtx.Unlock()
		}(id, net)
	}
	wg.Wait()
	return confs, nil
}
