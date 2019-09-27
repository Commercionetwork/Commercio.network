package pricefeed

import (
	"errors"
	"fmt"
	"strings"

	ctypes "github.com/commercionetwork/commercionetwork/x/common/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - docs genesis state
type GenesisState struct {
	Oracles ctypes.Addresses
	Assets  ctypes.Strings
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	return GenesisState{Oracles: ctypes.Addresses{}, Assets: ctypes.Strings{}}
}

// InitGenesis sets docs information for genesis.
func InitGenesis(ctx sdk.Context, keeper Keeper, genState GenesisState) {
	for _, oracle := range genState.Oracles {
		keeper.AddOracle(ctx, oracle)
	}

	for _, asset := range genState.Assets {
		keeper.AddAsset(ctx, asset)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {

	oracles := keeper.GetOracles(ctx)
	assets := keeper.GetAssets(ctx)

	return GenesisState{
		Oracles: oracles,
		Assets:  assets,
	}
}

// ValidateGenesis performs basic validation of genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(state GenesisState) error {
	for _, asset := range state.Assets {
		if len(strings.TrimSpace(asset)) == 0 {
			return errors.New(fmt.Sprintf("%s, is empty", asset))
		}
	}
	for _, oracle := range state.Oracles {
		if oracle.Empty() {
			return sdk.ErrInvalidAddress("Found Empty oracle address")
		}
	}
	return nil
}
