package keeper

import (
	"fmt"

	sdk "github.com/reapchain/cosmos-sdk/types"

	incentivestypes "github.com/reapchain/reapchain/v4/x/incentives/types"
	"github.com/reapchain/reapchain/v4/x/inflation/types"
)

// MintAndAllocateInflation performs inflation minting and allocation
func (k Keeper) MintAndAllocateInflation(ctx sdk.Context, coin sdk.Coin) error {
	currentInflationAmount, err := sdk.NewIntFromString(k.GetCurrentInflationAmount(ctx))
	if err == false {
		return nil
	}

	maxInflationAmount, err := sdk.NewIntFromString(k.GetMaxInflationAmount(ctx))
	if err == false {
		return nil
	}

	fmt.Println("stompesi - currentInflationAmount", currentInflationAmount)
	fmt.Println("stompesi - maxInflationAmount", maxInflationAmount)

	if currentInflationAmount.Equal(maxInflationAmount) {
		return nil
	}

	if coin.Amount.Add(currentInflationAmount).GT(maxInflationAmount) {
		coin.Amount = maxInflationAmount.Sub(currentInflationAmount)
	}

	fmt.Println("stompesi - add amount", coin.Amount)

	// Mint coins for distribution
	if err := k.MintCoins(ctx, coin); err != nil {
		return err
	}

	// Allocate minted coins according to allocation proportions (staking, usage
	// incentives, community pool)
	if err := k.AllocateExponentialInflation(ctx, coin); err != nil {
		return err
	}

	k.SetCurrentInflation(ctx, currentInflationAmount.Add(coin.Amount).String())
	return nil
}

// MintCoins implements an alias call to the underlying supply keeper's
// MintCoins to be used in BeginBlocker.
func (k Keeper) MintCoins(ctx sdk.Context, newCoin sdk.Coin) error {
	newCoins := sdk.NewCoins(newCoin)

	// skip as no coins need to be minted
	if newCoins.Empty() {
		return nil
	}

	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}

func NewInt(u uint64) {
	panic("unimplemented")
}

// AllocateExponentialInflation allocates coins from the inflation to external
// modules according to allocation proportions:
//   - staking rewards -> sdk `auth` module fee collector
//   - usage incentives -> `x/incentives` module
//   - community pool -> `sdk `distr` module community pool
func (k Keeper) AllocateExponentialInflation(ctx sdk.Context, mintedCoin sdk.Coin) error {
	params := k.GetParams(ctx)
	proportions := params.InflationDistribution

	// Allocate staking rewards into fee collector account
	stakingRewardsAmt := sdk.NewCoins(k.GetProportions(ctx, mintedCoin, proportions.StakingRewards))
	err := k.bankKeeper.SendCoinsFromModuleToModule(
		ctx,
		types.ModuleName,
		k.feeCollectorName,
		stakingRewardsAmt,
	)
	if err != nil {
		return err
	}

	// Allocate usage incentives to incentives module account
	usageIncentivesAmt := sdk.NewCoins(k.GetProportions(ctx, mintedCoin, proportions.UsageIncentives))
	err = k.bankKeeper.SendCoinsFromModuleToModule(
		ctx,
		types.ModuleName,
		incentivestypes.ModuleName,
		usageIncentivesAmt,
	)
	if err != nil {
		return err
	}

	// Allocate community pool amount (remaining module balance) to community
	// pool address
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	communityPoolAmt := k.bankKeeper.GetAllBalances(ctx, moduleAddr)
	return k.distrKeeper.FundCommunityPool(
		ctx,
		communityPoolAmt,
		moduleAddr,
	)
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
