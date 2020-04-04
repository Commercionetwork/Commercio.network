package keeper

import (
	"fmt"
	"testing"

	sdkErr "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/stretchr/testify/require"

	"github.com/commercionetwork/commercionetwork/x/commerciomint/types"
	"github.com/commercionetwork/commercionetwork/x/pricefeed"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestKeeper_StoreCdp(t *testing.T) {
	ctx, bk, _, _, k := SetupTestInput()
	//handler := NewHandler(k)

	_, _ = bk.AddCoins(ctx, k.supplyKeeper.GetModuleAddress(types.ModuleName), sdk.NewCoins(testCdp.Deposit))
	_ = bk.SetCoins(ctx, testCdp.Owner, testCdp.Credits)
	require.Equal(t, 0, len(k.GetCdps(ctx)))
	k.StoreCdp(ctx, testCdp)
	require.Equal(t, 1, len(k.GetCdps(ctx)))
	cdp, found := k.GetCdp(ctx, testCdp.Owner, testCdp.CreatedAt)
	require.True(t, found)
	require.Equal(t, testCdp.Owner, cdp.Owner)
	require.Equal(t, testCdp.CreatedAt, cdp.CreatedAt)
}

// --------------
// --- Credits
// --------------

func TestKeeper_SetCreditsDenom(t *testing.T) {
	ctx, _, _, _, k := SetupTestInput()
	denom := "test"
	k.SetCreditsDenom(ctx, denom)

	store := ctx.KVStore(k.storeKey)
	denomBz := store.Get([]byte(types.CreditsDenomStoreKey))
	require.Equal(t, denom, string(denomBz))
}

func TestKeeper_GetCreditsDenom(t *testing.T) {
	ctx, _, _, _, k := SetupTestInput()
	denom := "test"
	k.SetCreditsDenom(ctx, denom)
	actual := k.GetCreditsDenom(ctx)
	require.Equal(t, denom, actual)
}

// --------------
// --- CDPs
// --------------

func TestKeeper_AddCdp(t *testing.T) {
	testData := []struct {
		name             string
		cdps             types.Cdps
		newCdp           types.Cdp
		shouldBeInserted bool
	}{
		{
			name:             "Existing CDP is not inserted",
			cdps:             types.Cdps{testCdp},
			newCdp:           testCdp,
			shouldBeInserted: false,
		},
		{
			name:             "New CDP is inserted properly",
			cdps:             types.Cdps{},
			newCdp:           testCdp,
			shouldBeInserted: true,
		},
	}

	for _, test := range testData {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, _, _, _, k := SetupTestInput()

			for _, cdp := range test.cdps {
				k.AddCdp(ctx, cdp)
			}

			k.AddCdp(ctx, test.newCdp)

			result := k.GetCdps(ctx)

			if test.shouldBeInserted {
				require.Len(t, result, len(test.cdps)+1)
			} else {
				require.Len(t, result, len(test.cdps))
			}
		})
	}
}

