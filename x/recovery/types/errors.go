package types

import (
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
)

// errors
var (
	ErrBlockedAddress = sdkerrors.Register(ModuleName, 2, "blocked address")
)
