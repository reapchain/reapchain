package bech32ibc_test

import (
	"testing"
	"time"

	simapp "github.com/osmosis-labs/bech32-ibc/app"
	tmproto "github.com/reapchain/reapchain-core/proto/tendermint/types"
	"github.com/reapchain/reapchain/v8/x/bech32ibc/types"
	"github.com/stretchr/testify/require"
)

func TestBech32IBCProposalHandler(t *testing.T) {
	now := time.Now()
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	ctx = ctx.WithBlockTime(now.Add(time.Second))

	proposal := &types.UpdateHrpIbcChannelProposal{
		Title:         "update hrp ibc channel",
		Description:   "update hrp ibc channel",
		Hrp:           "akash",
		SourceChannel: "1",
	}
	handler := app.GovKeeper.Router().GetRoute(proposal.ProposalRoute())
	require.NotPanics(t, func() {
		err := handler(ctx, proposal)
		require.Error(t, err) // error for not existing source channel
	})
}
