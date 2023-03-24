package protocol

import (
	"context"
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/controller"
	"github.com/sonrhq/core/internal/resolver"
	v1 "github.com/sonrhq/core/types/highway/v1"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                          Auth API Rest Implementation                          ||
// ! ||--------------------------------------------------------------------------------||

func Keygen(c *fiber.Ctx) error {
	req := new(v1.KeygenRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Get the origin from the request.
	service, err := resolver.GetService(context.Background(), req.Origin)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	// Generate the keypair.
	cred, err := service.VerifyCreationChallenge(req.GetCredentialResponse())
	if err != nil {
		return c.Status(401).SendString(err.Error())
	}

	cont, err := controller.NewController(context.Background(), cred)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	res := &v1.KeygenResponse{
		Success: true,
		Did:     cont.DidDocument().Id,
		DidDocument: cont.DidDocument(),
	}
	return c.JSON(res)
}

func Login(c *fiber.Ctx) error {
	req := new(v1.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Get the origin from the request.
	_, err := resolver.GetService(context.Background(), req.Origin)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return nil
}

func QueryDocument(c *fiber.Ctx) error {
	did := c.Params("did")

	// Get the origin from the request.
	doc, err := resolver.GetDID(context.Background(), did)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	resp := &v1.QueryDocumentResponse{
		Success:        (doc != nil),
		AccountAddress: doc.DIDIdentifier(),
		DidDocument:    doc,
	}
	return c.JSON(resp)
}

func QueryService(c *fiber.Ctx) error {
	origin := c.Params("origin", "localhost")
	// Get the origin from the request.
	service, err := resolver.GetService(context.Background(), origin)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	challenge, err := service.IssueChallenge()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	resp := &v1.QueryServiceResponse{
		Challenge: string(challenge),
		RpName:    "Sonr",
		RpId:      service.Origin,
	}
	return c.JSON(resp)
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                        Accounts API Rest Implementation                        ||
// ! ||--------------------------------------------------------------------------------||
func CreateAccount(c *fiber.Ctx) error {
	return nil
}

func ListAccounts(c *fiber.Ctx) error {
	return nil
}

func GetAccount(c *fiber.Ctx) error {
	return nil
}

func DeleteAccount(c *fiber.Ctx) error {
	return nil
}

func SignMessage(c *fiber.Ctx) error {
	return nil
}

func VerifyMessage(c *fiber.Ctx) error {
	return nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                          Vault API Rest Implementation                         ||
// ! ||--------------------------------------------------------------------------------||

func AddShare(c *fiber.Ctx) error {
	req := new(v1.AddShareRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err := resolver.InsertRecord(req.Key, req.Value)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(&v1.AddShareResponse{
		Success: true,
		Key: req.Key,
	})
}

func SyncShare(c *fiber.Ctx) error {
	req := new(v1.SyncShareRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	record,err := resolver.GetRecord(req.Key)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(&v1.SyncShareResponse{
		Success: true,
		Value: base64.StdEncoding.EncodeToString(record),
		Key: req.Key,
	})
}

func RefreshShare(c *fiber.Ctx) error {
		return c.Status(500).JSON(fiber.Map{"error": "not implemented"})
}
