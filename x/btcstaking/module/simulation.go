package btcstaking

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/AliErcanOzgokce/babylon-relayer/testutil/sample"
	btcstakingsimulation "github.com/AliErcanOzgokce/babylon-relayer/x/btcstaking/simulation"
	"github.com/AliErcanOzgokce/babylon-relayer/x/btcstaking/types"
)

// avoid unused import issue
var (
	_ = btcstakingsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgSendStakingData = "op_weight_msg_send_staking_data"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendStakingData int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	btcstakingGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&btcstakingGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSendStakingData int
	simState.AppParams.GetOrGenerate(opWeightMsgSendStakingData, &weightMsgSendStakingData, nil,
		func(_ *rand.Rand) {
			weightMsgSendStakingData = defaultWeightMsgSendStakingData
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSendStakingData,
		btcstakingsimulation.SimulateMsgSendStakingData(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSendStakingData,
			defaultWeightMsgSendStakingData,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				btcstakingsimulation.SimulateMsgSendStakingData(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
