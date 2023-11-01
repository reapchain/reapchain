package escrow

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	authkeeper "github.com/reapchain/cosmos-sdk/x/auth/keeper"

	"github.com/reapchain/reapchain/v8/x/escrow/keeper"
	"github.com/reapchain/reapchain/v8/x/escrow/types"
)

// InitGenesis import module genesis
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
	data types.GenesisState,
) {
	k.SetParams(ctx, data.Params)

	if acc := accountKeeper.GetModuleAccount(ctx, types.ModuleName); acc == nil {
		// NOTE: shouldn't occur
		panic("the escrow module account has not been set")
	}

	for _, denom := range data.RegisteredDenoms {
		k.RegisterDenom(ctx, denom)
	}
	for _, escrowPool := range data.GetEscrowPools() {
		k.SetEscrowPool(ctx, escrowPool)
	}
}

// ExportGenesis export module status
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:           k.GetParams(ctx),
		RegisteredDenoms: k.GetRegisteredDenoms(ctx),
		EscrowPools:      k.GetAllEscrowPools(ctx),
	}
}
