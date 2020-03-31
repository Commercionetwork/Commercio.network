package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestCurrentPrice_Equals_true(t *testing.T) {
	actual := testPrice.Equals(testPrice)
	require.True(t, actual)
}

func TestCurrentPrice_Equals_false(t *testing.T) {
	curPrice := Price{
		AssetName: testPrice.AssetName,
		Value:     sdk.NewDec(2),
		Expiry:    sdk.NewInt(5),
	}

	actual := testPrice.Equals(curPrice)
	require.False(t, actual)
}

func TestCurrentPrices_AppendIfMissing_Notfound(t *testing.T) {
	prices := Prices{}
	actual, found := prices.AppendIfMissing(testPrice)
	expected := Prices{testPrice}
	require.Equal(t, expected, actual)
	require.False(t, found)
}

func TestCurrentPrices_AppendIfMissing_found(t *testing.T) {
	prices := Prices{testPrice}
	actual, found := prices.AppendIfMissing(testPrice)
	require.Nil(t, actual)
	require.True(t, found)
}

func TestPrice_Equals_true(t *testing.T) {
	actual := testPrice.Equals(testPrice)
	require.True(t, actual)
}

// -----------------
// --- OraclePrice
// -----------------

var testRawPrice = OraclePrice{
	Oracle:  testOracle,
	Price:   Price{AssetName: "uatom", Value: sdk.NewDecWithPrec(15423, 2), Expiry: sdk.NewInt(1100)},
	Created: sdk.NewInt(0),
}

func TestRawPrices_UpdatePriceOrAppendIfMissing(t *testing.T) {
	testData := []struct {
		name            string
		prices          OraclePrices
		price           OraclePrice
		shouldBeUpdated bool
	}{
		{
			name:            "New price inserted correctly",
			prices:          OraclePrices{},
			price:           testRawPrice,
			shouldBeUpdated: true,
		},
		{
			name:            "Price already inserted is not appended",
			prices:          OraclePrices{testRawPrice},
			price:           testRawPrice,
			shouldBeUpdated: false,
		},
		{
			name:   "Different expiration date price is replaced",
			prices: OraclePrices{testRawPrice},
			price: OraclePrice{
				Oracle: testRawPrice.Oracle,
				Price: Price{
					AssetName: testRawPrice.Price.AssetName,
					Value:     testRawPrice.Price.Value,
					Expiry:    testRawPrice.Price.Expiry.Add(sdk.NewInt(10)),
				},
				Created: testRawPrice.Created,
			},
			shouldBeUpdated: true,
		},
		{
			name:   "Different creation date price is replaced",
			prices: OraclePrices{testRawPrice},
			price: OraclePrice{
				Oracle:  testRawPrice.Oracle,
				Created: testRawPrice.Created.Add(sdk.NewInt(10)),
				Price:   testRawPrice.Price,
			},
			shouldBeUpdated: true,
		},
	}

	for _, test := range testData {
		test := test
		t.Run(test.name, func(t *testing.T) {
			prices, updated := test.prices.UpdatePriceOrAppendIfMissing(test.price)
			require.Equal(t, test.shouldBeUpdated, updated)
			if test.shouldBeUpdated {
				require.Contains(t, prices, test.price)
			}
		})
	}
}
