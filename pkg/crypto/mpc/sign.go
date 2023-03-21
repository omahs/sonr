package mpc

import (
	"fmt"
	"sync"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/crypto/mpc/internal/algorithm"
	"github.com/sonrhq/core/pkg/crypto/mpc/internal/network"

	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// SerializeECDSASecp256k1Signature marshals an ECDSA signature to DER format for use with the CMP protocol
func SerializeECDSASecp256k1Signature(sig *crypto.MPCECDSASignature) ([]byte, error) {
	rBytes, err := sig.R.MarshalBinary()
	if err != nil {
		return nil, err
	}
	sBytes, err := sig.S.MarshalBinary()
	if err != nil {
		return nil, err
	}

	sigBytes := make([]byte, 65)
	// 0 pad the byte arrays from the left if they aren't big enough.
	copy(sigBytes[33-len(rBytes):33], rBytes)
	copy(sigBytes[65-len(sBytes):65], sBytes)
	return sigBytes, nil
}

// ConfigFunc is a function type that returns a *cmp.Config.
type ConfigFunc func() *cmp.Config

// SignCMP signs a message with the given private key using the CMP protocol.
func SignCMP(configProviders []interface{}, m []byte) ([]byte, error) {
    configs := make([]*cmp.Config, len(configProviders))
    for i, cp := range configProviders {
        confProvider, ok := cp.(ConfigFunc)
        if !ok {
            return nil, fmt.Errorf("invalid config provider at index %d", i)
        }
        configs[i] = confProvider()
    }

    peers := make([]crypto.PartyID, len(configs))
    for i, c := range configs {
        peers[i] = c.ID
    }
    net := network.NewOfflineNetwork(peers...)
    doneChan := make(chan *crypto.MPCECDSASignature, 1)
    var wg sync.WaitGroup
    for _, c := range configs {
        wg.Add(1)
        go func(conf *cmp.Config) {
            pl := crypto.NewMPCPool(0)
            defer pl.TearDown()
            sig, err := algorithm.CmpSign(conf, m, net.Ls(), net, &wg, pl)
            if err != nil {
                return
            }
            doneChan <- sig
        }(c)
    }
    wg.Wait()
    sig := <-doneChan
    return SerializeECDSASecp256k1Signature(sig)
}
