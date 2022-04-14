package recovery

import (
	sdk "github.com/reapchain/cosmos-sdk/types"

	"github.com/reapchain/reapchain/x/recovery/keeper"
	"github.com/reapchain/reapchain/x/recovery/types"
)

// InitGenesis import module genesis
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	data types.GenesisState,
) {
	k.SetParams(ctx, data.Params)
}

// ExportGenesis export module status
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params: k.GetParams(ctx),
	}
}
