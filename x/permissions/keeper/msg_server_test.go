package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/reapchain/cosmos-sdk/types"
	keepertest "github.com/reapchain/reapchain/v8/testutil/keeper"
	"github.com/reapchain/reapchain/v8/x/permissions/keeper"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.PermissionsKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
