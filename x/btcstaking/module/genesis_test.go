package btcstaking_test

import (
	"testing"

	keepertest "github.com/AliErcanOzgokce/babylon-relayer/testutil/keeper"
	"github.com/AliErcanOzgokce/babylon-relayer/testutil/nullify"
	btcstaking "github.com/AliErcanOzgokce/babylon-relayer/x/btcstaking/module"
	"github.com/AliErcanOzgokce/babylon-relayer/x/btcstaking/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BtcstakingKeeper(t)
	btcstaking.InitGenesis(ctx, k, genesisState)
	got := btcstaking.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
