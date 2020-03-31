package keeper

import (
	"testing"

	"github.com/commercionetwork/commercionetwork/x/vbr/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

var request abci.RequestQuery

func TestQuerier_getBlockRewardsPoolFunds(t *testing.T) {
	var cdc, ctx, k, _, _ = SetupTestInput(false)
	var querier = NewQuerier(k)

	k.SetTotalRewardPool(ctx, TestBlockRewardsPool)

	path := []string{types.QueryBlockRewardsPoolFunds}
	actual, _ := querier(ctx, path, request)

	expected, _ := codec.MarshalJSONIndent(cdc, &TestBlockRewardsPool)

	require.Equal(t, expected, actual)
}
