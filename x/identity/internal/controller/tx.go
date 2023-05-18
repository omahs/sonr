package controller

import (
	"github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/identity/types/models"
)

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
