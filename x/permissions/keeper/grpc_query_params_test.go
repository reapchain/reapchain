package keeper_test

import (
	"testing"

	sdk "github.com/reapchain/cosmos-sdk/types"
	testkeeper "github.com/reapchain/reapchain/v8/testutil/keeper"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.PermissionsKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
