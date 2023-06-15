package keeper

import (
	"context"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
)

type msgServer struct {
	Keeper
}

func (m msgServer) RemoveStandingMemberProposal(ctx context.Context, proposal *types.MsgRemoveStandingMemberProposal) (*types.MsgRemoveStandingMemberProposalResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) RegisterStandingMemberProposal(ctx context.Context, proposal *types.MsgRegisterStandingMemberProposal) (*types.MsgRegisterStandingMemberProposalResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
