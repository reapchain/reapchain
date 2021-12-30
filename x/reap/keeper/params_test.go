package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/tharsis/evmos/testutil/keeper"
	"github.com/tharsis/evmos/x/reap/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ReapKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
