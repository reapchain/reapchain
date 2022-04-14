package types

import (
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
)

// errors
var (
	ErrClaimsRecordNotFound = sdkerrors.Register(ModuleName, 2, "claims record not found")
	ErrInvalidAction        = sdkerrors.Register(ModuleName, 3, "invalid claim action type")
)
