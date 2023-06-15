package client

import (
	govclient "github.com/reapchain/cosmos-sdk/x/gov/client"

	"github.com/reapchain/reapchain/v8/x/permissions/client/cli"
	"github.com/reapchain/reapchain/v8/x/permissions/client/rest"
)

var (
	RegisterStandingMemberProposal = govclient.NewProposalHandler(cli.CmdRegisterStandingMemberProposal, rest.RegisterStandingMemberProposalRESTHandler)
	RemoveStandingMemberProposal   = govclient.NewProposalHandler(cli.CmdRemoveStandingMemberProposal, rest.RemoveStandingMemberProposalRESTHandler)
	ReplaceStandingMemberProposal  = govclient.NewProposalHandler(cli.CmdReplaceStandingMemberProposal, rest.ReplaceStandingMemberProposalRESTHandler)
)
