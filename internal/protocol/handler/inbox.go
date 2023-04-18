package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/protocol/middleware"
	v1 "github.com/sonrhq/core/types/highway/v1"
	"github.com/sonrhq/core/x/identity/controller"
)

func ReadInboxMessages(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	req := new(v1.ReadMailRequest)
	if err := c.BodyParser(req); err != nil {
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

	msgs, err := cont.ReadMail(req.AccountAddress)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	fromBodyMap := make(map[string]string)
	for _, msg := range msgs {
		fromBodyMap[msg.Sender] = msg.Content
	}

	return c.JSON(fiber.Map{
		"success":  true,
		"messages": fromBodyMap,
	})
}

func SendInboxMessage(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	req := new(v1.SendMailRequest)
	if err := c.BodyParser(req); err != nil {
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

	err = cont.SendMail(req.FromAddress, req.ToAddress, req.Message)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