func TestKeeper_OpenCdp(t *testing.T) {
	testData := []struct {
		name            string
		owner           sdk.AccAddress
		amount          sdk.Coin
		tokenPrice      pricefeed.Price
		userFunds       sdk.Coins
		error           error
		returnedCredits sdk.Coins
	}{
		{
			name:       "Invalid deposited amount",
			owner:      testCdp.Owner,
			amount:     sdk.NewInt64Coin("testcoin", 0),
			tokenPrice: pricefeed.EmptyPrice(),
			error: sdkErr.Wrap(sdkErr.ErrInvalidCoins, fmt.Sprintf(
				"Invalid deposit amount: %s",
				sdk.NewCoins(sdk.NewInt64Coin("testcoin", 0)),
			)),
		},
		{
			name:       "Token price not found",
			owner:      testCdp.Owner,
			amount:     testCdp.Deposit,
			tokenPrice: pricefeed.EmptyPrice(),
			error:      sdkErr.Wrap(sdkErr.ErrUnknownRequest, fmt.Sprintf("no current price for given denom: %s", testCdp.Deposit.Denom)),
		},
		{
			name:       "Not enough funds inside user wallet",
			amount:     testCdp.Deposit,
			owner:      testCdp.Owner,
			tokenPrice: pricefeed.NewPrice("ucommercio", sdk.NewDec(10), sdk.NewInt(1000)),
			error: sdkErr.Wrap(sdkErr.ErrInsufficientFunds, fmt.Sprintf(
				"insufficient account funds; %s < %s",
				sdk.Coins{},
				sdk.NewCoins(sdk.NewInt64Coin("ucommercio", 100)),
			)),
		},
		{
			name:            "Successful opening",
			amount:          testCdp.Deposit,
			owner:           testCdp.Owner,
			tokenPrice:      pricefeed.NewPrice(testLiquidityDenom, sdk.NewDec(10), sdk.NewInt(1000)),
			userFunds:       sdk.NewCoins(testCdp.Deposit),
			returnedCredits: sdk.NewCoins(sdk.NewInt64Coin(testCreditsDenom, 10*50)),
		},
	}

	for _, test := range testData {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, bk, pfk, _, k := SetupTestInput()

			// Setup
			if !test.userFunds.Empty() {
				_ = bk.SetCoins(ctx, test.owner, test.userFunds)
			}
			if !test.tokenPrice.Equals(pricefeed.EmptyPrice()) {
				pfk.SetCurrentPrice(ctx, test.tokenPrice)
			}

			err := k.OpenCdp(ctx, test.owner, test.amount)
			if test.error != nil {
				require.Equal(t, test.error.Error(), err.Error())
			} else {
				require.NoError(t, err)
			}

			if !test.returnedCredits.IsEqual(sdk.Coins{}) {
				actual := bk.GetCoins(ctx, test.owner)
				require.Equal(t, test.returnedCredits, actual)
			}
		})
	}

}

func TestKeeper_GetCdpsByOwner(t *testing.T) {
	t.Run("Empty list is returned properly", func(t *testing.T) {
		ctx, _, _, _, k := SetupTestInput()
		require.Empty(t, k.CdpsByOwner(ctx, testCdpOwner))
	})

	t.Run("Existing list is returned properly", func(t *testing.T) {
		ctx, _, _, _, k := SetupTestInput()

		k.AddCdp(ctx, testCdp)

		store := ctx.KVStore(k.storeKey)
		var cdps types.Cdps
		k.cdc.MustUnmarshalBinaryBare(store.Get(k.getCdpKey(testCdp.Owner)), &cdps)
		require.Equal(t, types.Cdps{testCdp}, k.CdpsByOwner(ctx, testCdpOwner))
	})
}

func TestKeeper_GetCdpByOwnerAndTimeStamp(t *testing.T) {
	t.Run("Existing cdp with given timestamp returned properly", func(t *testing.T) {
		ctx, _, _, _, k := SetupTestInput()
		k.AddCdp(ctx, testCdp)

		store := ctx.KVStore(k.storeKey)
		var cdps types.Cdps
		k.cdc.MustUnmarshalBinaryBare(store.Get(k.getCdpKey(testCdp.Owner)), &cdps)
		actual, _ := k.GetCdp(ctx, testCdpOwner, 10)
		require.Equal(t, testCdp, actual)
	})
	t.Run("not existent cdp with given timestamp return empty cdp and false", func(t *testing.T) {
		ctx, _, _, _, k := SetupTestInput()

		store := ctx.KVStore(k.storeKey)
		var cdps types.Cdps
		k.cdc.MustUnmarshalBinaryBare(store.Get(k.getCdpKey(testCdp.Owner)), &cdps)
		actual, isFalse := k.GetCdp(ctx, testCdpOwner, 10)
		require.Equal(t, types.Cdp{}, actual)
		require.False(t, isFalse)
	})
}

