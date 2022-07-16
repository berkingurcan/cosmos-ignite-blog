package keeper_test

import (
	"testing"

	testkeeper "berkingblog/testutil/keeper"
	"berkingblog/x/berkingblog/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BerkingblogKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
