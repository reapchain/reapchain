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
	ProposalTypeAddToEscrowPool        string = "AddToEscrowPool"
)

// Implements Proposal Interface
var (
	_ govtypes.Content = &RegisterEscrowDenomProposal{}
	_ govtypes.Content = &ToggleEscrowConversionProposal{}
	_ govtypes.Content = &AddToEscrowPoolProposal{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeRegisterEscrowDenom)
	govtypes.RegisterProposalType(ProposalTypeToggleEscrowConversion)
	govtypes.RegisterProposalType(ProposalTypeAddToEscrowPool)

	govtypes.RegisterProposalTypeCodec(&RegisterEscrowDenomProposal{}, "escrow/RegisterEscrowDenomProposal")
	govtypes.RegisterProposalTypeCodec(&ToggleEscrowConversionProposal{}, "escrow/ToggleEscrowConversionProposal")
	govtypes.RegisterProposalTypeCodec(&AddToEscrowPoolProposal{}, "escrow/AddToEscrowPoolProposal")

}

// CreateDenomDescription generates a string with the coin description
func CreateDenomDescription(address string) string {
	return fmt.Sprintf("Cosmos coin token representation of %s", address)
}

func NewRegisterEscrowDenomProposal(title, description, denom string, initialPoolBalance sdk.Int) govtypes.Content {
	return &RegisterEscrowDenomProposal{
		Title:              title,
		Description:        description,
		Denom:              denom,
		InitialPoolBalance: initialPoolBalance,
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

func NewRegisterEscrowDenomAndConvertProposal(title, description, denom string, initialPoolBalance sdk.Int, proposer string, receiver string) govtypes.Content {
	return &RegisterEscrowDenomAndConvertProposal{
		Title:              title,
		Description:        description,
		Denom:              denom,
		InitialPoolBalance: initialPoolBalance,
		Proposer:           proposer,
		Receiver:           receiver,
	}
}

// ProposalRoute returns router key for this proposal
func (*RegisterEscrowDenomAndConvertProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (*RegisterEscrowDenomAndConvertProposal) ProposalType() string {
	return ProposalTypeRegisterEscrowDenom
}

// ValidateBasic performs a stateless check of the proposal fields
func (rtbp *RegisterEscrowDenomAndConvertProposal) ValidateBasic() error {
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

func NewAddToEscrowPoolProposal(title, description, denom string, amount sdk.Int) govtypes.Content {
	return &AddToEscrowPoolProposal{
		Title:       title,
		Description: description,
		Denom:       denom,
		Amount:      amount,
	}
}

// ProposalRoute returns router key for this proposal
func (*AddToEscrowPoolProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (*AddToEscrowPoolProposal) ProposalType() string {
	return ProposalTypeRegisterEscrowDenom
}

// ValidateBasic performs a stateless check of the proposal fields
func (aesp *AddToEscrowPoolProposal) ValidateBasic() error {

	EVMAddress := removePrefix(aesp.Denom)
	if err := ethermint.ValidateAddress(EVMAddress); err != nil {
		return sdkerrors.Wrap(err, "Invalid EVM address")
	}
	return govtypes.ValidateAbstract(aesp)
}

func NewAddToEscrowPoolAndConvertProposal(title, description, denom string, amount sdk.Int, proposer string, receiver string) govtypes.Content {
	return &AddToEscrowPoolAndConvertProposal{
		Title:       title,
		Description: description,
		Denom:       denom,
		Amount:      amount,
		Proposer:    proposer,
		Receiver:    receiver,
	}
}

// ProposalRoute returns router key for this proposal
func (*AddToEscrowPoolAndConvertProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (*AddToEscrowPoolAndConvertProposal) ProposalType() string {
	return ProposalTypeRegisterEscrowDenom
}

// ValidateBasic performs a stateless check of the proposal fields
func (aesp *AddToEscrowPoolAndConvertProposal) ValidateBasic() error {

	EVMAddress := removePrefix(aesp.Denom)
	if err := ethermint.ValidateAddress(EVMAddress); err != nil {
		return sdkerrors.Wrap(err, "Invalid EVM address")
	}
	return govtypes.ValidateAbstract(aesp)
}
