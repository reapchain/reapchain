package client

import (
	govclient "github.com/reapchain/cosmos-sdk/x/gov/client"
	"github.com/reapchain/mercury/x/intrarelayer/client/cli"
	"github.com/reapchain/mercury/x/intrarelayer/client/rest"
)

var (
	RegisterCoinProposalHandler         = govclient.NewProposalHandler(cli.NewRegisterCoinProposalCmd, rest.RegisterCoinProposalRESTHandler)
	RegisterERC20ProposalHandler        = govclient.NewProposalHandler(cli.NewRegisterERC20ProposalCmd, rest.RegisterERC20ProposalRESTHandler)
	ToggleTokenRelayProposalHandler     = govclient.NewProposalHandler(cli.NewToggleTokenRelayProposalCmd, rest.ToggleTokenRelayRESTHandler)
	UpdateTokenPairERC20ProposalHandler = govclient.NewProposalHandler(cli.NewUpdateTokenPairERC20ProposalCmd, rest.UpdateTokenPairERC20ProposalRESTHandler)
)
