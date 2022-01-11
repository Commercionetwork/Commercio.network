package keeper

import (
	"context"
	"testing"

	"github.com/commercionetwork/commercionetwork/x/documents/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context, Keeper) {
	keeper, ctx := setupKeeper(t)
	return NewMsgServerImpl(*keeper), sdk.WrapSDKContext(ctx), *keeper
}
