package v2

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	paramtypes "github.com/reapchain/cosmos-sdk/x/params/types"

	"github.com/reapchain/reapchain/x/claims/types"
)

// MigrateStore adds the new parameters AuthorizedChannels and EVMChannels
// to the claims paramstore.
func MigrateStore(ctx sdk.Context, paramstore *paramtypes.Subspace) error {
	if !paramstore.HasKeyTable() {
		ps := paramstore.WithKeyTable(types.ParamKeyTable())
		paramstore = &ps
	}

	paramstore.Set(ctx, types.ParamStoreKeyAuthorizedChannels, types.DefaultAuthorizedChannels)
	paramstore.Set(ctx, types.ParamStoreKeyEVMChannels, types.DefaultEVMChannels)
	return nil
}
