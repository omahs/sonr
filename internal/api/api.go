package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

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
	if isFiber() {
		setupFiber(ctx)

	} else {
		setupConnect(ctx)
	}
}

func setupConnect(ctx client.Context) {
	hway = &Protocol{ctx: ctx}
	mux := http.NewServeMux()
	mux.Handle(AuthenticationHandler())
	mux.Handle(MpcHandler())
	mux.Handle(VaultHandler())
	go hway.serveConnect(mux)
}

func setupFiber(ctx client.Context) {
	hway = &Protocol{ctx: ctx}
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	app.Post("/highway/auth/keygen", Keygen)
	app.Post("/highway/auth/login", Login)
	app.Get("/highway/auth/service/:origin", QueryService)
	app.Get("/highway/auth/document/:did", QueryDocument)
	app.Post("/highway/vault/add", AddShare)
	app.Post("/highway/vault/sync", SyncShare)
	go hway.serveFiber(app)
}

func (p *Protocol) serveConnect(mux *http.ServeMux) {
	if hasTLSCert() {
		http.ListenAndServeTLS(
			fmt.Sprintf(":%s", getServerPort()),
			getTLSCert(),
			getTLSKey(),
			mux,
		)
	} else {
		http.ListenAndServe(
			fmt.Sprintf(":%s", getServerPort()),
			h2c.NewHandler(mux, &http2.Server{}),
		)
	}
}

func (p *Protocol) serveFiber(app *fiber.App) {
	if hasTLSCert() {
		app.ListenTLS(
			fmt.Sprintf(":%s", getServerPort()),
			getTLSCert(),
			getTLSKey(),
		)
	} else {
		app.Listen(
			fmt.Sprintf(":%s", getServerPort()),
		)
	}
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

func isFiber() bool {
	return os.Getenv("HIGHWAY_MODE") != "connect"
}
