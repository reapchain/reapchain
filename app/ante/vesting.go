package ante

import (
	"fmt"
	"github.com/reapchain/cosmos-sdk/codec"
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	"github.com/reapchain/cosmos-sdk/x/authz"
	stakingtypes "github.com/reapchain/cosmos-sdk/x/staking/types"
	evmtypes "github.com/reapchain/ethermint/x/evm/types"
	vestingtypes "github.com/reapchain/reapchain/v8/x/vesting/types"
	"time"
)

// EthVestingTransactionDecorator validates if clawback vesting accounts are
// permitted to perform Ethereum Tx.
type EthVestingTransactionDecorator struct {
	ak evmtypes.AccountKeeper
}

func NewEthVestingTransactionDecorator(ak evmtypes.AccountKeeper) EthVestingTransactionDecorator {
	return EthVestingTransactionDecorator{
		ak: ak,
	}
}

// AnteHandle validates that a clawback vesting account has surpassed the
// vesting cliff and lockup period.
//
// This AnteHandler decorator will fail if:
//   - the message is not a MsgEthereumTx
//   - sender account cannot be found
//   - sender account is not a ClawbackvestingAccount
//   - blocktime is before surpassing vesting cliff end (with zero vested coins) AND
//   - blocktime is before surpassing all lockup periods (with non-zero locked coins)
func (vtd EthVestingTransactionDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
				"invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil),
			)
		}

		acc := vtd.ak.GetAccount(ctx, msgEthTx.GetFrom())
		if acc == nil {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress,
				"account %s does not exist", acc)
		}

		// Check that this decorator only applies to clawback vesting accounts
		clawbackAccount, isClawback := acc.(*vestingtypes.ClawbackVestingAccount)
		if !isClawback {
			return next(ctx, tx, simulate)
		}

		// Error if vesting cliff has not passed (with zero vested coins). This
		// rule does not apply for existing clawback accounts that receive a new
		// grant while there are already vested coins on the account.
		vested := clawbackAccount.GetVestedCoins(ctx.BlockTime())
		if len(vested) == 0 {
			return ctx, sdkerrors.Wrapf(vestingtypes.ErrInsufficientVestedCoins,
				"cannot perform Ethereum tx with clawback vesting account, that has no vested coins: %s", vested,
			)
		}

		// Error if account has locked coins (before surpassing all lockup periods)
		islocked := clawbackAccount.HasLockedCoins(ctx.BlockTime())
		if islocked {
			return ctx, sdkerrors.Wrapf(vestingtypes.ErrVestingLockup,
				"cannot perform Ethereum tx with clawback vesting account, that has locked coins: %s", vested,
			)
		}
	}

	return next(ctx, tx, simulate)
}

// TODO: remove once Cosmos SDK is upgraded to v0.46

// VestingDelegationDecorator validates delegation of vested coins
type VestingDelegationDecorator struct {
	ak  evmtypes.AccountKeeper
	sk  vestingtypes.StakingKeeper
	cdc codec.BinaryCodec
	bk  evmtypes.BankKeeper
}

// NewVestingDelegationDecorator creates a new VestingDelegationDecorator
func NewVestingDelegationDecorator(ak evmtypes.AccountKeeper, bk evmtypes.BankKeeper, sk vestingtypes.StakingKeeper, cdc codec.BinaryCodec) VestingDelegationDecorator {
	return VestingDelegationDecorator{
		ak:  ak,
		sk:  sk,
		cdc: cdc,
		bk:  bk,
	}
}

// AnteHandle checks if the tx contains a staking delegation.
// It errors if the coins are still locked or the bond amount is greater than
// the coins already vested
func (vdd VestingDelegationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		switch msg := msg.(type) {
		case *authz.MsgExec:
			// Check for bypassing authorization
			if err := vdd.validateAuthz(ctx, msg); err != nil {
				return ctx, err
			}
		default:
			if err := vdd.validateMsg(ctx, msg); err != nil {
				return ctx, err
			}
		}
	}

	return next(ctx, tx, simulate)
}

// validateAuthz validates the authorization internal message
func (vdd VestingDelegationDecorator) validateAuthz(ctx sdk.Context, execMsg *authz.MsgExec) error {
	for _, v := range execMsg.Msgs {
		var innerMsg sdk.Msg
		if err := vdd.cdc.UnpackAny(v, &innerMsg); err != nil {
			return sdkerrors.Wrap(err, "cannot unmarshal authz exec msgs")
		}

		if err := vdd.validateMsg(ctx, innerMsg); err != nil {
			return err
		}
	}

	return nil
}

// validateMsg checks that the only vested coins can be delegated
func (vdd VestingDelegationDecorator) validateMsg(ctx sdk.Context, msg sdk.Msg) error {
	delegateMsg, ok := msg.(*stakingtypes.MsgDelegate)
	if !ok {
		return nil
	}

	for _, addr := range msg.GetSigners() {
		acc := vdd.ak.GetAccount(ctx, addr)
		if acc == nil {
			return sdkerrors.Wrapf(
				sdkerrors.ErrUnknownAddress,
				"account %s does not exist", addr,
			)
		}

		clawbackAccount, isClawback := acc.(*vestingtypes.ClawbackVestingAccount)
		if !isClawback {
			// continue to next decorator as this logic only applies to vesting
			return nil
		}

		// error if bond amount is > vested coins
		bondDenom := vdd.sk.BondDenom(ctx)
		coins := clawbackAccount.GetVestedOnly(ctx.BlockTime())

		// Commented Vesting Balance only Check for Staking
		// This code only allows for Vesting Accounts to Stake their fully vested coins, not including any already owned coins.
		//if coins == nil || coins.Empty() {
		//	return sdkerrors.Wrap(
		//		vestingtypes.ErrInsufficientVestedCoins,
		//		"account has no vested coins",
		//	)
		//}

		vested := coins.AmountOf(bondDenom)

		balance := vdd.bk.GetBalance(ctx, acc.GetAddress(), bondDenom)

		currentlyVestingAmounts := clawbackAccount.GetVestingCoins(ctx.BlockTime())

		for _, currentVestingCoin := range currentlyVestingAmounts {
			balanceWithoutCurrentVesting := balance.Sub(currentVestingCoin)
			fmt.Println("CURRENTLY VESTING AMOUNT: ", currentVestingCoin.Amount)

			if balanceWithoutCurrentVesting.Amount.LT(delegateMsg.Amount.Amount) {
				return sdkerrors.Wrapf(
					vestingtypes.ErrInsufficientVestedCoins,
					"cannot delegate unvested coins. coins vested < delegation amount (%s < %s)",
					vested, delegateMsg.Amount.Amount,
				)
			}
		}

		fmt.Printf("\n\n")
		fmt.Println("==================================================")
		fmt.Println("LOG OUTPUT - REAPCHAIN SOURCE CODE", time.Now().Format(time.RFC822))
		fmt.Println("ADDRESS", acc.GetAddress().String())
		fmt.Println("BALANCE", balance)
		fmt.Println("==================================================")
		fmt.Printf("\n\n")

	}

	return nil
}
