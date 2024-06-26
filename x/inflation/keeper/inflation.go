package keeper

import (
	sdk "github.com/reapchain/cosmos-sdk/types"

	incentivestypes "github.com/reapchain/reapchain/v8/x/incentives/types"
	"github.com/reapchain/reapchain/v8/x/inflation/types"
)

// MintAndAllocateInflation performs inflation minting and allocation
func (k Keeper) MintAndAllocateInflation(
	ctx sdk.Context,
	coin sdk.Coin,
) (
	staking, incentives, communityPool sdk.Coins,
	err error,
) {
	params := k.GetParams(ctx)

	currentInflationAmount, errors := sdk.NewIntFromString(k.GetCurrentInflationAmount(ctx))
	if !errors {
		return nil, nil, nil, nil
	}

	maxInflationAmount, errors := sdk.NewIntFromString(k.GetMaxInflationAmount(ctx))
	if !errors {
		return nil, nil, nil, nil
	}
	if currentInflationAmount.Equal(maxInflationAmount) {
		params.EnableInflation = false
		k.SetParams(ctx, params)
		return nil, nil, nil, nil
	}

	if coin.Amount.Add(currentInflationAmount).GT(maxInflationAmount) {
		coin.Amount = maxInflationAmount.Sub(currentInflationAmount)
	}

	// Mint coins for distribution
	if err := k.MintCoins(ctx, coin); err != nil {
		return nil, nil, nil, err
	}

	// Allocate minted coins according to allocation proportions (staking, usage
	// incentives, community pool)
	if _, _, _, err := k.AllocateExponentialInflation(ctx, coin); err != nil {
		return nil, nil, nil, err
	}

	k.SetCurrentInflation(ctx, currentInflationAmount.Add(coin.Amount).String())
	return nil, nil, nil, nil

}

// MintCoins implements an alias call to the underlying supply keeper's
// MintCoins to be used in BeginBlocker.
func (k Keeper) MintCoins(ctx sdk.Context, coin sdk.Coin) error {
	coins := sdk.NewCoins(coin)

	// skip as no coins need to be minted
	if coins.Empty() {
		return nil
	}

	return k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
}

func NewInt(u uint64) {
	panic("unimplemented")
}

// AllocateExponentialInflation allocates coins from the inflation to external
// modules according to allocation proportions:
//   - staking rewards -> sdk `auth` module fee collector
//   - usage incentives -> `x/incentives` module
//   - community pool -> `sdk `distr` module community pool
func (k Keeper) AllocateExponentialInflation(
	ctx sdk.Context,
	mintedCoin sdk.Coin,
) (
	staking, incentives, communityPool sdk.Coins,
	err error,
) {
	params := k.GetParams(ctx)
	proportions := params.InflationDistribution

	// Allocate staking rewards into fee collector account
	staking = sdk.NewCoins(k.GetProportions(ctx, mintedCoin, proportions.StakingRewards))
	err = k.bankKeeper.SendCoinsFromModuleToModule(
		ctx,
		types.ModuleName,
		k.feeCollectorName,
		staking,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	// Allocate usage incentives to incentives module account
	incentives = sdk.NewCoins(k.GetProportions(ctx, mintedCoin, proportions.UsageIncentives))
	err = k.bankKeeper.SendCoinsFromModuleToModule(
		ctx,
		types.ModuleName,
		incentivestypes.ModuleName,
		incentives,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	// Allocate community pool amount (remaining module balance) to community
	// pool address
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	communityPool = k.bankKeeper.GetAllBalances(ctx, moduleAddr)
	err = k.distrKeeper.FundCommunityPool(
		ctx,
		communityPool,
		moduleAddr,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	return staking, incentives, communityPool, nil
}

// GetAllocationProportion calculates the proportion of coins that is to be
// allocated during inflation for a given distribution.
func (k Keeper) GetProportions(
	ctx sdk.Context,
	coin sdk.Coin,
	distribution sdk.Dec,
) sdk.Coin {
	return sdk.NewCoin(
		coin.Denom,
		coin.Amount.ToDec().Mul(distribution).TruncateInt(),
	)
}

// BondedRatio the fraction of the staking tokens which are currently bonded
func (k Keeper) BondedRatio(ctx sdk.Context) sdk.Dec {
	stakeSupply := k.stakingKeeper.StakingTokenSupply(ctx)

	if !stakeSupply.IsPositive() {
		return sdk.ZeroDec()
	}

	return k.stakingKeeper.TotalBondedTokens(ctx).ToDec().QuoInt(stakeSupply)
}

// GetCirculatingSupply returns the bank supply of the mintDenom
func (k Keeper) GetCirculatingSupply(ctx sdk.Context) sdk.Dec {
	mintDenom := k.GetParams(ctx).MintDenom

	circulatingSupply := k.bankKeeper.GetSupply(ctx, mintDenom).Amount.ToDec()

	return circulatingSupply
}

// GetInflationRate returns the inflation rate for the current period.
func (k Keeper) GetInflationRate(ctx sdk.Context) sdk.Dec {
	epochMintProvision, _ := k.GetEpochMintProvision(ctx)
	if epochMintProvision.IsZero() {
		return sdk.ZeroDec()
	}

	epp := k.GetEpochsPerPeriod(ctx)
	if epp == 0 {
		return sdk.ZeroDec()
	}

	epochsPerPeriod := sdk.NewDec(epp)

	circulatingSupply := k.GetCirculatingSupply(ctx)
	if circulatingSupply.IsZero() {
		return sdk.ZeroDec()
	}

	// EpochMintProvision * 365 / circulatingSupply * 100
	return epochMintProvision.Mul(epochsPerPeriod).Quo(circulatingSupply).Mul(sdk.NewDec(100))
}
