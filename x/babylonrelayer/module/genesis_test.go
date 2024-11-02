package babylonrelayer_test

import (
	"testing"

	keepertest "github.com/AliErcanOzgokce/babylon-relayer/testutil/keeper"
	"github.com/AliErcanOzgokce/babylon-relayer/testutil/nullify"
	babylonrelayer "github.com/AliErcanOzgokce/babylon-relayer/x/babylonrelayer/module"
	"github.com/AliErcanOzgokce/babylon-relayer/x/babylonrelayer/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BabylonrelayerKeeper(t)
	babylonrelayer.InitGenesis(ctx, k, genesisState)
	got := babylonrelayer.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
