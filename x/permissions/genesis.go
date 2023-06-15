package permissions

import (
	"fmt"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v8/x/permissions/keeper"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
	"time"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, sk types.StakingKeeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	validators := sk.GetAllValidators(ctx)
	for _, validator := range validators {

		if validator.Type == "standing" {
			fmt.Println("\n\n---------------------------------------------------------------------")
			fmt.Println("PERMISSIONS MODULE - InitGenesis", time.Now().Format(time.RFC822))
			fmt.Println("ADDRESS: ", validator.OperatorAddress)
			fmt.Println("VALIDATOR TYPE: ", validator.Type)

			validatorAddress, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
			accountAddress := sdk.AccAddress(validatorAddress).String()

			var whiteListedValidator = types.WhitelistedValidator{
				ValidatorAddress: validator.OperatorAddress,
				Moniker:          validator.GetMoniker(),
				AccountAddress:   accountAddress,
			}
			_ = k.AppendWhitelistedValidator(
				ctx,
				whiteListedValidator,
			)
		}
	}

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
