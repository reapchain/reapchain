package client

import (
	govclient "github.com/reapchain/cosmos-sdk/x/gov/client"

	"github.com/reapchain/reapchain/v8/x/escrow/client/cli"
	"github.com/reapchain/reapchain/v8/x/escrow/client/rest"
)

var (
	RegisterEscrowDenomProposalHandler    = govclient.NewProposalHandler(cli.NewRegisterEscrowDenomProposalCmd, rest.RegisterEscrowDenomProposalRESTHandler)
	ToggleEscrowConversionProposalHandler = govclient.NewProposalHandler(cli.NewToggleEscrowConversionProposalCmd, rest.ToggleEscrowConversionRESTHandler)
	AddEscrowSupplyProposalHandler        = govclient.NewProposalHandler(cli.NewAddEscrowSupplyProposalCmd, rest.AddEscrowSupplyRESTHandler)
)
