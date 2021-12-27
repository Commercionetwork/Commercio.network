package keeper

import (
	"fmt"

	sdkErr "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/commercionetwork/commercionetwork/x/commerciomint/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryGetEtp:
			return queryGetEtp(ctx, path[1:], k, legacyQuerierCdc)
		case types.QueryGetEtpsByOwner:
			return queryGetEtpsByOwner(ctx, path[1:], k, legacyQuerierCdc)
		case types.QueryGetallEtps:
			return queryGetAllEtp(ctx, k, legacyQuerierCdc)
		case types.QueryConversionRateRest:
			return queryConversionRate(ctx, k, legacyQuerierCdc)
		case types.QueryFreezePeriodRest:
			return queryFreezePeriod(ctx, k,legacyQuerierCdc)
		default:
			return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest, fmt.Sprintf("Unknown %s query endpoint", types.ModuleName))
		}
	}
}

func queryGetEtp(ctx sdk.Context, path []string, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	id := path[0]
	etp, ok := k.GetPositionById(ctx, id)
	if !ok {
		return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest, fmt.Sprintf("Posistion with id: %s not found!", id))
	}

	etpbz, err := codec.MarshalJSONIndent(legacyQuerierCdc, etp)
	if err != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest, "could not marshal result to JSON")
	}

	return etpbz, nil
}

func queryGetEtpsByOwner(ctx sdk.Context, path []string, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	ownerAddr, e := sdk.AccAddressFromBech32(path[0])
	if e != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrInvalidAddress, fmt.Sprintf("The given address: %s is not valid!", path[0]))
	}
	etps := k.GetAllPositionsOwnedBy(ctx, ownerAddr)
	etpsBz, err := codec.MarshalJSONIndent(legacyQuerierCdc, etps)
	if err != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest, "could not marshal result to JSON")
	}

	return etpsBz, nil
}

func queryGetAllEtp(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	etps := k.GetAllPositions(ctx)
	etpsBz, err := codec.MarshalJSONIndent(legacyQuerierCdc, etps)
	if err != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest, "could not marshal result to JSON")
	}

	return etpsBz, nil
}

func queryConversionRate(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	return codec.MarshalJSONIndent(legacyQuerierCdc, k.GetConversionRate(ctx))
}

func queryFreezePeriod(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	return codec.MarshalJSONIndent(legacyQuerierCdc, k.GetFreezePeriod(ctx))
}
