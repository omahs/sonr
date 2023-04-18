package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/internal/protocol/v2/middleware"
	v1 "github.com/sonrhq/core/types/highway/v1"
	"github.com/sonrhq/core/x/identity/controller"
)

func GetService(c *fiber.Ctx) error {
	origin := c.Params("origin", "localhost")
	// Get the origin from the request.
	service, err := local.Context().GetService(context.Background(), origin)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"success": true,
		"service": service,
	},)
}

func ListServices(c *fiber.Ctx) error {
	serviceList, err := local.Context().GetAllServices(context.Background())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"success": true,
		"services": serviceList,
	},)
}

func GetServiceAssertion(c *fiber.Ctx) error {
	origin := c.Params("origin", "localhost")
	alias := c.Params("alias", "admin")

	_, doc, err := local.Context().CheckAlias(context.Background(), alias)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Get the origin from the request.
	service, err := local.Context().GetService(context.Background(), origin)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	challenge, err := service.GetCredentialAssertionOptions(doc)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	resp := &v1.QueryServiceAssertionResponse{
		AssertionOptions: challenge,
		Alias:            alias,
		Origin:           service.Origin,
	}
	return c.JSON(resp)
}

func VerifyServiceAssertion(c *fiber.Ctx) error {
	req := new(v1.KeygenRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	ok, _, err := local.Context().CheckAlias(context.Background(), req.Username)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if !ok {
		return c.Status(400).SendString("Username already taken.")
	}

	// Get the origin from the request.
	// uuid := req.Uuid
	service, _ := local.Context().GetService(context.Background(), req.Origin)
	if service == nil {
		// Try to get the service from the localhost
		service, _ = local.Context().GetService(context.Background(), "localhost")
	}

	// Check if service is still nil - return internal server error
	if service == nil {
		return c.Status(500).SendString("Internal Server Error.")
	}

	// Checking if the credential response is valid.
	cred, err := service.VerifyCreationChallenge(req.CredentialResponse)
	if err != nil {
		c.Status(400).SendString(err.Error())
	}

	// Create a new controller with the credential.
	cont, err := controller.NewController(controller.WithWebauthnCredential(cred), controller.WithBroadcastTx(), controller.WithUsername(req.Username))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	usr := middleware.NewUser(cont, req.Username)
	// Create the Claims
	jwt, err := usr.JWT()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	accs, err := usr.ListAccounts()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	res := &v1.KeygenResponse{
		Success:  true,
		Did:      cont.Did(),
		Primary:  cont.PrimaryIdentity(),
		Accounts: accs,
		TransactionHash: cont.PrimaryTxHash(),
		Jwt:      jwt,
	}
	return c.JSON(res)
}

func GetServiceAttestion(c *fiber.Ctx) error {
	origin := c.Params("origin", "localhost")
	alias := c.Params("alias", "admin")
	// Get the origin from the request.
	service, err := local.Context().GetService(context.Background(), origin)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	challenge, err := service.GetCredentialCreationOptions(alias)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	resp := &v1.QueryServiceResponse{
		CredentialOptions: challenge,
		RpName:            "Sonr",
		RpId:              service.Origin,
	}
	return c.JSON(resp)
}

func VerifyServiceAttestion(c *fiber.Ctx) error {
	req := new(v1.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	doc, err := local.Context().GetDID(c.Context(), req.AccountAddress)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if doc == nil && req.GetUsername() != "" {
		ok, ddoc, err := local.Context().CheckAlias(c.Context(), req.Username)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		if !ok {
			return c.Status(400).SendString("Username not found.")
		}
		doc = ddoc
	}

	cont, err := controller.LoadController(doc)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	usr := middleware.NewUser(cont, req.GetUsername())
	// Create the Claims
	jwt, err := usr.JWT()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	res := &v1.LoginResponse{
		Success: true,
		Did:     cont.Did(),
		Jwt:     jwt,
		Address: cont.Address(),
	}
	return c.JSON(res)
}
