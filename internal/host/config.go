package host

import (
	"fmt"
	"time"

	"net"
	"os"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/multiformats/go-multiaddr"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/textileio/go-threads/core/thread"
)

// ** ─── Constants ────────────────────────────────────────────────────────
// Bootstrap Peer Discovery Interval
const refreshInterval = time.Second * 5

// Global Service Protocol ID
const globalProtocol = protocol.ID("/sonr/global-service/0.1")

// Textile Client API URL
const textileApiUrl = "api.hub.textile.io:443"

// Textile Miner Index Target
const textileMinerIdx = "f022352"

// ** ─── HostNode Info/Status Methods ────────────────────────────────────────────────────────
// @ Close Libp2p Host
func (h *hostNode) Close() {
	h.host.Close()
}

// @ Return Host Peer ID
func (hn *hostNode) ID() peer.ID {
	return hn.id
}

// @ Returns HostNode Peer Addr Info
func (hn *hostNode) Info() peer.AddrInfo {
	peerInfo := peer.AddrInfo{
		ID:    hn.host.ID(),
		Addrs: hn.host.Addrs(),
	}
	return peerInfo
}

// @ Returns Instance Host
func (hn *hostNode) Host() host.Host {
	return hn.host
}

// @ Returns Host Node MultiAddr
func (hn *hostNode) MultiAddr() (multiaddr.Multiaddr, *md.SonrError) {
	pi := hn.Info()
	addrs, err := peer.AddrInfoToP2pAddrs(&pi)
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_HOST_INFO)
	}
	return addrs[0], nil
}

// @ Returns Instance Host
func (hn *hostNode) PubKey() thread.PubKey {
	return hn.tileIdentity.GetPublic()
}

// ** ─── Address MANAGEMENT ────────────────────────────────────────────────────────
// # Return Bootstrap List Address Info
func getBootstrapAddrInfo() ([]peer.AddrInfo, error) {
	// Create Bootstrapper List
	var bootstrappers []ma.Multiaddr
	for _, s := range []string{
		// Libp2p Default
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
		"/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
		"/ip4/104.131.131.82/udp/4001/quic/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	} {
		ma, err := ma.NewMultiaddr(s)
		if err != nil {
			return nil, err
		}
		bootstrappers = append(bootstrappers, ma)
	}

	ds := make([]peer.AddrInfo, 0, len(bootstrappers))
	for i := range bootstrappers {
		info, err := peer.AddrInfoFromP2pAddr(bootstrappers[i])
		if err != nil {
			continue
		}
		ds = append(ds, *info)
	}
	return ds, nil
}

// # getFreePort asks the kernel for a free open port
func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

// # Return Internal Addr Strings ^ //
func getInternalAddrStrings() []string {
	// Initialize
	p, _ := getFreePort()
	listenAddrs := []string{
		fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", p),
	}
	return listenAddrs
}

// # Return Device Listening Addresses ^ //
func getExternalAddrStrings() ([]string, error) {
	// Initialize
	listenAddrs := []string{}
	hasIpv4 := false
	hasIpv6 := false

	// Get iPv4 Addresses
	ip4Addrs, err := iPv4Addrs()
	if err == nil {
		hasIpv4 = true
	}

	// Add iPv4 Addresses
	if hasIpv4 {
		listenAddrs = append(listenAddrs, ip4Addrs...)
	}

	// Neither iPv6 nor iPv4 found
	if !hasIpv4 && !hasIpv6 {
		return nil, errors.New("No IP Addresses found")
	}

	// Return Listen Addr Strings
	return listenAddrs, nil
}

// # Returns Node Public iPv4 Address
func iPv4Addrs() ([]string, error) {
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)

	p, _ := getFreePort()

	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv4 := addr.To4(); ipv4 != nil {
			ip4 := ipv4.String()
			return []string{
				fmt.Sprintf("/ip4/%s/tcp/%d", ip4, p),
			}, nil

		}
	}
	return nil, errors.New("No IPV4 found")
}
