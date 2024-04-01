package tt1_test

import (
	"testing"

	keepertest "tt1/testutil/keeper"
	"tt1/testutil/nullify"
	tt1 "tt1/x/tt1/module"
	"tt1/x/tt1/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Tt1Keeper(t)
	tt1.InitGenesis(ctx, k, genesisState)
	got := tt1.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
