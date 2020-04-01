package keeper

import (
	"fmt"

	sdkErr "github.com/cosmos/cosmos-sdk/types/errors"

	ctypes "github.com/commercionetwork/commercionetwork/x/common/types"
	"github.com/commercionetwork/commercionetwork/x/vbr/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/staking/exported"
	"github.com/cosmos/cosmos-sdk/x/supply"
	supplyExported "github.com/cosmos/cosmos-sdk/x/supply/exported"
)

type Keeper struct {
	cdc          *codec.Codec
	storeKey     sdk.StoreKey
	distKeeper   distribution.Keeper
	supplyKeeper supply.Keeper
}

func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, dk distribution.Keeper, sk supply.Keeper) Keeper {
	return Keeper{
		cdc:          cdc,
		storeKey:     storeKey,
		distKeeper:   dk,
		supplyKeeper: sk,
	}
}

// -------------
// --- Pool
// -------------

// SetTotalRewardPool allows to set the value of the total rewards pool that has left
func (k Keeper) SetTotalRewardPool(ctx sdk.Context, updatedPool sdk.DecCoins) {
	store := ctx.KVStore(k.storeKey)
	if !updatedPool.Empty() {
		store.Set([]byte(types.PoolStoreKey), k.cdc.MustMarshalBinaryBare(&updatedPool))
	} else {
		store.Delete([]byte(types.PoolStoreKey))
	}
}

// GetTotalRewardPool returns the current total rewards pool amount
func (k Keeper) GetTotalRewardPool(ctx sdk.Context) sdk.DecCoins {
	macc := k.supplyKeeper.GetModuleAccount(ctx, types.ModuleName)
	mcoins := macc.GetCoins()

	return sdk.NewDecCoinsFromCoins(mcoins...)
}

// --------------------------
// --- Yearly reward pool
// --------------------------

// GetYearlyRewardPool returns the reward pool that has been assigned for the current year or rewards
func (k Keeper) GetYearlyRewardPool(ctx sdk.Context) (pool sdk.DecCoins) {
	store := ctx.KVStore(k.storeKey)
	k.cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.YearlyPoolStoreKey)), &pool)
	return pool
}

// SetYearlyRewardPool sets the given yearlyPool to be the current year's reward pool
func (k Keeper) SetYearlyRewardPool(ctx sdk.Context, yearlyPool sdk.DecCoins) {
	store := ctx.KVStore(k.storeKey)
	if !yearlyPool.Empty() {
		store.Set([]byte(types.YearlyPoolStoreKey), k.cdc.MustMarshalBinaryBare(&yearlyPool))
	} else {
		store.Delete([]byte(types.YearlyPoolStoreKey))
	}
}

// --------------------
// --- Year number
// --------------------

var (
	DPY = sdk.NewDecWithPrec(36525, 2) // Days Per Year
	HPD = sdk.NewDecWithPrec(24, 0)    //  Hours Per Day
	MPH = sdk.NewDecWithPrec(60, 0)    //  Minutes Per Hour
	BPM = sdk.NewDecWithPrec(12, 0)    // Blocks Per Minutes

	BPY = DPY.Mul(HPD).Mul(MPH).Mul(BPM) // Blocks Per Year
)

func computeYearFromBlockHeight(blockHeight int64) int64 {
	// Divide the current block number to the number of blocks per year to get the year value
	// Truncate the result so that 1.99 years = 1 year and not 2
	blocksPerYear := DPY.Mul(HPD).Mul(MPH).Mul(BPM)
	return sdk.NewDec(blockHeight).Quo(blocksPerYear).TruncateInt64()
}

// SetYearNumber set the given year to be the current year number
func (k Keeper) SetYearNumber(ctx sdk.Context, year int64) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(types.YearNumberStoreKey), k.cdc.MustMarshalBinaryBare(year))
}

// GetYearNumber returns the current year number
func (k Keeper) GetYearNumber(ctx sdk.Context) (year int64) {
	store := ctx.KVStore(k.storeKey)
	actualBz := store.Get([]byte(types.YearNumberStoreKey))
	if actualBz == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryBare(actualBz, &year)
	return year
}

