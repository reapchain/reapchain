package v2_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/reapchain/cosmos-sdk/testutil"
	sdk "github.com/reapchain/cosmos-sdk/types"
	paramtypes "github.com/reapchain/cosmos-sdk/x/params/types"

	"github.com/reapchain/ethermint/encoding"

	"github.com/reapchain/reapchain/v8/app"
	v2 "github.com/reapchain/reapchain/v8/x/escrow/migrations/v2"
	escrowtypes "github.com/reapchain/reapchain/v8/x/escrow/types"
)

func TestUpdateParams(t *testing.T) {
	encCfg := encoding.MakeConfig(app.ModuleBasics)
	erc20Key := sdk.NewKVStoreKey(escrowtypes.StoreKey)
	tErc20Key := sdk.NewTransientStoreKey(fmt.Sprintf("%s_test", escrowtypes.StoreKey))
	ctx := testutil.DefaultContext(erc20Key, tErc20Key)
	paramstore := paramtypes.NewSubspace(
		encCfg.Marshaler, encCfg.Amino, erc20Key, tErc20Key, "erc20",
	)
	paramstore = paramstore.WithKeyTable(escrowtypes.ParamKeyTable())
	require.True(t, paramstore.HasKeyTable())

	// check no params
	require.False(t, paramstore.Has(ctx, escrowtypes.ParamStoreKeyEnableEscrow))

	// Run migrations
	err := v2.UpdateParams(ctx, &paramstore)
	require.NoError(t, err)

	// Make sure the params are set
	require.True(t, paramstore.Has(ctx, escrowtypes.ParamStoreKeyEnableEscrow))

	var enableERC20, enableEVMHook bool

	// Make sure the new params are set
	require.NotPanics(t, func() {
		paramstore.Get(ctx, escrowtypes.ParamStoreKeyEnableEscrow, &enableERC20)
	})

	// check the params are updated
	require.True(t, enableERC20)
	require.True(t, enableEVMHook)
}
