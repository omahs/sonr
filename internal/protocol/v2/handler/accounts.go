package handler

import (
	"encoding/base64"
	v1 "github.com/sonrhq/core/types/highway/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/protocol/v2/middleware"
	"github.com/sonrhq/core/x/identity/controller"
)

func GetAccount(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
		primeID, err := usr.PrimaryIdentity()
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	cont, err := controller.LoadController(primeID)
	if err != nil {
		return c.Status(401).SendString(err.Error())
	}
	address := c.Params("address", "")
	acc, err := cont.GetAccount(address)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"success": true,
		"account": acc.ToProto(),
		"coin_type": acc.CoinType().Ticker(),
		"name": acc.Name(),
		},
	)
}

func ListAccounts(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	primeID, err := usr.PrimaryIdentity()
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	cont, err := controller.LoadController(primeID)
	if err != nil {
		return c.Status(401).SendString(err.Error())
	}
	accs, err := cont.ListAccounts()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var accsList []interface{}
	for _, acc := range accs {
		accsList = append(accsList, fiber.Map{
			"account": acc.ToProto(),
			"coin_type": acc.CoinType().Ticker(),
			"name": acc.Name(),
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"accounts": accsList,
		},)
}

func CreateAccount(c *fiber.Ctx) error {
	coinType := c.Params("coin_type")
	name := c.Params("name")
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	primeID, err := usr.PrimaryIdentity()
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	cont, err := controller.LoadController(primeID)
	if err != nil {
		return c.Status(401).SendString(err.Error())
	}
	ct := crypto.CoinTypeFromName(coinType)
	acc, err := cont.CreateAccount(name, ct)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success": true,
		"account": acc.ToProto(),
		"coin_type": acc.CoinType().Ticker(),
		"name": acc.Name(),
		},
	)
}

func SignWithAccount(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	req := new(v1.SignMessageRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	primeID, err := usr.PrimaryIdentity()
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	cont, err := controller.LoadController(primeID)
	if err != nil {
		return c.Status(401).SendString(err.Error())
	}
	bz, err := base64.RawStdEncoding.DecodeString(req.Message)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	sig, err := cont.Sign(req.Did, bz)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success": true,
		"signature": base64.RawStdEncoding.EncodeToString(sig),
		"message": req.Message,
		"did": req.Did,
	})
}

func VerifyWithAccount(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	req := new(v1.VerifyMessageRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	bz, err := base64.RawStdEncoding.DecodeString(req.Message)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	sig, err := base64.RawStdEncoding.DecodeString(req.Signature)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	primeID, err := usr.PrimaryIdentity()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	cont, err := controller.LoadController(primeID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	ok, err := cont.Verify(req.Did, bz, sig)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success": true,
		"verified": ok,
		"message": req.Message,
		"signature": req.Signature,
		"did": req.Did,
	})
}

func SendTransaction(c *fiber.Ctx) error {
	return nil
}
