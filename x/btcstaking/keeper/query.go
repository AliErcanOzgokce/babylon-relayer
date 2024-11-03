package keeper

import (
	"github.com/AliErcanOzgokce/babylon-relayer/x/btcstaking/types"
)

var _ types.QueryServer = Keeper{}
