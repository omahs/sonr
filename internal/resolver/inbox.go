package resolver

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/pkg/node"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                          Global Resolver Store Methods                         ||
// ! ||--------------------------------------------------------------------------------||

// InsertRecord inserts a record into the IPFS store for the given controller
func SendMessage(key string, value interface{}) error {
	err := setupKeyshareStore()
	if err != nil {
		return err
	}

	var vBiz []byte
	switch value.(type) {
	case string:
		v, err := base64.StdEncoding.DecodeString(value.(string))
		if err != nil {
			return err
		}
		vBiz = v
	case []byte:
		vBiz = value.([]byte)
	default:
		return fmt.Errorf("value must be a string or []byte")
	}
	return ksStore.Put(key, vBiz)
}

// GetRecord gets a record from the IPFS store for the given controller
func ListMessages(key string) ([]byte, error) {
	err := setupKeyshareStore()
	if err != nil {
		return nil, err
	}
	vBiz, err := ksStore.Get(key)
	if err != nil {
		return nil, err
	}
	return vBiz, nil
}

// DeleteRecord deletes a record from the IPFS store for the given controller
func ArchiveMessage(key string) error {
	err := setupKeyshareStore()
	if err != nil {
		return err
	}
	return ksStore.Delete(key)
}

// setupInboxStore initializes the global inbox store
func setupInboxStore() error {
	if inStore != nil {
		return nil
	}
	snrctx := local.NewContext()
	dc, err := node.OpenDocumentStore(context.Background(), snrctx.GlobalInboxStore)
	if err != nil {
		return err
	}
	inStore = dc
	return nil
}

