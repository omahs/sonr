package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/internal/protocol/middleware"
	"github.com/sonrhq/core/x/identity"
)

func GetService(c *fiber.Ctx) error {
	q := middleware.ParseQuery(c)
	service, err := q.GetService()
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success": true,
		"service": service,
	})
}

func ListServices(c *fiber.Ctx) error {
	serviceList, err := local.Context().GetAllServices(c.Context())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success":  true,
		"services": serviceList,
	})
}

func GetServiceAttestion(c *fiber.Ctx) error {
	q := middleware.ParseQuery(c)
	service, err := q.GetService()
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	ucw, err := local.Context().OldestUnclaimedWallet(c.Context())
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	wc := identity.LoadClaimableWallet(ucw)
	chal, err := wc.IssueChallenge()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	challenge, err := service.GetCredentialCreationOptions(q.Alias(), chal, q.IsMobile())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"alias":             q.Alias(),
		"attestion_options": challenge,
		"origin":            q.Origin(),
		"challenge":         string(chal),
	})

}

func VerifyServiceAttestion(c *fiber.Ctx) error {
	q := middleware.ParseQuery(c)
	if !q.HasAttestion() {
		return c.Status(400).SendString("Missing attestion.")
	}

	// Get the origin from the request.
	service, err := q.GetService()
	if err != nil {
		return c.SendStatus(fiber.ErrNotFound.Code)
	}

	ucw, err := local.Context().OldestUnclaimedWallet(c.Context())
	if err != nil {
		return c.Status(404).SendString(fmt.Sprintf("Failed to find unclaimed wallet: %s", err.Error()))
	}

	claims := identity.LoadClaimableWallet(ucw)
	chal, err := claims.IssueChallenge()
	if err != nil {
		return c.Status(500).SendString(fmt.Sprintf("Failed to issue challenge: %s", err.Error()))
	}

	// Checking if the credential response is valid.
	cred, err := service.VerifyCreationChallenge(q.Attestion(), chal)
	if err != nil {
		return c.Status(403).SendString(fmt.Sprintf("Failed to verify attestion: %s", err.Error()))
	}

	cont, err := claims.Assign(cred, q.Alias())
	if err != nil {
		return c.Status(500).SendString(fmt.Sprintf("Failed to assign credential: %s", err.Error()))
	}

	usr := middleware.NewUser(cont, q.Alias())
	jwt, err := usr.JWT()
	if err != nil {
		return c.Status(500).SendString(fmt.Sprintf("Failed to create JWT: %s", err.Error()))
	}
	return c.JSON(fiber.Map{
		"success": true,
		"did":     cont.Did(),
		"primary": cont.PrimaryIdentity(),
		"tx_hash": cont.PrimaryTxHash(),
		"jwt":     jwt,
		"address": cont.Address(),
	})
}

func GetServiceAssertion(c *fiber.Ctx) error {
	q := middleware.ParseQuery(c)
	doc, err := q.GetDID()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	service, err := q.GetService()
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	challenge, err := service.GetCredentialAssertionOptions(doc, q.IsMobile())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"did":               doc.Id,
		"assertion_options": challenge,
		"origin":            q.Origin(),
	})
}

func VerifyServiceAssertion(c *fiber.Ctx) error {
	origin := c.Params("origin", "localhost")
	did := c.Query("did")
	assertion := c.Query("assertion")

	doc, err := local.Context().GetDID(c.Context(), did)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	service, err := local.Context().GetService(c.Context(), origin)
	if err != nil {
		return err
	}

	if err := service.VerifyAssertionChallenge(assertion, doc.ListCredentialVerificationMethods()...); err != nil {
		return c.Status(403).SendString(err.Error())
	}

	cont, err := identity.LoadController(doc)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	usr := middleware.NewUser(cont, doc.FindUsername())
	// Create the Claims
	jwt, err := usr.JWT()
	if err != nil {
		return c.Status(401).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success": true,
		"did":     cont.Did(),
		"jwt":     jwt,
		"address": cont.Address(),
	})
}
