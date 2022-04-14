package keeper

import (
	"github.com/reapchain/reapchain/x/reap/types"
)

var _ types.QueryServer = Keeper{}
