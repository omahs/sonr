package api

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
	req := &v1.KeygenRequest{}
	err := req.Unmarshal(c.Request().Body())
	if err != nil {
		return err
	}

	// Get the origin from the request.
	service, err := resolver.GetService(context.Background(), req.Origin)
	if err != nil {
		return err
	}
	// Generate the keypair.
		cred, err := service.VerifyCreationChallenge(req.GetCredentialResponse())
	if err != nil {
		return err
	}

	cont, err := controller.NewController(context.Background(), cred)
	if err != nil {
		return err
	}
	res := &v1.KeygenResponse{
		Success: true,
		Did:     cont.DidDocument().Id,
		DidDocument: cont.DidDocument(),
	}
	return c.JSON(res)
}

func Login(c *fiber.Ctx) error {
	req := &v1.LoginRequest{}
	err := req.Unmarshal(c.Request().Body())
	if err != nil {
		return err
	}
	// Get the origin from the request.
	_, err = resolver.GetService(context.Background(), req.Origin)
	if err != nil {
		return err
	}

	return nil
}

func QueryDocument(c *fiber.Ctx) error {
	did := c.Params("did")

	// Get the origin from the request.
	doc, err := resolver.GetDID(context.Background(), did)
	if err != nil {
		return err
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
		return err
	}
	challenge, err := service.IssueChallenge()
	if err != nil {
		return err
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
	req := &v1.AddShareRequest{}
	err := req.Unmarshal(c.Request().Body())
	if err != nil {
		return err
	}
	err = resolver.InsertRecord(req.Key, req.Value)
	if err != nil {
		return err
	}
	return c.JSON(&v1.AddShareResponse{
		Success: true,
	})
}

func SyncShare(c *fiber.Ctx) error {
	req := &v1.SyncShareRequest{}
	err := req.Unmarshal(c.Request().Body())
	if err != nil {
		return err
	}
	record,err := resolver.GetRecord(req.Key)
	if err != nil {
		return err
	}
	return c.JSON(&v1.SyncShareResponse{
		Success: true,
		Value: base64.StdEncoding.EncodeToString(record),
		Key: req.Key,
	})
}

func RefreshShare(c *fiber.Ctx) error {
	return nil
}
