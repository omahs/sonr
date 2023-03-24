package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	highway "github.com/sonrhq/core/types/highway/v1/highwayv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var hway *Protocol

type Protocol struct {
	ctx client.Context
	highway.UnimplementedMpcHandler
	highway.UnimplementedAuthenticationHandler
	highway.UnimplementedVaultHandler
}

func RegisterHighway(ctx client.Context) {
	hway = &Protocol{ctx: ctx}
	mux := http.NewServeMux()
	mux.Handle(AuthenticationHandler())
	mux.Handle(MpcHandler())
	mux.Handle(VaultHandler())
	go hway.serveConnectHTTP(mux)
}

func (p *Protocol) serveConnectHTTP(mux *http.ServeMux) {
	if hasTLSCert() {
		http.ListenAndServeTLS(
			fmt.Sprintf(":%s", getServerPort()),
			getTLSCert(),
			getTLSKey(),
			mux,
		)
	} else {
		http.ListenAndServe(
			fmt.Sprintf("%s:%s", getServerHost(), getServerPort()),
			h2c.NewHandler(mux, &http2.Server{}),
		)
	}
}

func getServerHost() string {
	if host := os.Getenv("CONNECT_SERVER_ADDRESS"); host != "" {
		log.Printf("using CONNECT_SERVER_ADDRESS: %s", host)
		return host
	}
	return "localhost"
}

func getServerPort() string {
	if port := os.Getenv("CONNECT_SERVER_PORT"); port != "" {
		log.Printf("using CONNECT_SERVER_PORT: %s", port)
		return port
	}
	return "8080"
}

func getTLSCert() string {
	if cert := os.Getenv("CONNECT_SERVER_TLS_CERT"); cert != "" {
		log.Printf("using CONNECT_SERVER_TLS_CERT: %s", cert)
		return cert
	}
	return ""
}

func getTLSKey() string {
	if key := os.Getenv("CONNECT_SERVER_TLS_KEY"); key != "" {
		log.Printf("using CONNECT_SERVER_TLS_KEY: %s", key)
		return key
	}
	return ""
}

func hasTLSCert() bool {
	return getTLSCert() != "" && getTLSKey() != "" && !isDev()
}

func isDev() bool {
	return os.Getenv("ENVIRONMENT") == "dev"
}
