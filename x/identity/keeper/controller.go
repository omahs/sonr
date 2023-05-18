package keeper

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	// "github.com/sonrhq/core/internal/vault"
	// "github.com/sonrhq/core/internal/vault"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/crypto/mpc"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/internal/vault"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/identity/types/models"
	servicetypes "github.com/sonrhq/core/x/service/types"
)

var PrimaryAccountaddress string = "primary"

type Controller interface {
	// The `Address()` function is a method of the `didController` struct that returns the address of the
	// primary account associated with the controller. It takes a pointer to the `didController` struct as
	// its receiver and returns a string representing the address of the primary account.
	Address() string

	// The `Did()` function is a method of the `didController` struct that returns the DID (Decentralized
	// Identifier) associated with the controller's primary account. It takes a pointer to the
	// `didController` struct as its receiver and returns a string representing the DID.
	Did() string

	// PrimaryIdentity returns the controller's DID document
	GetIdentity() *types.Identity

	// BlockchainIdentities returns the controller's blockchain identities
	BlockchainIdentities() []*types.DidDocument

	// Createmodels.Account creates a new models.Account for the controller
	CreateAccount(name string, coinType crypto.CoinType) (models.Account, error)

	// GetAccount returns an account by Address or DID
	GetAccount(id string) (models.Account, error)

	// Listmodels.Accounts returns the controller's models.Accounts
	ListAccounts() ([]models.Account, error)

	// SendMail sends a message between two Controllers
	SendMail(address string, to string, body string) error

	// ReadMail reads the controller's inbox
	ReadMail(address string) ([]*models.InboxMessage, error)

	// Sign signs a message with the controller's models.Account
	Sign(address string, msg []byte) ([]byte, error)

	// Verify verifies a signature with the controller's models.Account
	Verify(address string, msg []byte, sig []byte) (bool, error)
}

type didController struct {
	primary    models.Account
	primaryDoc *types.DidDocument
	identity   *types.Identity
	blockchain []models.Account

	currCredential *servicetypes.WebauthnCredential
	disableIPFS    bool
	aka            string
	broadcastChan  chan *local.BroadcastTxResponse
}

func NewController(options ...ControllerOption) (Controller, error) {
	opts := defaultOptions()
	for _, option := range options {
		option(opts)
	}

	doneCh := make(chan models.Account)
	errCh := make(chan error)
	go generateInitialAccount(context.Background(), opts.WebauthnCredential, doneCh, errCh, opts)

	select {
	case acc := <-doneCh:
		cn, err := setupController(context.Background(), acc, opts)
		if err != nil {
			return nil, err
		}
		return cn, nil
	case err := <-errCh:
		return nil, err
	}
}

// The function loads a controller with a primary account and a list of blockchain accounts from a
// given identity.
func LoadController(doc *types.Identity) (Controller, error) {
	acc, err := vault.GetAccount(doc.Id)
	if err != nil {
		return nil, err
	}
	blockAccDids := doc.CapabilityDelegation
	var blockAccs []models.Account
	for _, did := range blockAccDids {
		acc, err := vault.GetAccount(did)
		if err != nil {
			return nil, err
		}
		blockAccs = append(blockAccs, acc)
	}
	cn := &didController{
		primary:    acc,
		identity:   doc,
		blockchain: blockAccs,
	}
	return cn, nil
}

// The function loads a controller with a primary account and a list of blockchain accounts from a
// given DID document.
func LoadControllerWithDid(doc *types.DidDocument) (Controller, error) {
	acc, err := vault.GetAccount(doc.Id)
	if err != nil {
		return nil, err
	}
	blockAccDids := doc.CapabilityDelegation
	var blockAccs []models.Account
	for _, did := range blockAccDids {
		acc, err := vault.GetAccount(did)
		if err != nil {
			return nil, err
		}
		blockAccs = append(blockAccs, acc)
	}
	cn := &didController{
		primary:    acc,
		primaryDoc: doc,
		blockchain: blockAccs,
	}
	return cn, nil
}

// The `Address()` function is a method of the `didController` struct that returns the address of the
// primary account associated with the controller. It takes a pointer to the `didController` struct as
// its receiver and returns a string representing the address of the primary account.
func (dc *didController) Address() string {
	return dc.primary.Address()
}

// The `Did()` function is a method of the `didController` struct that returns the DID (Decentralized
// Identifier) associated with the controller's primary account. It takes a pointer to the
// `didController` struct as its receiver and returns a string representing the DID.
func (dc *didController) Did() string {
	return dc.primaryDoc.Id
}

// The `PrimaryIdentity()` function is a method of the `didController` struct that returns the DID
// document associated with the controller's primary account. It takes a pointer to the `didController`
// struct as its receiver and returns a pointer to a `types.DidDocument` representing the primary
// account's DID document.
func (dc *didController) GetIdentity() *types.Identity {
	return dc.identity
}

