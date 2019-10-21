package keeper

import (
	"fmt"
	"time"

	"github.com/commercionetwork/commercionetwork/x/mint/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case types.QueryGetCdp:
			return queryGetCdp(ctx, path[1:], keeper)
		case types.QueryGetCdps:
			return queryGetCdps(ctx, path[1:], keeper)
		default:
			return nil, sdk.ErrUnknownRequest(fmt.Sprintf("Unknown %s query endpoint", types.ModuleName))
		}
	}
}

func queryGetCdp(ctx sdk.Context, path []string, keeper Keeper) ([]byte, sdk.Error) {
	ownerAddr, _ := sdk.AccAddressFromBech32(path[0])
	timestamp, err := time.Parse(time.RFC3339, path[1])
	if err != nil {
		return nil, sdk.ErrUnknownRequest("timestamp format isn't ISO8601")
	}
	cdp, found := keeper.GetCdpByOwnerAndTimeStamp(ctx, ownerAddr, timestamp)

	if !found {
		return nil, sdk.ErrUnknownRequest("couldn't find any cdp associated with the given address and timestamp")
	}

	cdpBz, err := codec.MarshalJSONIndent(keeper.cdc, &cdp)
	if err != nil {
		return nil, sdk.ErrUnknownRequest("Could not marshal result to JSON")
	}

	return cdpBz, nil
}

func queryGetCdps(ctx sdk.Context, path []string, keeper Keeper) ([]byte, sdk.Error) {
	ownerAddr, _ := sdk.AccAddressFromBech32(path[0])
	cdps := keeper.GetCdpsByOwner(ctx, ownerAddr)

	cdpsBz, err := codec.MarshalJSONIndent(keeper.cdc, cdps)
	if err != nil {
		return nil, sdk.ErrUnknownRequest("Could not marshal result to JSON")
	}

	return cdpsBz, nil
}
