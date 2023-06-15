package permissions

import (
	"github.com/reapchain/cosmos-sdk/crypto/keys/ed25519"
	"math/rand"

	"github.com/reapchain/cosmos-sdk/baseapp"
	simappparams "github.com/reapchain/cosmos-sdk/simapp/params"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/types/module"
	simtypes "github.com/reapchain/cosmos-sdk/types/simulation"
	"github.com/reapchain/cosmos-sdk/x/simulation"
	permissionssimulation "github.com/reapchain/reapchain/v8/x/permissions/simulation"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
)

// avoid unused import issue
var (
	_ = AccAddress
	_ = permissionssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
// this line is used by starport scaffolding # simapp/module/const
)

func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	permissionsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&permissionsGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	permissionsParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyForcedUnbondingTime), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(permissionsParams.ForcedUnbondingTime))
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyGovernanceMinimumInitialDepositEnabled), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(permissionsParams.GovernanceMinimumInitialDepositEnabled))
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyGovernanceMinimumInitialDepositPercentage), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(permissionsParams.GovernanceMinimumInitialDepositPercentage))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
