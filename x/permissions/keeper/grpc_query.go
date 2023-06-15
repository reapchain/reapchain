package keeper

import (
	"github.com/reapchain/reapchain/v8/x/permissions/types"
)

var _ types.QueryServer = Keeper{}
