package service

import (
	"math/rand"

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

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateServiceRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateServiceRecord, &weightMsgCreateServiceRecord, nil,
		func(_ *rand.Rand) {
			weightMsgCreateServiceRecord = defaultWeightMsgCreateServiceRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateServiceRecord,
		servicesimulation.SimulateMsgCreateServiceRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateServiceRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateServiceRecord, &weightMsgUpdateServiceRecord, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateServiceRecord = defaultWeightMsgUpdateServiceRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateServiceRecord,
		servicesimulation.SimulateMsgUpdateServiceRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteServiceRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteServiceRecord, &weightMsgDeleteServiceRecord, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteServiceRecord = defaultWeightMsgDeleteServiceRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteServiceRecord,
		servicesimulation.SimulateMsgDeleteServiceRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
