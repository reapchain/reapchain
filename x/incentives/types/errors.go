package types

import (
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
)

// errors
var (
	ErrInternalIncentive = sdkerrors.Register(ModuleName, 2, "internal incentives error")
)
