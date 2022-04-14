package types

import (
	"github.com/reapchain/cosmos-sdk/codec"
	codectypes "github.com/reapchain/cosmos-sdk/codec/types"
)

var ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
