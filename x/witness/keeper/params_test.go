package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "reapchain/testutil/keeper"
	"github.com/reapchain/reapchain/v4/x/witness/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.WitnessKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
