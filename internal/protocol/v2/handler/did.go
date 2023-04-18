package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/local"
)

func GetDID(c *fiber.Ctx) error {
	did := c.Params("did")

	// Get the origin from the request.
	doc, err := local.Context().GetDID(context.Background(), did)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"did":      did,
		"document": doc,
	})
}

func GetDIDByAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	available, doc, err := local.Context().CheckAlias(context.Background(), alias)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	if doc == nil {
		return c.JSON(fiber.Map{
			"available": available,
		})
	}

		return c.JSON(fiber.Map{
			"available": available,
			"did":       doc.Id,
			"document":  doc,
		})

}

func ListDIDs(c *fiber.Ctx) error {
	docs, err := local.Context().GetAllDIDs(context.Background())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"documents": docs,
	})
}


