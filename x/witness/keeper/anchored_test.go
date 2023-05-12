package keeper_test

import (
	"testing"

	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "reapchain/testutil/keeper"
	"reapchain/testutil/nullify"
	"github.com/reapchain/reapchain/v4/x/witness/keeper"
	"github.com/reapchain/reapchain/v4/x/witness/types"
)

func createNAnchored(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Anchored {
	items := make([]types.Anchored, n)
	for i := range items {
		items[i].Id = keeper.AppendAnchored(ctx, items[i])
	}
	return items
}

func TestAnchoredGet(t *testing.T) {
	keeper, ctx := keepertest.WitnessKeeper(t)
	items := createNAnchored(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAnchored(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAnchoredRemove(t *testing.T) {
	keeper, ctx := keepertest.WitnessKeeper(t)
	items := createNAnchored(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAnchored(ctx, item.Id)
		_, found := keeper.GetAnchored(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAnchoredGetAll(t *testing.T) {
	keeper, ctx := keepertest.WitnessKeeper(t)
	items := createNAnchored(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAnchored(ctx)),
	)
}

func TestAnchoredCount(t *testing.T) {
	keeper, ctx := keepertest.WitnessKeeper(t)
	items := createNAnchored(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAnchoredCount(ctx))
}
