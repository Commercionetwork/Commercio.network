package keeper

import (
	"errors"

	"github.com/commercionetwork/commercionetwork/x/government/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	StoreKey sdk.StoreKey
	cdc      *codec.Codec
}

func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		StoreKey: storeKey,
		cdc:      cdc,
	}
}

// SetGovernmentAddress allows to set the given address as the one that
// the government will use later
func (keeper Keeper) SetGovernmentAddress(ctx sdk.Context, address sdk.AccAddress) error {
	store := ctx.KVStore(keeper.StoreKey)

	if store.Has([]byte(types.GovernmentStoreKey)) {
		return errors.New("government address already set")
	}

	store.Set([]byte(types.GovernmentStoreKey), address)
	return nil
}

// GetGovernmentAddress returns the address that the government has currently
func (keeper Keeper) GetGovernmentAddress(ctx sdk.Context) sdk.AccAddress {
	store := ctx.KVStore(keeper.StoreKey)
	return store.Get([]byte(types.GovernmentStoreKey))
}

// SetTumblerAddress allows to set the given address as the one that
// the tumbler will use later
func (keeper Keeper) SetTumblerAddress(ctx sdk.Context, address sdk.AccAddress) error {
	store := ctx.KVStore(keeper.StoreKey)

	store.Set([]byte(types.TumblerStoreKey), address)
	return nil
}

// GetTumblerAddress returns the address that the tumbler has currently
func (keeper Keeper) GetTumblerAddress(ctx sdk.Context) sdk.AccAddress {
	store := ctx.KVStore(keeper.StoreKey)
	return store.Get([]byte(types.TumblerStoreKey))
}
