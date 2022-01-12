package keeper

import (
	"github.com/reapchain/mercury/x/reap/types"
)

var _ types.QueryServer = Keeper{}
