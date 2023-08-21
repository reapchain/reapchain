package keeper_test

import (
	"testing"

	"github.com/osmosis-labs/bech32-ibc/app"
	"github.com/reapchain/cosmos-sdk/baseapp"
	sdk "github.com/reapchain/cosmos-sdk/types"
	tmproto "github.com/reapchain/reapchain-core/proto/tendermint/types"
	"github.com/reapchain/reapchain/v8/x/bech32ibc/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	app         *app.App
	ctx         sdk.Context
	queryClient types.QueryClient
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app = app.Setup(false)
	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{})

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, suite.app.Bech32IBCKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
