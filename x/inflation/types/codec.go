package types

import (
	"github.com/reapchain/cosmos-sdk/codec"
	codectypes "github.com/reapchain/cosmos-sdk/codec/types"
)

// ModuleCdc references the global incentives module codec. Note, the codec
// should ONLY be used in certain instances of tests and for JSON encoding.
//
// The actual codec used for serialization should be provided to
// modules/incentives and defined at the application level.
var ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

// RegisterInterfaces register implementations
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {}