func TestKeeper_CloseCdp(t *testing.T) {

	t.Run("Non existing CDP returns error", func(t *testing.T) {
		ctx, _, _, _, k := SetupTestInput()

		err := k.CloseCdp(ctx, testCdp.Owner, testCdp.CreatedAt)
		errMsg := fmt.Sprintf("CDP for user with address %s and timestamp %d does not exist", testCdpOwner, testCdp.CreatedAt)
		require.Equal(t, sdkErr.Wrap(sdkErr.ErrUnknownRequest, errMsg).Error(), err.Error())
	})

	t.Run("Existing CDP is closed properly", func(t *testing.T) {
		ctx, bk, _, _, k := SetupTestInput()

		k.AddCdp(ctx, testCdp)
		_ = k.supplyKeeper.MintCoins(ctx, types.ModuleName, testLiquidityPool)
		_, _ = bk.AddCoins(ctx, testCdpOwner, testCdp.Credits)

		require.NoError(t, k.CloseCdp(ctx, testCdpOwner, testCdp.CreatedAt))
		require.Equal(t, testCdp.Deposit, bk.GetCoins(ctx, testCdpOwner))
	})

}

func TestKeeper_DeleteCdp(t *testing.T) {
	testData := []struct {
		name            string
		existingCdps    types.Cdps
		deletedCdp      types.Cdp
		shouldBeDeleted bool
	}{
		{
			name:            "Existing CDP is deleted",
			existingCdps:    types.Cdps{testCdp},
			deletedCdp:      testCdp,
			shouldBeDeleted: true,
		},
		{
			name:         "Non existent CDP is not deleted",
			existingCdps: types.Cdps{testCdp},
			deletedCdp: types.Cdp{
				Owner:     testCdp.Owner,
				Deposit:   testCdp.Deposit,
				Credits:   testCdp.Credits,
				CreatedAt: testCdp.CreatedAt + 1,
			},
			shouldBeDeleted: false,
		},
	}

	for _, test := range testData {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, _, _, _, k := SetupTestInput()

			for _, cdp := range test.existingCdps {
				k.AddCdp(ctx, cdp)
			}

			k.deleteCdp(ctx, test.deletedCdp)

			result := k.GetCdps(ctx)
			if test.shouldBeDeleted {
				require.Len(t, result, len(test.existingCdps)-1)
			} else {
				require.Len(t, result, len(test.existingCdps))
			}
		})
	}
}

func TestKeeper_AutoLiquidateCdp(t *testing.T) {
	ctx, bk, pfk, _, k := SetupTestInput()
	// Setup
	if !testCdp.Deposit.IsZero() {
		_ = bk.SetCoins(ctx, testCdpOwner, sdk.NewCoins(testCdp.Deposit))
	}
	tokenPrice := pricefeed.NewPrice(testLiquidityDenom, sdk.NewDec(10), sdk.NewInt(1000))
	if !tokenPrice.Equals(pricefeed.EmptyPrice()) {
		pfk.SetCurrentPrice(ctx, tokenPrice)
	}
	require.NoError(t, k.OpenCdp(ctx, testCdp.Owner, testCdp.Deposit))
	cdps := k.CdpsByOwner(ctx, testCdp.Owner)
	require.Equal(t, 1, len(cdps))
	yes, err := k.ShouldLiquidateCdp(ctx, cdps[0])
	require.NoError(t, err)
	require.True(t, yes)
	require.NotPanics(t, func() { k.AutoLiquidateCdps(ctx) })
	require.Equal(t, 0, len(k.CdpsByOwner(ctx, testCdp.Owner)))
}

// --------------
// --- CdpCollateralRate
// --------------

func TestKeeper_SetCdpCollateralRate(t *testing.T) {
	ctx, _, _, _, k := SetupTestInput()
	require.Error(t, k.SetCollateralRate(ctx, sdk.NewInt(0).ToDec()))
	require.Error(t, k.SetCollateralRate(ctx, sdk.NewInt(-1).ToDec()))
	require.NoError(t, k.SetCollateralRate(ctx, sdk.NewInt(2).ToDec()))
	rate := sdk.NewDec(3)
	require.NoError(t, k.SetCollateralRate(ctx, rate))

	var got sdk.Dec
	k.cdc.MustUnmarshalBinaryBare(ctx.KVStore(k.storeKey).Get([]byte(types.CollateralRateKey)), &got)
	require.True(t, rate.Equal(got), got.String())
}

func TestKeeper_GetCdpCollateralRate(t *testing.T) {
	ctx, _, _, _, k := SetupTestInput()
	rate := sdk.NewDec(3)
	require.NoError(t, k.SetCollateralRate(ctx, rate))
	require.Equal(t, rate, k.GetCollateralRate(ctx))
}
