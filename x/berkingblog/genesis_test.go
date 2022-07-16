package berkingblog_test

import (
	"testing"

	keepertest "berkingblog/testutil/keeper"
	"berkingblog/testutil/nullify"
	"berkingblog/x/berkingblog"
	"berkingblog/x/berkingblog/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BerkingblogKeeper(t)
	berkingblog.InitGenesis(ctx, *k, genesisState)
	got := berkingblog.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