// The `BlockchainIdentities()` function is a method of the `didController` struct that returns an
// array of `*types.DidDocument` representing the DID documents of all the blockchain identities
// associated with the controller. It takes a pointer to the `didController` struct as its receiver and
// returns an array of pointers to `types.DidDocument`.
func (dc *didController) BlockchainIdentities() []*types.DidDocument {
	var docs []*types.DidDocument
	for _, acc := range dc.blockchain {
		fmt.Println(acc)
	}
	return docs
}

// Returns a list of all the accounts associated with the controller. It
// returns an array of `models.Account` and an error. The method first checks if the primary account
// exists and then appends it to the list of blockchain accounts associated with the controller.
// Finally, it returns the list of accounts.
func (dc *didController) ListAccounts() ([]models.Account, error) {
	if dc.primary == nil {
		return nil, fmt.Errorf("no Primary Account found")
	}
	return append([]models.Account{dc.primary}, dc.blockchain...), nil
}

func (dc *didController) CreateAccount(name string, coinType crypto.CoinType) (models.Account, error) {
	ctCount := 0
	for _, acc := range dc.blockchain {
		if acc.CoinType() == coinType {
			ctCount++
		}
	}
	newAcc, err := dc.primary.DeriveAccount(coinType, ctCount, name)
	if err != nil {
		return nil, err
	}

	// Add account to the vault
	if !dc.disableIPFS {
		err = vault.InsertAccount(newAcc)
		if err != nil {
			return nil, err
		}
	}

	// Add the new models.Account to the controller
	dc.blockchain = append(dc.blockchain, newAcc)
	return newAcc, nil
}

// Getmodels.Account returns the controller's models.Account from the Address
func (dc *didController) GetAccount(address string) (models.Account, error) {
	if strings.Contains(address, "did:") {
		return dc.GetAccountByDid(address)
	}
	for _, acc := range dc.blockchain {
		if acc.Address() == address {
			return acc, nil
		}
	}
	return nil, fmt.Errorf("models.Account not found")
}

// GetAccountByDid returns the controller's models.Account from the DID
func (dc *didController) GetAccountByDid(did string) (models.Account, error) {
	if dc.primaryDoc.Id == did {
		return dc.primary, nil
	}
	for _, acc := range dc.blockchain {
		if acc.Did() == did {
			return acc, nil
		}
	}
	return nil, fmt.Errorf("models.Account not found")
}

// Sign signs a message with the controller's selected models.Account
func (dc *didController) Sign(address string, msg []byte) ([]byte, error) {
	acc, err := dc.GetAccount(address)
	if err != nil {
		return nil, err
	}
	return acc.Sign(msg)
}

// Verify verifies a signature with the controller's selected models.Account
func (dc *didController) Verify(address string, msg []byte, sig []byte) (bool, error) {
	acc, err := dc.GetAccount(address)
	if err != nil {

		return false, err
	}
	return acc.Verify(msg, sig)
}

// SendMail sends a mail from the controller's selected models.Account
func (dc *didController) SendMail(address string, to string, body string) error {
	acc, err := dc.GetAccount(address)
	if err != nil {
		return err
	}
	msg, err := acc.CreateInboxMessage(to, body)
	if err != nil {
		return err
	}
	err = vault.WriteInbox(to, msg)
	if err != nil {
		return err
	}
	return nil
}

