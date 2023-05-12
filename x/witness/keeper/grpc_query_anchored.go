package keeper

import (
	"context"

	"github.com/reapchain/cosmos-sdk/store/prefix"
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	"github.com/reapchain/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/reapchain/reapchain/v4/x/witness/types"
)

func (k Keeper) AnchoredAll(c context.Context, req *types.QueryAllAnchoredRequest) (*types.QueryAllAnchoredResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var anchoreds []types.Anchored
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	anchoredStore := prefix.NewStore(store, types.KeyPrefix(types.AnchoredKey))

	pageRes, err := query.Paginate(anchoredStore, req.Pagination, func(key []byte, value []byte) error {
		var anchored types.Anchored
		if err := k.cdc.Unmarshal(value, &anchored); err != nil {
			return err
		}

		anchoreds = append(anchoreds, anchored)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAnchoredResponse{Anchored: anchoreds, Pagination: pageRes}, nil
}

func (k Keeper) Anchored(c context.Context, req *types.QueryGetAnchoredRequest) (*types.QueryGetAnchoredResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	anchored, found := k.GetAnchored(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAnchoredResponse{Anchored: anchored}, nil
}
