package keeper_test

import (
	"testing"

	testkeeper "github.com/reapchain/reapchain/v8/testutil/keeper"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PermissionsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.PermissionsUnbondingTime, k.PermissionsUnbondingTime(ctx))
	require.EqualValues(t, params.PermissionsMinimumInitialDepositEnabled, k.PermissionsMinimumInitialDepositEnabled(ctx))
	require.EqualValues(t, params.PermissionsMinimumInitialDepositPercentage, k.PermissionsMinimumInitialDepositPercentage(ctx))
}
