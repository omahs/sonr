package domain_test

import (
	"testing"

	keepertest "github.com/sonrhq/core/testutil/keeper"
	"github.com/sonrhq/core/testutil/nullify"
	"github.com/sonrhq/core/x/domain"
	"github.com/sonrhq/core/x/domain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TLDRecordList: []types.TLDRecord{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		SLDRecordList: []types.SLDRecord{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DomainKeeper(t)
	domain.InitGenesis(ctx, *k, genesisState)
	got := domain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TLDRecordList, got.TLDRecordList)
	require.ElementsMatch(t, genesisState.SLDRecordList, got.SLDRecordList)
	// this line is used by starport scaffolding # genesis/test/assert
}
