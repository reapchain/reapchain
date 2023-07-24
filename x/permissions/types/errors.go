package types

// DONTCOVER

import (
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
)

// x/permissions module sentinel errors
var (
	ErrMaximumNumberOfStandingMember     = sdkerrors.Register(ModuleName, 2, "cannot register anymore validators to whitelist, maximum capacity")
	ErrValidatorAlreadyRegistered        = sdkerrors.Register(ModuleName, 3, "cannot register validator, already in whitelist")
	ErrAccountsDontMatch                 = sdkerrors.Register(ModuleName, 4, "validator address & account address do not match")
	ErrValidatorNotFound                 = sdkerrors.Register(ModuleName, 5, "validator address not found in whitelist")
	ErrInvalidValidatorAddress           = sdkerrors.Register(ModuleName, 6, "validator address invalid")
	ErrUnauthorizedStandingMemberAddress = sdkerrors.Register(ModuleName, 7, "address is not authorized to be a standing member")
	ErrInsufficientInitialDeposit        = sdkerrors.Register(ModuleName, 8, "insufficient initial deposit, please deposit more")
	ErrNotMatchingMonikers               = sdkerrors.Register(ModuleName, 9, "try again with matching moniker")
	ErrInvalidDelegation                 = sdkerrors.Register(ModuleName, 10, "validator not in whitelist, delegation is not possible")
)
