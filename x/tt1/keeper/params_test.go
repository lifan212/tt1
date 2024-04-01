package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "tt1/testutil/keeper"
	"tt1/x/tt1/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.Tt1Keeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
