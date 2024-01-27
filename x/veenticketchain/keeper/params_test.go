package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "veenticketchain/testutil/keeper"
	"veenticketchain/x/veenticketchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.VeenticketchainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
