package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/reapchain/reapchain-core/crypto/tmhash"
	tmproto "github.com/reapchain/reapchain-core/proto/podc/types"
	tmversion "github.com/reapchain/reapchain-core/proto/podc/version"
	"github.com/reapchain/reapchain-core/version"

	"github.com/reapchain/ethermint/tests"
	feemarkettypes "github.com/reapchain/ethermint/x/feemarket/types"

	"github.com/reapchain/cosmos-sdk/baseapp"
	sdk "github.com/reapchain/cosmos-sdk/types"

	"github.com/reapchain/reapchain/v8/app"
	claimstypes "github.com/reapchain/reapchain/v8/x/claims/types"
	"github.com/reapchain/reapchain/v8/x/recovery/types"
)

var (
	ibcAtomDenom = "ibc/A4DB47A9D3CF9A068D454513891B526702455D3EF08FB9EB558C561F9DC2B701"
	ibcOsmoDenom = "ibc/ED07A3391A112B175915CD8FAF43A2DA8E4790EDE12566649D0C2F97716B8518"
	erc20Denom   = "erc20/0xdac17f958d2ee523a2206206994597c13d831ec7"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx sdk.Context

	app         *app.Reapchain
	queryClient types.QueryClient
}

func (suite *KeeperTestSuite) SetupTest() {
	// consensus key
	consAddress := sdk.ConsAddress(tests.GenerateAddress().Bytes())

	suite.app = app.Setup(false, feemarkettypes.DefaultGenesisState())
	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{
		Height:          1,
		ChainID:         "reapchain_221230-1",
		Time:            time.Now().UTC(),
		ProposerAddress: consAddress.Bytes(),

		Version: tmversion.Consensus{
			Block: version.BlockProtocol,
		},
		LastBlockId: tmproto.BlockID{
			Hash: tmhash.Sum([]byte("block_id")),
			PartSetHeader: tmproto.PartSetHeader{
				Total: 11,
				Hash:  tmhash.Sum([]byte("partset_header")),
			},
		},
		AppHash:            tmhash.Sum([]byte("app")),
		DataHash:           tmhash.Sum([]byte("data")),
		EvidenceHash:       tmhash.Sum([]byte("evidence")),
		ValidatorsHash:     tmhash.Sum([]byte("validators")),
		NextValidatorsHash: tmhash.Sum([]byte("next_validators")),
		ConsensusHash:      tmhash.Sum([]byte("consensus")),
		LastResultsHash:    tmhash.Sum([]byte("last_result")),
	})

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, suite.app.RecoveryKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)

	claimsParams := claimstypes.DefaultParams()
	claimsParams.AirdropStartTime = suite.ctx.BlockTime()
	suite.app.ClaimsKeeper.SetParams(suite.ctx, claimsParams)

	stakingParams := suite.app.StakingKeeper.GetParams(suite.ctx)
	stakingParams.BondDenom = claimsParams.GetClaimsDenom()
	suite.app.StakingKeeper.SetParams(suite.ctx, stakingParams)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
