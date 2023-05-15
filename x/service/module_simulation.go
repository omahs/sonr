package service

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sonrhq/core/testutil/sample"
	servicesimulation "github.com/sonrhq/core/x/service/simulation"
	"github.com/sonrhq/core/x/service/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = servicesimulation.FindAccount
	_ = sims.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateServiceRecord = "op_weight_msg_service_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateServiceRecord int = 100

	opWeightMsgUpdateServiceRecord = "op_weight_msg_service_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateServiceRecord int = 100

	opWeightMsgDeleteServiceRecord = "op_weight_msg_service_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteServiceRecord int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	serviceGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ServiceRecordList: []types.ServiceRecord{
			{
				Controller: sample.AccAddress(),
				Id:         "0",
			},
			{
				Controller: sample.AccAddress(),
				Id:         "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&serviceGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}


// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