// UpdateYearlyPool allows to update the current yearly pool based on whenever we
// are now in a new year or not.
// If we are in the same year, we do nothing. Otherwise, we set the new year's pool to be the
// 20% of the total reward pool.
func (k Keeper) UpdateYearlyPool(ctx sdk.Context, blockHeight int64) {
	previousYearNumber := k.GetYearNumber(ctx)
	currentYearNumber := computeYearFromBlockHeight(blockHeight)

	// Check if the year number has changed and thus we need to update the yearly reward pool
	if previousYearNumber != currentYearNumber {

		// Get the reward pool
		rewardPool := k.GetTotalRewardPool(ctx)

		// Compute a pool in which each coin is 20% of the total pool
		yearlyRewardPool := make(sdk.DecCoins, len(rewardPool))
		for index, coin := range rewardPool {
			// Each new coin amount must be 20% (= 1/5) of the previous
			yearlyRewardPool[index] = sdk.NewDecCoinFromDec(coin.Denom, coin.Amount.QuoInt64(5))
		}

		// Set the new yearly reward pool and year number
		k.SetYearlyRewardPool(ctx, yearlyRewardPool)
		k.SetYearNumber(ctx, currentYearNumber)
	}
}

// ---------------------------
// --- Reward distribution
// ---------------------------

// ComputeProposerReward computes the final reward for the validator block's proposer
func (k Keeper) ComputeProposerReward(ctx sdk.Context, validatorsCount int64,
	proposer exported.ValidatorI, totalStakedTokens sdk.Int) sdk.DecCoins {

	// Get the maximum year reward by multiplying the yearly pool by V/100
	Ry := k.GetYearlyRewardPool(ctx).MulDec(sdk.NewDec(validatorsCount)).QuoDec(sdk.NewDec(100))

	// Cap the yearly reward limit per validator by dividing the yearly reward by 100
	RLyn := Ry.QuoDec(sdk.NewDec(100))

	// Compute the voting power for this validator at the current block
	VPnb := proposer.GetBondedTokens().Quo(totalStakedTokens)

	// Compute the half validator set
	halfV := sdk.NewDec(1).QuoInt64(validatorsCount)
	isTopValidator := VPnb.ToDec().GT(halfV)

	// Compute the multiplying factor based on whenever the validator is in the top half or not.
	// If it's in the top half list, the validator should receive a lower quantity of tokens as it
	// will validate more blocks.
	// If it's in the bottom half of the list it should receive a higher amount as it will validate
	// less blocks
	var multiplyingFactor sdk.Dec
	if isTopValidator {
		multiplyingFactor = sdk.NewDec(1).QuoInt(VPnb)
	} else {
		multiplyingFactor = sdk.NewDec(validatorsCount)
	}

	// Compute the final reward
	Rnb := RLyn.QuoDec(BPY).MulDec(multiplyingFactor)

	return Rnb
}

// DistributeBlockRewards distributes the computed reward to the block proposer
func (k Keeper) DistributeBlockRewards(ctx sdk.Context, validator exported.ValidatorI, reward sdk.DecCoins) error {
	rewardPool := k.GetTotalRewardPool(ctx)
	yearlyPool := k.GetYearlyRewardPool(ctx)

	// Check if the yearly pool and the total pool have enough funds
	if ctypes.IsAllGTE(rewardPool, reward) && ctypes.IsAllGTE(yearlyPool, reward) {
		// truncate fractional part and only take the integer part into account
		rewardInt, _ := reward.TruncateDecimal()
		k.SetYearlyRewardPool(ctx, yearlyPool.Sub(sdk.NewDecCoinsFromCoins(rewardInt...)))

		err := k.supplyKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, distribution.ModuleName, rewardInt)
		if err != nil {
			return fmt.Errorf("could not send tokens from vbr to distribution module accounts: %w", err)
		}
		k.distKeeper.AllocateTokensToValidator(ctx, validator, sdk.NewDecCoinsFromCoins(rewardInt...))
	} else {
		return sdkErr.Wrap(sdkErr.ErrInsufficientFunds, "Pool hasn't got enough funds to supply validator's rewards")
	}

	return nil
}

// VbrAccount returns vbr's ModuleAccount
func (k Keeper) VbrAccount(ctx sdk.Context) supplyExported.ModuleAccountI {
	return k.supplyKeeper.GetModuleAccount(ctx, types.ModuleName)
}

// MintVBRTokens mints coins into the vbr's ModuleAccount
func (k Keeper) MintVBRTokens(ctx sdk.Context, coins sdk.Coins) error {
	if err := k.supplyKeeper.MintCoins(ctx, types.ModuleName, coins); err != nil {
		return fmt.Errorf("could not mint requested coins: %w", err)
	}

	return nil
}
