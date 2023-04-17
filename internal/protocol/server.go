package protocol

import (
	"context"
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/internal/node"
	"github.com/sonrhq/core/types/crypto"
	v1 "github.com/sonrhq/core/types/highway/v1"
	"github.com/sonrhq/core/x/identity/controller"
)

func RegisterHighway(ctx client.Context) {
	app := initHttpTransport(ctx)
	node.StartLocalIPFS()
	go serveFiber(app.App)
}

func serveFiber(app *fiber.App) {
	if local.Context().HasTlsCert() {
		app.ListenTLS(
			local.Context().FiberListenAddress(),
			local.Context().TlsCertPath,
			local.Context().TlsKeyPath,
		)
	} else {
		app.Listen(local.Context().FiberListenAddress())
	}
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                          Auth API Rest Implementation                          ||
// ! ||--------------------------------------------------------------------------------||

func (htt *HttpTransport) Keygen(c *fiber.Ctx) error {
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
	usr := controller.NewUser(cont, req.Username)
	// Create the Claims
	jwt, err := usr.JWT([]byte("secret"))
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

func (htt *HttpTransport) Login(c *fiber.Ctx) error {
	req := new(v1.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	doc, err := local.Context().GetDID(c.Context(), req.AccountAddress)
	if err != nil {
		return err
	}

	if doc == nil && req.GetUsername() != "" {
		ok, ddoc, err := local.Context().CheckAlias(c.Context(), req.Username)
		if err != nil {
			return err
		}
		if !ok {
			return c.Status(400).SendString("Username not found.")
		}
		doc = ddoc
	}

	cont, err := controller.LoadController(doc)
	if err != nil {
		return err
	}
	usr := controller.NewUser(cont, req.GetUsername())
	// Create the Claims
	jwt, err := usr.JWT([]byte("secret"))
	if err != nil {
		return err
	}

	res := &v1.LoginResponse{
		Success: true,
		Did:     cont.Did(),
		Jwt:     jwt,
		Address: cont.Address(),
	}
	return c.JSON(res)
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                        Query Method for the Highway API                        ||
// ! ||--------------------------------------------------------------------------------||

func QueryDocument(c *fiber.Ctx) error {
	did := c.Params("did")

	// Get the origin from the request.
	doc, err := local.Context().GetDID(context.Background(), did)
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

func (htt *HttpTransport) QueryAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	available, doc, err := local.Context().CheckAlias(context.Background(), alias)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	if doc == nil {
		resp := &v1.QueryAliasResponse{
			Available: true,
		}
		return c.JSON(resp)
	}
	resp := &v1.QueryAliasResponse{
		DidDocument: doc,
		Did:         doc.Id,
		Available:   available,
	}
	return c.JSON(resp)
}

func (htt *HttpTransport) QueryDocument(c *fiber.Ctx) error {
	did := c.Params("did")
	// Get the origin from the request.
	doc, err := local.Context().GetDID(context.Background(), did)
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

func (htt *HttpTransport) QueryServiceCreation(c *fiber.Ctx) error {
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

func (htt *HttpTransport) QueryServiceAssertion(c *fiber.Ctx) error {
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

// ! ||--------------------------------------------------------------------------------||
// ! ||                        Accounts API Rest Implementation                        ||
// ! ||--------------------------------------------------------------------------------||
func (htt *HttpTransport) IsAuthorized(c *fiber.Ctx) error {
	usr, err := htt.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "User is authorized",
		"user":    usr,
	})
}

func (htt *HttpTransport) CreateAccount(c *fiber.Ctx) error {
	usr, err := htt.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	req := new(v1.CreateAccountRequest)
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
	ct := crypto.CoinTypeFromName(req.CoinType)
	acc, err := cont.CreateAccount(req.Name, ct)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	res := &v1.CreateAccountResponse{
		Success: true,
		Accounts: []*v1.Account{
			acc.ToProto(),
		},
	}
	return c.JSON(res)
}

func (htt *HttpTransport) ListAccounts(c *fiber.Ctx) error {
	usr, err := htt.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	primeID, err := usr.PrimaryIdentity()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	cont, err := controller.LoadController(primeID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	accs, err := cont.ListAccounts()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	res := &v1.ListAccountsResponse{
		Success: true,
	}
	for _, acc := range accs {
		res.Accounts = append(res.Accounts, acc.ToProto())
	}
	return c.JSON(res)
}

func (htt *HttpTransport) GetAccount(c *fiber.Ctx) error {
	usr, err := htt.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
		primeID, err := usr.PrimaryIdentity()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	cont, err := controller.LoadController(primeID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	address := c.Params("address", "")
	acc, err := cont.GetAccount(address)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	res := &v1.GetAccountResponse{
		Success:  true,
		CoinType: acc.CoinType().Name(),
		Accounts: []*v1.Account{
			acc.ToProto(),
		},
	}
	return c.JSON(res)
}

func (htt *HttpTransport) SignMessage(c *fiber.Ctx) error {
	usr, err := htt.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	req := new(v1.SignMessageRequest)
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
	bz, err := base64.RawStdEncoding.DecodeString(req.Message)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	sig, err := cont.Sign(req.Did, bz)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	res := &v1.SignMessageResponse{
		Success:   true,
		Signature: base64.RawStdEncoding.EncodeToString(sig),
		Message:   req.Message,
		Did:       req.Did,
	}
	return c.JSON(res)
}

func (htt *HttpTransport) VerifyMessage(c *fiber.Ctx) error {
	usr, err := htt.FetchUser(c)
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
	res := &v1.VerifyMessageResponse{
		Success: ok,
		Did:     req.Did,
	}
	return c.JSON(res)
}

func (htt *HttpTransport) SendMail(c *fiber.Ctx) error {
	usr, err := htt.FetchUser(c)
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
	res := &v1.SendMailResponse{
		Success: true,
	}
	return c.JSON(res)
}

func (htt *HttpTransport) ReadMail(c *fiber.Ctx) error {
	usr, err := htt.FetchUser(c)
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
	res := &v1.ReadMailResponse{
		Success:  true,
		Messages: fromBodyMap,
	}
	return c.JSON(res)
}
