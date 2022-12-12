package keeper_test

import (
	sdk "github.com/reapchain/cosmos-sdk/types"

	"github.com/reapchain/reapchain/v4/x/claims/types"
)

func (suite *KeeperTestSuite) TestsClaimsRecords() {
	addr1, err := sdk.AccAddressFromBech32("reap10hz0qunf49j9rvn3twrtdu0cm8hnav47vawaxm")
	suite.Require().NoError(err)
	addr2, err := sdk.AccAddressFromBech32("reap1ktkajgqv25830rgve5qwhwfs6m4hjruaml33ff")
	suite.Require().NoError(err)

	cr1 := types.NewClaimsRecord(sdk.NewInt(1000))
	cr2 := types.NewClaimsRecord(sdk.NewInt(200))
	cr2.MarkClaimed(types.ActionDelegate)

	expRecords := []types.ClaimsRecordAddress{
		{
			Address:                addr2.String(),
			InitialClaimableAmount: cr2.InitialClaimableAmount,
			ActionsCompleted:       cr2.ActionsCompleted,
		},
		{
			Address:                addr1.String(),
			InitialClaimableAmount: cr1.InitialClaimableAmount,
			ActionsCompleted:       cr1.ActionsCompleted,
		},
	}

	suite.app.ClaimsKeeper.SetClaimsRecord(suite.ctx, addr1, cr1)
	suite.app.ClaimsKeeper.SetClaimsRecord(suite.ctx, addr2, cr2)

	records := suite.app.ClaimsKeeper.GetClaimsRecords(suite.ctx)
	suite.Require().Equal(expRecords, records)
}
