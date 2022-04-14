package types

import (
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
)

// errors
var (
	ErrInsufficientVestedCoins = sdkerrors.Register(ModuleName, 2, "insufficient vested coins error")
	ErrVestingLockup           = sdkerrors.Register(ModuleName, 3, "vesting lockup error")
)
