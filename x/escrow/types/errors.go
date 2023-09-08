package types

import (
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
)

// errors
var (
	ErrEscrowDisabled       = sdkerrors.Register(ModuleName, 2, "escrow module is disabled")
	ErrDenomNotFound        = sdkerrors.Register(ModuleName, 3, "denomination not found")
	ErrDenomAlreadyExists   = sdkerrors.Register(ModuleName, 4, "denomination already exists")
	ErrDenominationDisabled = sdkerrors.Register(ModuleName, 5, "escrow conversion for denomination is disabled")
)
