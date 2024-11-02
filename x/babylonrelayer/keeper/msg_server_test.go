package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/AliErcanOzgokce/babylon-relayer/testutil/keeper"
	"github.com/AliErcanOzgokce/babylon-relayer/x/babylonrelayer/keeper"
	"github.com/AliErcanOzgokce/babylon-relayer/x/babylonrelayer/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.BabylonrelayerKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
