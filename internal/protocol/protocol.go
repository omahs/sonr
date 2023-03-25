package protocol

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/helmet/v2"
	"github.com/sonrhq/core/internal/local"
)

var hway *Protocol

type Protocol struct {
	ctx client.Context
}

func RegisterHighway(ctx client.Context) {
	setupFiber(ctx)
}

func setupFiber(ctx client.Context) {
	// Setup the fiber app.
	hway = &Protocol{ctx: ctx}
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Middleware
	app.Use(cors.New())
	app.Use(helmet.New())

	// Status Methods
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK. Highway version v0.6.0. Running on HTTP/TLS")
	})

	// Query Methods
	app.Get("/highway/query/service/:origin", timeout.New(QueryService, time.Second*5))
	app.Get("/highway/query/document/:did", timeout.New(QueryDocument, time.Second*5))

	// Auth Methods
	app.Post("/highway/auth/keygen", timeout.New(Keygen, time.Second*10))
	app.Post("/highway/auth/login", timeout.New(Login, time.Second*10))
	app.Post("/highway/vault/add", timeout.New(AddShare, time.Second*5))
	app.Post("/highway/vault/sync", timeout.New(SyncShare, time.Second*5))
	go hway.serveFiber(app)
}

func (p *Protocol) serveFiber(app *fiber.App) {
	snrctx := local.NewContext()
	if snrctx.HasTlsCert() {
		app.ListenTLS(
			fmt.Sprintf(":%s", snrctx.HighwayPort()),
			snrctx.TlsCertPath,
			snrctx.TlsKeyPath,
		)
	} else {
		if snrctx.IsDev() {

			app.Listen(
				fmt.Sprintf(":%s", snrctx.HighwayPort()),
			)
		} else {
			app.Listen(
				fmt.Sprintf("%s:%s", snrctx.GrpcEndpoint(), snrctx.HighwayPort()),
			)
		}
	}
}
