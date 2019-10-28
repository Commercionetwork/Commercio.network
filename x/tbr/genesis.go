package tbr

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GenesisState struct {
	PoolAmount       sdk.DecCoins `json:"pool_amount"`
	YearlyPoolAmount sdk.DecCoins `json:"yearly_pool_amount"`
	YearNumber       int64        `json:"year_number"`
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// InitGenesis sets the initial Block Reward Pool amount for genesis.
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	// Set the reward pool - Should never be nil as its validated inside the ValidateGenesis method
	keeper.SetTotalRewardPool(ctx, data.PoolAmount)

	// Default yearly reward pool amount
	if data.YearlyPoolAmount == nil {
		data.YearlyPoolAmount = data.PoolAmount.QuoDec(sdk.NewDec(5))
	}

	// Set the yearly reward pool and year number
	keeper.SetYearlyRewardPool(ctx, data.YearlyPoolAmount)
	keeper.SetYearNumber(ctx, data.YearNumber)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	return GenesisState{
		PoolAmount:       keeper.GetTotalRewardPool(ctx),
		YearlyPoolAmount: keeper.GetYearlyRewardPool(ctx),
		YearNumber:       keeper.GetYearNumber(ctx),
	}
}

// ValidateGenesis performs basic validation of genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error {
	if data.PoolAmount == nil || data.PoolAmount.Empty() {
		return errors.New("transaction block reward pool cannot be empty")
	}

	if !data.PoolAmount.IsValid() {
		return fmt.Errorf("invalid transaction block reward pool: %s", data.PoolAmount.String())
	}

	if !data.YearlyPoolAmount.IsValid() {
		return fmt.Errorf("invalid yearly transaction block reward pool: %s", data.YearlyPoolAmount.String())
	}

	return nil
}
