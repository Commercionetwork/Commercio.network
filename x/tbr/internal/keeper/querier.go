package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"

	tbrTypes "github.com/commercionetwork/commercionetwork/x/tbr/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case tbrTypes.QueryBlockRewardsPoolFunds:
			return queryGetBlockRewardsPoolFunds(ctx, path[1:], keeper)
		default:
			return nil, sdk.ErrUnknownRequest(fmt.Sprintf("Unknown %s query endpoint", tbrTypes.ModuleName))
		}
	}
}

func queryGetBlockRewardsPoolFunds(ctx sdk.Context, _ []string, keeper Keeper) (res []byte, err sdk.Error) {
	funds := keeper.GetTotalRewardPool(ctx)

	fundsBz, err2 := codec.MarshalJSONIndent(keeper.cdc, funds)
	if err2 != nil {
		return nil, sdk.ErrUnknownRequest("Could not marshal result to JSON")
	}

	return fundsBz, nil
}