// ReadMail reads a mail from the controller's selected models.Account
func (dc *didController) ReadMail(address string) ([]*models.InboxMessage, error) {
	acc, err := dc.GetAccount(address)
	if err != nil {
		return nil, err
	}
	return vault.ReadInbox(acc.Address())
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                         Blockchain Transaction Methods                         ||
// ! ||--------------------------------------------------------------------------------||
// CreatePrimaryIdentity sends a transaction to create a new DID document with the provided account
func (c *didController) CreatePrimaryIdentity(doc *types.DidDocument, acc models.Account, alias string, wallet_id uint32) {
	go func() {
		msg := types.NewMsgCreateDidDocument(acc.Address(), wallet_id, alias, doc)
		resp, err := c.primary.SendSonrTx(msg)
		if err != nil {
			return
		}
		c.broadcastChan <- resp
	}()
}

// UpdatePrimaryIdentity sends a transaction to update an existing DID document with the provided account
func (c *didController) UpdatePrimaryIdentity(docs ...*types.DidDocument) {
	go func() {
		msg := types.NewMsgUpdateDidDocument(c.primary.Address(), c.primaryDoc, docs...)
		resp, err := c.primary.SendSonrTx(msg)
		if err != nil {
			return
		}
		c.broadcastChan <- resp
	}()
}

// RegisterIdentity sends a transaction to register a new identity with the provided account
func (c *didController) RegisterIdentity(id *types.Identity, alias string, wallet_id uint32, relationships ...*types.VerificationRelationship) {
	go func() {
		msg := types.NewMsgRegisterIdentity(c.primary.Address(), wallet_id, alias, id, relationships...)
		resp, err := c.primary.SendSonrTx(msg)
		if err != nil {
			return
		}
		c.broadcastChan <- resp
	}()
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                  Configuration                                 ||
// ! ||--------------------------------------------------------------------------------||

type ControllerOptions struct {
	// The controller's on config generated handler
	OnConfigGenerated []mpc.OnConfigGenerated

	// Credential to authorize the controller
	WebauthnCredential *servicetypes.WebauthnCredential

	// Disable IPFS
	DisableIPFS bool

	// Broadcast the transaction
	BroadcastTx bool

	// Username for the controller
	Username string

	errChan       chan error
	broadcastChan chan *local.BroadcastTxResponse
}

func defaultOptions() *ControllerOptions {
	return &ControllerOptions{
		OnConfigGenerated: []mpc.OnConfigGenerated{},
		DisableIPFS:       false,
		BroadcastTx:       false,
		Username:          "",
		errChan:           make(chan error),
		broadcastChan:     make(chan *local.BroadcastTxResponse),
	}
}

type ControllerOption func(*ControllerOptions)

func WithUsername(username string) ControllerOption {
	return func(o *ControllerOptions) {
		o.Username = username
	}
}

func WithConfigHandlers(handlers ...mpc.OnConfigGenerated) ControllerOption {
	return func(o *ControllerOptions) {
		o.OnConfigGenerated = handlers
	}
}

func WithWebauthnCredential(cred *servicetypes.WebauthnCredential) ControllerOption {
	return func(o *ControllerOptions) {
		o.WebauthnCredential = cred
	}
}

func WithIPFSDisabled() ControllerOption {
	return func(o *ControllerOptions) {
		o.DisableIPFS = true
	}
}

func WithBroadcastTx(brdcastChan chan *local.BroadcastTxResponse) ControllerOption {
	return func(o *ControllerOptions) {
		o.BroadcastTx = true
		o.broadcastChan = brdcastChan
	}
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                          Helper Methods for Controller                         ||
// ! ||--------------------------------------------------------------------------------||

func generateInitialAccount(ctx context.Context, credential *servicetypes.WebauthnCredential, doneCh chan models.Account, errChan chan error, opts *ControllerOptions) {
	shardName := crypto.PartyID(base64.RawStdEncoding.EncodeToString(credential.Id))
	// Call Handler for keygen
	confs, err := mpc.Keygen(shardName, mpc.WithHandlers(opts.OnConfigGenerated...))
	if err != nil {
		errChan <- err
	}

	pubKey, err := crypto.NewPubKeyFromCmpConfig(confs[0])
	if err != nil {
		errChan <- err
	}

	rootDid, _ := crypto.SONRCoinType.FormatDID(pubKey)
	var kss []models.KeyShare
	for _, conf := range confs {
		ksb, err := conf.MarshalBinary()
		if err != nil {
			errChan <- err
		}
		ksDid := fmt.Sprintf("%s#%s", rootDid, conf.ID)
		ks, err := models.NewKeyshare(ksDid, ksb, crypto.SONRCoinType)
		if err != nil {
			errChan <- err
		}
		kss = append(kss, ks)
	}
	doneCh <- models.NewAccount(kss, crypto.SONRCoinType)
}

func setupController(ctx context.Context, primary models.Account, opts *ControllerOptions) (Controller, error) {
	if !opts.DisableIPFS {
		err := vault.InsertAccount(primary)
		if err != nil {
			return nil, err
		}
	}

	doc := types.NewPrimaryIdentity(primary.Did(), primary.PubKey(), nil)
	if opts.WebauthnCredential != nil {
		cred, err := servicetypes.ValidateWebauthnCredential(opts.WebauthnCredential, primary.Did())
		if err != nil {
			return nil, err
		}
		doc = types.NewPrimaryIdentity(primary.Did(), primary.PubKey(), cred.ToVerificationMethod())
		if !opts.DisableIPFS {
			err = vault.StoreCredential(cred)
			if err != nil {
				return nil, err
			}
		}
	}

	if opts.Username != "" {
		doc.AlsoKnownAs = []string{opts.Username}
	}

	cont := &didController{
		primary:     primary,
		blockchain:  []models.Account{},
		primaryDoc:  doc,
		disableIPFS: opts.DisableIPFS,
		aka:         doc.FindUsername(),
	}
	return cont, nil
}
