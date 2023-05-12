package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/reapchain/cosmos-sdk/types"
	keepertest "reapchain/testutil/keeper"
	"github.com/reapchain/reapchain/v4/x/witness/keeper"
	"github.com/reapchain/reapchain/v4/x/witness/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.WitnessKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
