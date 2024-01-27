package keeper

import (
	"veenticketchain/x/veenticketchain/types"
)

var _ types.QueryServer = Keeper{}
