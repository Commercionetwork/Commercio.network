package keeper

import (
	"fmt"
	"strings"
	"testing"

	"github.com/commercionetwork/commercionetwork/x/vbr/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

var msgIncrementsBRPool = types.MsgIncrementBlockRewardsPool{
	Funder: TestFunder,
	Amount: TestAmount,
}

func TestValidMsg_IncrementBRPool(t *testing.T) {
	_, ctx, k, _, bk := SetupTestInput(false)

	_ = bk.SetCoins(ctx, TestFunder, TestAmount)
	handler := NewHandler(k)

	_, err := handler(ctx, msgIncrementsBRPool)
	require.NoError(t, err)

	macc := k.VbrAccount(ctx)

	initialPool, _ := TestBlockRewardsPool.TruncateDecimal()
	expectedTotalAmount := initialPool.Add(TestAmount...)

	require.Equal(t, expectedTotalAmount, macc.GetCoins())
}

func TestInvalidMsg(t *testing.T) {
	_, ctx, k, _, _ := SetupTestInput(false)

	handler := NewHandler(k)

	_, err := handler(ctx, sdk.NewTestMsg())

	require.Error(t, err)
	require.True(t, strings.Contains(err.Error(), fmt.Sprintf("Unrecognized %s message type", types.ModuleName)))
}
