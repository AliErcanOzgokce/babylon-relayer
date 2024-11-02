package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/AliErcanOzgokce/babylon-relayer/testutil/keeper"
	"github.com/AliErcanOzgokce/babylon-relayer/x/babylonrelayer/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BabylonrelayerKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
