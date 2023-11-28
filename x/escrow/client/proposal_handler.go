package client

import (
	govclient "github.com/reapchain/cosmos-sdk/x/gov/client"

	"github.com/reapchain/reapchain/v8/x/escrow/client/cli"
	"github.com/reapchain/reapchain/v8/x/escrow/client/rest"
)

var (
	RegisterEscrowDenomProposalHandler           = govclient.NewProposalHandler(cli.NewRegisterEscrowDenomProposalCmd, rest.RegisterEscrowDenomProposalRESTHandler)
	RegisterEscrowDenomAndConvertProposalHandler = govclient.NewProposalHandler(cli.NewRegisterEscrowDenomAndConvertProposalCmd, rest.RegisterEscrowDenomAndConvertProposalRESTHandler)

	ToggleEscrowConversionProposalHandler    = govclient.NewProposalHandler(cli.NewToggleEscrowConversionProposalCmd, rest.ToggleEscrowConversionRESTHandler)
	AddToEscrowPoolProposalHandler           = govclient.NewProposalHandler(cli.NewAddToEscrowPoolProposalCmd, rest.AddToEscrowPoolRESTHandler)
	AddToEscrowPoolAndConvertProposalHandler = govclient.NewProposalHandler(cli.NewAddToEscrowPoolAndConvertProposalCmd, rest.AddToEscrowPoolAndConvertRESTHandler)
)
