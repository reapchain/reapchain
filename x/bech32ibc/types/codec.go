package types

import (
	"github.com/reapchain/cosmos-sdk/codec"
	cdctypes "github.com/reapchain/cosmos-sdk/codec/types"
	govtypes "github.com/reapchain/cosmos-sdk/x/gov/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&UpdateHrpIbcChannelProposal{}, "osmosis/UpdateHrpIbcChannelProposal", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&UpdateHrpIbcChannelProposal{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
