package keeper

import (
	"github.com/reapchain/reapchain/v4/x/witness/types"
)

var _ types.QueryServer = Keeper{}
