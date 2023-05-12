package witness_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "reapchain/testutil/keeper"
	"reapchain/testutil/nullify"
	"github.com/reapchain/reapchain/v4/x/witness"
	"github.com/reapchain/reapchain/v4/x/witness/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AnchoredList: []types.Anchored{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AnchoredCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WitnessKeeper(t)
	witness.InitGenesis(ctx, *k, genesisState)
	got := witness.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AnchoredList, got.AnchoredList)
	require.Equal(t, genesisState.AnchoredCount, got.AnchoredCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
