package witness

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v4/x/witness/keeper"
	"github.com/reapchain/reapchain/v4/x/witness/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the anchored
	for _, elem := range genState.AnchoredList {
		k.SetAnchored(ctx, elem)
	}

	// Set anchored count
	k.SetAnchoredCount(ctx, genState.AnchoredCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AnchoredList = k.GetAllAnchored(ctx)
	genesis.AnchoredCount = k.GetAnchoredCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
