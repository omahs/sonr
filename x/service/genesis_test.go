package service_test

import (
	"testing"

	keepertest "github.com/sonrhq/core/testutil/keeper"
	"github.com/sonrhq/core/testutil/nullify"
	"github.com/sonrhq/core/x/service"
	"github.com/sonrhq/core/x/service/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ServiceRecordList: []types.ServiceRecord{
			{
				Id: "0",
			},
			{
				Id: "1",
			},
		},
		ServiceRelationshipsList: []types.ServiceRelationship{
			{
				Did: "0",
			},
			{
				Did: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ServiceKeeper(t)
	service.InitGenesis(ctx, *k, genesisState)
	got := service.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ServiceRecordList, got.ServiceRecordList)
	require.ElementsMatch(t, genesisState.ServiceRelationshipsList, got.ServiceRelationshipsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
