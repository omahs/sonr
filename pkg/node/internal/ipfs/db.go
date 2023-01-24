package ipfs

import (
	"context"

	orbitdb "berty.tech/go-orbit-db"
	"berty.tech/go-orbit-db/iface"
	"github.com/sonrhq/core/pkg/node/config"
)

type dbImpl struct {
	// IPFS URL of the remote node
	ctx     context.Context
	node    config.IPFSNode
	orbitDb iface.OrbitDB
}

// Initialize creates a new node with default options
func newDBInstance(ctx context.Context, n config.IPFSNode) (config.IPFSDB, error) {
	db, err := orbitdb.NewOrbitDB(ctx, n.CoreAPI(), &orbitdb.NewOrbitDBOptions{})
	if err != nil {
		return nil, err
	}
	return &dbImpl{
		ctx:     ctx,
		node:    n,
		orbitDb: db,
	}, nil
}

// GetDocsStore creates or loads a document database from given name
func (r *dbImpl) GetDocsStore(username string) (iface.DocumentStore, error) {
	addr, err := r.fetchDocsAddress(username)
	if err != nil {
		return nil, err
	}
	return r.orbitDb.Docs(r.ctx, addr, nil)
}

// GetEventLogStore creates or loads an event log database from given name
func (r *dbImpl) GetEventLogStore(username string) (iface.EventLogStore, error) {
	addr, err := r.fetchEventLogAddress(username)
	if err != nil {
		return nil, err
	}
	return r.orbitDb.Log(r.ctx, addr, nil)
}

// GetKeyValueStore creates or loads a key value database from given name
func (r *dbImpl) GetKeyValueStore(username string) (iface.KeyValueStore, error) {
	addr, err := r.fetchKeyValueAddress(username)
	if err != nil {
		return nil, err
	}
	return r.orbitDb.KeyValue(r.ctx, addr, nil)
}

//
// Helper functions
//

// fetchDocsAddress fetches the address of the document store for a given username
func (r *dbImpl) fetchDocsAddress(username string) (string, error) {
	addr, err := r.orbitDb.DetermineAddress(r.ctx, username, config.DB_DOCUMENT_STORE.String(), nil)
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}

// fetchEventLogAddress fetches the address of the event log for a given username
func (r *dbImpl) fetchEventLogAddress(username string) (string, error) {
	addr, err := r.orbitDb.DetermineAddress(r.ctx, username, config.DB_EVENT_LOG_STORE.String(), nil)
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}

// fetchKeyValueAddress fetches the address of the key value store for a given username
func (r *dbImpl) fetchKeyValueAddress(username string) (string, error) {
	addr, err := r.orbitDb.DetermineAddress(r.ctx, username, config.DB_KEY_VALUE_STORE.String(), nil)
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}
