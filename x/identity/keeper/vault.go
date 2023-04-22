package keeper

import (
	"context"
	"errors"
	"strings"

	"berty.tech/go-orbit-db/iface"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/internal/node"
	"github.com/sonrhq/core/x/identity/types/models"
	servicetypes "github.com/sonrhq/core/x/service/types"
)

var (
	v *vaultImpl
)

type Vault interface {
	// InsertAccount inserts the account and its keyshares into the IPFS store
	InsertAccount(acc models.Account) error

	// GetAccount gets the account and its keyshares from the IPFS store
	GetAccount(accDid string) (models.Account, error)

	// DeleteAccount deletes the account and its keyshares from the IPFS store
	DeleteAccount(accDid string) error

	// ReadInbox reads the inbox from the IPFS store
	ReadInbox(accDid string) ([]*models.InboxMessage, error)

	// WriteInbox writes the inbox to the IPFS store
	WriteInbox(toDid string, msg *models.InboxMessage) error
}

type vaultImpl struct {
	KsTable node.IPFSKVStore
	InTable node.IPFSDocsStore

	ctx context.Context
}

func setupVault() error {
	if v != nil {
		return nil
	}
	ctx := context.Background()
	snrctx := local.Context()
	kv, err := node.OpenKeyValueStore(ctx, snrctx.GlobalKvKsStore)
	if err != nil {
		return err
	}
	docs, err := node.OpenDocumentStore(ctx, snrctx.GlobalInboxDocsStore, nil)
	if err != nil {
		return err
	}
	vi := &vaultImpl{
		KsTable: kv,
		InTable: docs,
		ctx:     ctx,
	}
	v = vi
	return nil
}

func InsertAccount(acc models.Account) error {
	err := setupVault()
	if err != nil {
		return err
	}
	ksAccListVal := strings.Join(acc.ListKeyShares(), ",")
	_, err = v.KsTable.Put(v.ctx, accountPrefix(acc.Did()), []byte(ksAccListVal))
	if err != nil {
		return err
	}
	acc.MapKeyShare(func(ks models.KeyShare) models.KeyShare {
		_, err = v.KsTable.Put(v.ctx, keysharePrefix(ks.Did()), ks.Bytes())
		if err != nil {
			return nil
		}
		return ks
	})
	return nil
}

func GetAccount(accDid string) (models.Account, error) {
	err := setupVault()
	if err != nil {
		return nil, err
	}
	ksr, err := models.ParseAccountDID(accDid)
	if err != nil {
		return nil, err
	}

	vBiz, err := v.KsTable.Get(context.Background(), accountPrefix(accDid))
	if err != nil {
		return nil, err
	}

	ksAccListVal := strings.Split(string(vBiz), ",")
	var ksList []models.KeyShare
	for _, ksDid := range ksAccListVal {
		vBiz, err := v.KsTable.Get(context.Background(), keysharePrefix(ksDid))
		if err != nil {
			return nil, err
		}
		ks, err := models.NewKeyshare(ksDid, vBiz, ksr.CoinType, ksr.AccountName)
		if err != nil {
			return nil, err
		}
		ksList = append(ksList, ks)
	}
	acc := models.NewAccount(ksList, ksr.CoinType)
	return acc, nil
}

func DeleteAccount(accDid string) error {
	err := setupVault()
	if err != nil {
		return err
	}
	// Delete the keyshares
	vBiz, err := v.KsTable.Get(context.Background(), accountPrefix(accDid))
	if err != nil {
		return err
	}

	ksAccListVal := strings.Split(string(vBiz), ",")
	for _, ksDid := range ksAccListVal {
		_, err = v.KsTable.Delete(context.Background(), keysharePrefix(ksDid))
		if err != nil {
			return err
		}
	}

	// Delete the account
	_, err = v.KsTable.Delete(context.Background(), accountPrefix(accDid))
	if err != nil {
		return err
	}
	return nil
}

func FetchCredential(keyDid string) (servicetypes.Credential, error) {
	err := setupVault()
	if err != nil {
		return nil, err
	}
	// Delete the keyshares
	vBiz, err := v.KsTable.Get(context.Background(), webauthnPrefix(keyDid))
	if err != nil {
		return nil, err
	}
	return servicetypes.LoadJSONCredential(vBiz)
}

func StoreCredential(cred servicetypes.Credential) error {
	err := setupVault()
	if err != nil {
		return err
	}
	bz, err := cred.Marshal()
	if err != nil {
		return err
	}
	_, err = v.KsTable.Put(v.ctx, webauthnPrefix(cred.Did()), bz)
	if err != nil {
		return err
	}
	return nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                         Inbox handler for W2W messages                         ||
// ! ||--------------------------------------------------------------------------------||

// CreateInbox sets up a default inbox for the account
func (v *vaultImpl) CreateInbox(accDid string) error {
	inbox, err := models.CreateDefaultInboxMap(accDid)
	if err != nil {
		return err
	}
	_, err = v.InTable.Put(v.ctx, inbox)
	if err != nil {
		return err
	}
	return nil
}

// HasInbox checks if the account has an inbox
func (v *vaultImpl) HasInbox(accDid string) (bool, error) {
	inboxRaw, err := v.InTable.Get(v.ctx, accDid, &iface.DocumentStoreGetOptions{})
	if err != nil {
		return false, err
	}
	if len(inboxRaw) == 0 {
		return false, nil
	}
	return true, nil
}

// LoadInbox loads the inbox for the account
func (v *vaultImpl) LoadInbox(accDid string) (*models.Inbox, error) {
	// Check if the inbox exists
	hasInbox, err := v.HasInbox(accDid)
	if err != nil {
		return nil, err
	}
	if !hasInbox {
		err := v.CreateInbox(accDid)
		if err != nil {
			return nil, err
		}
	}

	// Load the inbox
	inboxRaw, err := v.InTable.Get(v.ctx, accDid, &iface.DocumentStoreGetOptions{})
	inboxMap, ok := inboxRaw[0].(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid inbox")
	}
	inbox, err := models.NewInboxFromMap(inboxMap)
	if err != nil {
		return nil, err
	}
	return inbox, nil
}

// ReadInbox reads the inbox for the account
func ReadInbox(accDid string) ([]*models.InboxMessage, error) {
	inbox, err := v.LoadInbox(accDid)
	if err != nil {
		return nil, err
	}
	return inbox.Messages, nil
}

// WriteInbox writes a message to the inbox for the account
func WriteInbox(toDid string, msg *models.InboxMessage) error {
	// Get the inbox
	inbox, err := v.LoadInbox(toDid)
	if err != nil {
		return err
	}
	// Add the message to the inbox
	inboxMap, err := inbox.AddMessageToMap(msg)
	if err != nil {
		return err
	}
	// Update the inbox
	_, err = v.InTable.Put(v.ctx, inboxMap)
	if err != nil {
		return err
	}
	return nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                         Helper Methods for Module Setup                        ||
// ! ||--------------------------------------------------------------------------------||

func keysharePrefix(v string) string {
	return "ks/" + v
}

func accountPrefix(v string) string {
	return "acc/" + v
}

func webauthnPrefix(v string) string {
	return "webauthn/" + v
}
