package veenticketchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "veenticketchain/testutil/keeper"
	"veenticketchain/testutil/nullify"
	"veenticketchain/x/veenticketchain/module"
	"veenticketchain/x/veenticketchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VeenticketchainKeeper(t)
	veenticketchain.InitGenesis(ctx, k, genesisState)
	got := veenticketchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
