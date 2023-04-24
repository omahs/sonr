package identity_test

import (
	"fmt"
	"testing"
	"time"

	keepertest "github.com/sonrhq/core/testutil/keeper"
	"github.com/sonrhq/core/testutil/nullify"
	"github.com/sonrhq/core/x/identity"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PrimaryIdentities: []types.DidDocument{
			{
				Id: "0",
			},
			{
				Id: "1",
			},
		},
		ClaimableWalletList: []types.ClaimableWallet{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ClaimableWalletCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IdentityKeeper(t)
	identity.InitGenesis(ctx, *k, genesisState)
	got := identity.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PrimaryIdentities, got.PrimaryIdentities)
	require.ElementsMatch(t, genesisState.ClaimableWalletList, got.ClaimableWalletList)
	require.Equal(t, genesisState.ClaimableWalletCount, got.ClaimableWalletCount)
	// this line is used by starport scaffolding # genesis/test/assert
}

func TestSequencer(t *testing.T) {
	sq := identity.NewSequencer()
	for i := 0; i < 20; i++ {
		w := sq.Next()
		if w == nil {
			t.Log("No wallet available")
			time.Sleep(time.Second)
			continue
		}
		t.Logf("Wallet %d: %s", w.Id, w.PublicKey)
		err := sq.Add()
		if err != nil {
			t.Error(err)
		}
	}
}

func TestQueue(t *testing.T) {
	var products = []string{
		"books",
		"computers",
	}
	// New products, which we need to add to our products storage.
	newProducts := []string{
		"apples",
		"oranges",
		"wine",
		"bread",
		"orange juice",
	}
	// New queue initialization.
	productsQueue := identity.NewQueue("NewProducts")
	var jobs []identity.Job

	// Range over new products.
	for _, newProduct := range newProducts {
		// We need to do this, because variables declared in for loops are passed by reference.
		// Otherwise, our closure will always receive the last item from the newProducts.
		product := newProduct
		// Defining of the closure, where we add a new product to our simple storage (products slice)
		action := func() error {
			products = append(products, product)
			return nil
		}
		// Append job to jobs slice.
		jobs = append(jobs, identity.Job{
			Name:   fmt.Sprintf("Importing new product: %s", newProduct),
			Action: action,
		})
	}

	// Adds jobs to the queue.
	productsQueue.AddJobs(jobs)

	// Defines a queue worker, which will execute our queue.
	worker := identity.NewWorker(productsQueue)
	// Execute jobs in queue.
	worker.DoWork()

	// Prints products storage after queue execution.
	defer fmt.Print(products)
}
