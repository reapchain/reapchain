package types

// DONTCOVER

import (
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
)

// x/permissions module sentinel errors
var (
	ErrMaximumNumberOfStandingMember     = sdkerrors.Register(ModuleName, 10, "\n\ncannot register anymore validators to whitelist, maximum capacity\n\n")
	ErrValidatorAlreadyRegistered        = sdkerrors.Register(ModuleName, 11, "\n\ncannot register validator, already in whitelist\n\n")
	ErrAccountsDontMatch                 = sdkerrors.Register(ModuleName, 12, "\n\nvalidator address & account address are not from the same keys\n\n")
	ErrValidatorNotFound                 = sdkerrors.Register(ModuleName, 13, "\n\nvalidator address not found in whitelist\n\n")
	ErrInvalidValidatorAddress           = sdkerrors.Register(ModuleName, 14, "\n\nvalidator address invalid\n\n")
	ErrUnauthorizedStandingMemberAddress = sdkerrors.Register(ModuleName, 15, "\n\naddresss is not authorized to be a standing member\n\n")
	ErrInsufficientInitialDeposit        = sdkerrors.Register(ModuleName, 16, "\n\ninsufficient initial deposit, please deposit more.\n\n")
	ErrNotMatchingMonikers               = sdkerrors.Register(ModuleName, 15, "\n\ntry again with matching moniker.\n\n")
)
