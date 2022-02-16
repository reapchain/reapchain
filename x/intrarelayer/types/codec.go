package types

import (
	"github.com/reapchain/cosmos-sdk/codec"
	codectypes "github.com/reapchain/cosmos-sdk/codec/types"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/types/msgservice"
	govtypes "github.com/reapchain/cosmos-sdk/x/gov/types"
)

// ModuleCdc references the global intrarelayer module codec. Note, the codec should
// ONLY be used in certain instances of tests and for JSON encoding.
//
// The actual codec used for serialization should be provided to modules/intrarelayer and
// defined at the application level.
var ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

// RegisterInterfaces register implementations
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgConvertCoin{},
		&MsgConvertERC20{},
	)
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&RegisterCoinProposal{},
		&RegisterERC20Proposal{},
		&ToggleTokenRelayProposal{},
		&UpdateTokenPairERC20Proposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
