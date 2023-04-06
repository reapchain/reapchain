package client

import (
	govclient "github.com/reapchain/cosmos-sdk/x/gov/client"

	"github.com/reapchain/reapchain/v8/x/incentives/client/cli"
	"github.com/reapchain/reapchain/v8/x/incentives/client/rest"
)

var (
	RegisterIncentiveProposalHandler = govclient.NewProposalHandler(cli.NewRegisterIncentiveProposalCmd, rest.RegisterIncentiveProposalRESTHandler)
	CancelIncentiveProposalHandler   = govclient.NewProposalHandler(cli.NewCancelIncentiveProposalCmd, rest.CancelIncentiveProposalRequestRESTHandler)
)
