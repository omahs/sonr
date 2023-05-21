package handler

import (
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/x/identity/client/gateway/middleware"
)

func CreateAccount(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	acc, err := usr.CreateAccount(c.Params("name"), crypto.CoinTypeFromName(c.Params("coin_type")))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success":   true,
		"account":   acc.ToProto(),
		"coin_type": acc.CoinType().Ticker(),
		"address":   acc.Address(),
	})
}

func SignWithAccount(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	bz, err := base64.RawStdEncoding.DecodeString(c.Query("message"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	sig, err := usr.Sign(c.Params("address"), bz)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success":   true,
		"signature": base64.RawStdEncoding.EncodeToString(sig),
		"message":   c.Query("message"),
		"address":   c.Params("address"),
	})
}

func VerifyWithAccount(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	bz, err := base64.RawStdEncoding.DecodeString(c.Query("message"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	sig, err := base64.RawStdEncoding.DecodeString(c.Query("signature"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	ok, err := usr.Verify(c.Params("address"), bz, sig)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success":   true,
		"verified":  ok,
		"message":   c.Query("message"),
		"signature": c.Query("signature"),
		"address":   c.Params("address"),
	})
}
