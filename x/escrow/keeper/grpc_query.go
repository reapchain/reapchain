package keeper

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/reapchain/cosmos-sdk/store/prefix"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/types/query"
	"github.com/reapchain/reapchain/v8/x/escrow/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) RegisteredDenoms(c context.Context, req *types.QueryRegisteredDenomsRequest) (*types.QueryRegisteredDenomsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	var registeredDenoms []types.RegisteredDenom
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixDenom)

	pageRes, err := query.Paginate(store, req.Pagination, func(_, value []byte) error {
		var registeredDenom types.RegisteredDenom
		if err := k.cdc.Unmarshal(value, &registeredDenom); err != nil {
			return err
		}
		registeredDenoms = append(registeredDenoms, registeredDenom)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryRegisteredDenomsResponse{
		RegisteredDenoms: registeredDenoms,
		Pagination:       pageRes,
	}, nil
}

func (k Keeper) EscrowPoolBalance(c context.Context, req *types.QueryEscrowPoolBalanceRequest) (*types.QueryEscrowPoolBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	escrowPool, _ := k.GetEscrowPoolByDenom(ctx, req.Denom)

	return &types.QueryEscrowPoolBalanceResponse{
		EscrowPoolBalance: escrowPool.Coins,
	}, nil
}

// Params returns the params of the erc20 module
func (k Keeper) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)
	return &types.QueryParamsResponse{Params: params}, nil
}
