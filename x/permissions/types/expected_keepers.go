package types

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/reapchain/cosmos-sdk/x/staking/types"
	"time"
)

type StakingKeeper interface {
	Validator(sdk.Context, sdk.ValAddress) stakingtypes.ValidatorI
	Unjail(sdk.Context, sdk.ConsAddress)

	UnbondingTime(ctx sdk.Context) (res time.Duration)

	GetUnbondingDelegation(
		ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress,
	) (ubd stakingtypes.UnbondingDelegation, found bool)

	GetAllValidators(ctx sdk.Context) (validators []stakingtypes.Validator)

	GetValidatorDelegations(ctx sdk.Context, valAddr sdk.ValAddress) (delegations []stakingtypes.Delegation)

	Undelegate(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, sharesAmount sdk.Dec) (time.Time, error)

	Delegation(sdk.Context, sdk.AccAddress, sdk.ValAddress) stakingtypes.DelegationI

	GetDelegation(ctx sdk.Context,
		delAddr sdk.AccAddress, valAddr sdk.ValAddress) (delegation stakingtypes.Delegation, found bool)

	//Interface Functinos for Force Unbonding
	MaxStandingMembers(ctx sdk.Context) (res uint32)
	BondDenom(ctx sdk.Context) (res string)
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)

	HasMaxUnbondingDelegationEntries(ctx sdk.Context,
		delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress) bool

	Unbond(
		ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, shares sdk.Dec,
	) (amount sdk.Int, err error)

	SetUnbondingDelegationEntry(
		ctx sdk.Context, delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress,
		creationHeight int64, minTime time.Time, balance sdk.Int,
	) stakingtypes.UnbondingDelegation

	InsertUBDQueue(ctx sdk.Context, ubd stakingtypes.UnbondingDelegation,
		completionTime time.Time)

	stakingtypes.StakingHooks
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SendCoinsFromModuleToModule(ctx sdk.Context, senderPool, recipientPool string, amt sdk.Coins) error
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

//// GovKeeper defines the expected governance keeper interface used on incentives
//type GovKeeper interface {
//	GetDepositParams(ctx sdk.Context) govtypes.DepositParams
//	GetVotingParams(ctx sdk.Context) govtypes.VotingParams
//}
