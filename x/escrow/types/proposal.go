package types

import (
	"fmt"
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	govtypes "github.com/reapchain/cosmos-sdk/x/gov/types"
	ethermint "github.com/reapchain/ethermint/types"
)

// constants
const (
	ProposalTypeRegisterEscrowDenom    string = "RegisterEscrowDenom"
	ProposalTypeToggleEscrowConversion string = "ToggleEscrowConversion"
	ProposalTypeAddEscrowSupply        string = "AddEscrowSupply"
)

// Implements Proposal Interface
var (
	_ govtypes.Content = &RegisterEscrowDenomProposal{}
	_ govtypes.Content = &ToggleEscrowConversionProposal{}
	_ govtypes.Content = &AddEscrowSupplyProposal{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeRegisterEscrowDenom)
	govtypes.RegisterProposalType(ProposalTypeToggleEscrowConversion)
	govtypes.RegisterProposalType(ProposalTypeAddEscrowSupply)

	govtypes.RegisterProposalTypeCodec(&RegisterEscrowDenomProposal{}, "escrow/RegisterEscrowDenomProposal")
	govtypes.RegisterProposalTypeCodec(&ToggleEscrowConversionProposal{}, "escrow/ToggleEscrowConversionProposal")
	govtypes.RegisterProposalTypeCodec(&AddEscrowSupplyProposal{}, "escrow/AddEscrowSupplyProposal")

}

// CreateDenomDescription generates a string with the coin description
func CreateDenomDescription(address string) string {
	return fmt.Sprintf("Cosmos coin token representation of %s", address)
}

func NewRegisterEscrowDenomProposal(title, description, denom string, initialSupply sdk.Int) govtypes.Content {
	return &RegisterEscrowDenomProposal{
		Title:         title,
		Description:   description,
		Denom:         denom,
		InitialSupply: initialSupply,
	}
}

// ProposalRoute returns router key for this proposal
func (*RegisterEscrowDenomProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (*RegisterEscrowDenomProposal) ProposalType() string {
	return ProposalTypeRegisterEscrowDenom
}

// ValidateBasic performs a stateless check of the proposal fields
func (rtbp *RegisterEscrowDenomProposal) ValidateBasic() error {
	EVMAddress := removePrefix(rtbp.Denom)
	if err := ethermint.ValidateAddress(EVMAddress); err != nil {
		return sdkerrors.Wrap(err, "Invalid EVM address")
	}
	return govtypes.ValidateAbstract(rtbp)
}

func NewToggleEscrowConversionProposal(title, description string, denom string) govtypes.Content {
	return &ToggleEscrowConversionProposal{
		Title:       title,
		Description: description,
		Denom:       denom,
	}
}

// ProposalRoute returns router key for this proposal
func (*ToggleEscrowConversionProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (*ToggleEscrowConversionProposal) ProposalType() string {
	return ProposalTypeToggleEscrowConversion
}

func (ttcp *ToggleEscrowConversionProposal) ValidateBasic() error {
	EVMAddress := removePrefix(ttcp.Denom)
	if err := ethermint.ValidateAddress(EVMAddress); err != nil {
		return sdkerrors.Wrap(err, "Invalid EVM address")
	}

	return govtypes.ValidateAbstract(ttcp)
}

func NewAddEscrowSupplyProposal(title, description, denom string) govtypes.Content {
	return &AddEscrowSupplyProposal{
		Title:       title,
		Description: description,
		Denom:       denom,
	}
}

// ProposalRoute returns router key for this proposal
func (*AddEscrowSupplyProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (*AddEscrowSupplyProposal) ProposalType() string {
	return ProposalTypeRegisterEscrowDenom
}

// ValidateBasic performs a stateless check of the proposal fields
func (aesp *AddEscrowSupplyProposal) ValidateBasic() error {

	EVMAddress := removePrefix(aesp.Denom)
	if err := ethermint.ValidateAddress(EVMAddress); err != nil {
		return sdkerrors.Wrap(err, "Invalid EVM address")
	}
	return govtypes.ValidateAbstract(aesp)
}
