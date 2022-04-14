package keeper

import (
	"context"

	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/x/inflation/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

// Period returns the current period of the inflation module.
func (k Keeper) Period(
	c context.Context,
	_ *types.QueryPeriodRequest,
) (*types.QueryPeriodResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	period := k.GetPeriod(ctx)
	return &types.QueryPeriodResponse{Period: period}, nil
}

// EpochMintProvision returns the EpochMintProvision of the inflation module.
func (k Keeper) EpochMintProvision(
	c context.Context,
	_ *types.QueryEpochMintProvisionRequest,
) (*types.QueryEpochMintProvisionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	epochMintProvision, found := k.GetEpochMintProvision(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "epoch mint provision not found")
	}

	mintDenom := k.GetParams(ctx).MintDenom
	coin := sdk.NewDecCoinFromDec(mintDenom, epochMintProvision)

	return &types.QueryEpochMintProvisionResponse{EpochMintProvision: coin}, nil
}

// SkippedEpochs returns the number of skipped Epochs of the inflation module.
func (k Keeper) SkippedEpochs(
	c context.Context,
	_ *types.QuerySkippedEpochsRequest,
) (*types.QuerySkippedEpochsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	skippedEpochs := k.GetSkippedEpochs(ctx)
	return &types.QuerySkippedEpochsResponse{SkippedEpochs: skippedEpochs}, nil
}

// InflationRate returns the number of skipped Epochs of the inflation module.
func (k Keeper) InflationRate(
	c context.Context,
	_ *types.QueryInflationRateRequest,
) (*types.QueryInflationRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	inflationRate := k.GetInflationRate(ctx)

	return &types.QueryInflationRateResponse{InflationRate: inflationRate}, nil
}

// TotalSupply returns the number of skipped Epochs of the inflation module.
func (k Keeper) TotalSupply(
	c context.Context,
	_ *types.QueryTotalSupplyRequest,
) (*types.QueryTotalSupplyResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	totalSupply := k.GetTotalSupply(ctx)

	mintDenom := k.GetParams(ctx).MintDenom
	coin := sdk.NewDecCoinFromDec(mintDenom, totalSupply)

	return &types.QueryTotalSupplyResponse{TotalSupply: coin}, nil
}

// Params returns params of the mint module.
func (k Keeper) Params(
	c context.Context,
	_ *types.QueryParamsRequest,
) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)
	return &types.QueryParamsResponse{Params: params}, nil
}
