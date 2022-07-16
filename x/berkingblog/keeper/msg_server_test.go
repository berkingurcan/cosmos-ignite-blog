package keeper_test

import (
	"context"
	"testing"

	keepertest "berkingblog/testutil/keeper"
	"berkingblog/x/berkingblog/keeper"
	"berkingblog/x/berkingblog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BerkingblogKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
