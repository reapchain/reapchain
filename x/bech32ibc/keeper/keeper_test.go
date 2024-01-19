package keeper_test

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	tmproto "github.com/reapchain/reapchain-core/proto/podc/types"
	"testing"

	"github.com/reapchain/cosmos-sdk/baseapp"
	"github.com/reapchain/reapchain/v8/app"
	"github.com/reapchain/reapchain/v8/x/bech32ibc/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	app         *app.Reapchain
	ctx         sdk.Context
	queryClient types.QueryClient
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app = app.Setup(false, nil)
	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{})

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, suite.app.Bech32IBCKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
