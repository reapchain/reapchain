package reap_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/tharsis/evmos/testutil/keeper"
	"github.com/tharsis/evmos/x/reap"
	"github.com/tharsis/evmos/x/reap/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ReapKeeper(t)
	reap.InitGenesis(ctx, *k, genesisState)
	got := reap.ExportGenesis(ctx, *k)
	require.NotNil(t, got)


	// this line is used by starport scaffolding # genesis/test/assert
}
