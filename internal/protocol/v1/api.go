package protocol

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/helmet/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/sonrhq/core/internal/local"
)

type HttpTransport struct {
	*fiber.App
	SessionStore  *session.Store
	ClientContext client.Context
}

func initHttpTransport(ctx client.Context) *HttpTransport {
	// HttPTransport
	rest := &HttpTransport{
		App: fiber.New(fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
		ClientContext: ctx,
		SessionStore:  session.New(),
	}

	// Middleware
	rest.Use(cors.New())
	rest.Use(helmet.New())

	// Query Methods
	rest.Get("/highway/query/alias/:alias", timeout.New(rest.QueryAlias, time.Second*5))
	rest.Get("/highway/query/document/:did", timeout.New(rest.QueryDocument, time.Second*5))

	// Service Methods
	rest.Get("/highway/service/:origin/attest/:alias", timeout.New(rest.QueryServiceCreation, time.Second*5))
	rest.Get("/highway/service/:origin/assert/:alias", timeout.New(rest.QueryServiceAssertion, time.Second*5))

	// Auth Methods
	rest.Post("/highway/auth/keygen", timeout.New(rest.Keygen, time.Second*10))
	rest.Post("/highway/auth/login", timeout.New(rest.Login, time.Second*10))

	// -- RESTRICTED METHODS --
	rest.Use(jwtware.New(jwtware.Config{
		SigningKey: local.Context().SigningKey(),
	}))

	// MPC Methods
	rest.Get("/highway/auth/check", timeout.New(rest.IsAuthorized, time.Second*5))
	rest.Get("/highway/accounts", timeout.New(rest.ListAccounts, time.Second*5))
	rest.Get("/highway/accounts/:address", timeout.New(rest.GetAccount, time.Second*5))
	rest.Post("/highway/accounts/create", timeout.New(rest.CreateAccount, time.Second*5))
	rest.Post("/highway/accounts/sign", timeout.New(rest.SignMessage, time.Second*5))
	rest.Post("/highway/accounts/verify", timeout.New(rest.VerifyMessage, time.Second*5))

	// Inbox Methods
	rest.Get("/highway/inbox/read", timeout.New(rest.ReadMail, time.Second*5))
	rest.Post("/highway/inbox/send", timeout.New(rest.SendMail, time.Second*5))
	return rest
}
