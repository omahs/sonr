package ipfs

import (
	"context"
	"fmt"
	"testing"

	"github.com/sonrhq/core/pkg/node/config"
	"github.com/stretchr/testify/assert"
)

func TestNewAddGet(t *testing.T) {
	// Call Run method and check for panic (if any)
	cnfg := config.DefaultConfig()
	node, err := Initialize(context.Background(), cnfg)
	if err != nil {
		t.Fatal(err)
	}

	// Add a file to the network
	cid, err := node.Add([]byte("Hello World!"))
	if err != nil {
		t.Fatal(err)
	}

	// Get the file from the network
	file, err := node.Get(cid)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("File: %s\n", file)
	fmt.Printf("CID: %s\n", cid)
	// Check if the file is the same as the one we added
	assert.Equal(t, []byte("Hello World!"), file)
}

func TestOrbitDB(t *testing.T) {
	cnfg := config.DefaultConfig()
	node, err := Initialize(context.Background(), cnfg)
	if err != nil {
		t.Fatal(err)
	}

	db, err := node.InitDB()
	if err != nil {
		t.Fatal(err)
	}

	docsStore, err := db.GetDocsStore("test")
	if err != nil {
		t.Fatal(err)
	}

	testData := map[string]interface{}{
		"_id":  "0",
		"test": "test",
	}
	op, err := docsStore.Put(context.Background(), testData)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("op: %v", op)
}
