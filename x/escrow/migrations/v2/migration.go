package v2

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	paramtypes "github.com/reapchain/cosmos-sdk/x/params/types"

	"github.com/reapchain/reapchain/v8/x/escrow/types"
)

// UpdateParams updates the module parameters EnableERC20 and EnableEMVHook
// values to true.
func UpdateParams(ctx sdk.Context, paramstore *paramtypes.Subspace) error {
	if !paramstore.HasKeyTable() {
		ps := paramstore.WithKeyTable(types.ParamKeyTable())
		paramstore = &ps
	}

	paramstore.Set(ctx, types.ParamStoreKeyEnableEscrow, true)
	return nil
}
