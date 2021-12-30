package keeper

import (
	"github.com/tharsis/evmos/x/reap/types"
)

var _ types.QueryServer = Keeper{}
