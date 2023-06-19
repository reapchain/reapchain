package permissions_test

import (
	"testing"

	keepertest "github.com/reapchain/reapchain/v8/testutil/keeper"
	"github.com/reapchain/reapchain/v8/testutil/nullify"
	"github.com/reapchain/reapchain/v8/x/permissions"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keepertest.PermissionsKeeper(t)
	permissions.InitGenesis(ctx, *k, genesisState)
	got := permissions.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

}
