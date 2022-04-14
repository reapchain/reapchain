package recovery

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	capabilitytypes "github.com/reapchain/cosmos-sdk/x/capability/types"
	channeltypes "github.com/reapchain/ibc-go/modules/core/04-channel/types"
	porttypes "github.com/reapchain/ibc-go/modules/core/05-port/types"
	"github.com/reapchain/ibc-go/modules/core/exported"

	"github.com/reapchain/reapchain/ibc"
	"github.com/reapchain/reapchain/x/recovery/keeper"
)

var _ porttypes.Middleware = &IBCMiddleware{}

// IBCMiddleware implements the ICS26 callbacks for the transfer middleware given
// the recovery keeper and the underlying application.
type IBCMiddleware struct {
	*ibc.Module
	keeper keeper.Keeper
}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper and underlying application
func NewIBCMiddleware(k keeper.Keeper, app porttypes.IBCModule) IBCMiddleware {
	return IBCMiddleware{
		Module: ibc.NewModule(app),
		keeper: k,
	}
}

// OnRecvPacket implements the IBCModule interface.
// If the acknowledgement fails, this callback will default to the ibc-core
// packet callback.
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) exported.Acknowledgement {
	ack := im.Module.OnRecvPacket(ctx, packet, relayer)

	// return if the acknowledgement is an error ACK
	if !ack.Success() {
		return ack
	}

	return im.keeper.OnRecvPacket(ctx, packet, ack)
}

// SendPacket implements the ICS4 Wrapper interface
func (im IBCMiddleware) SendPacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet exported.PacketI,
) error {
	return im.keeper.SendPacket(ctx, chanCap, packet)
}

// WriteAcknowledgement implements the ICS4 Wrapper interface
func (im IBCMiddleware) WriteAcknowledgement(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet exported.PacketI,
	ack exported.Acknowledgement,
) error {
	return im.keeper.WriteAcknowledgement(ctx, chanCap, packet, ack)
}
