package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestCurrentPrice_Equals_true(t *testing.T) {
	actual := TestPriceInfo.Equals(TestPriceInfo)
	assert.True(t, actual)
}

func TestCurrentPrice_Equals_false(t *testing.T) {
	curPrice := CurrentPrice{
		AssetName: TestPriceInfo.AssetName,
		Price:     sdk.NewDec(2),
		Expiry:    sdk.NewInt(5),
	}

	actual := TestPriceInfo.Equals(curPrice)
	assert.False(t, actual)
}

func TestCurrentPrices_GetPrice(t *testing.T) {
	prices := CurrentPrices{TestPriceInfo}
	actual, _ := prices.GetPrice(TestPriceInfo.AssetName)
	assert.Equal(t, TestPriceInfo, actual)
}

func TestCurrentPrices_GetPrice_NotFound(t *testing.T) {
	prices := CurrentPrices{}
	_, err := prices.GetPrice("testName")
	assert.Error(t, err)
}

func TestCurrentPrices_AppendIfMissing_Notfound(t *testing.T) {
	prices := CurrentPrices{}
	actual, found := prices.AppendIfMissing(TestPriceInfo)
	expected := CurrentPrices{TestPriceInfo}
	assert.Equal(t, expected, actual)
	assert.False(t, found)
}

func TestCurrentPrices_AppendIfMissing_found(t *testing.T) {
	prices := CurrentPrices{TestPriceInfo}
	actual, found := prices.AppendIfMissing(TestPriceInfo)
	assert.Nil(t, actual)
	assert.True(t, found)
}

func TestRawPrice_Equals_false(t *testing.T) {
	curPrice := CurrentPrice{
		AssetName: TestPriceInfo.AssetName,
		Price:     sdk.NewDec(2),
		Expiry:    sdk.NewInt(5),
	}
	rawPrice := RawPrice{
		Oracle:    TestOracle,
		PriceInfo: curPrice,
	}

	actual := rawPrice.Equals(TestRawPrice)
	assert.False(t, actual)
}

func TestRawPrice_Equals_true(t *testing.T) {
	actual := TestRawPrice.Equals(TestRawPrice)
	assert.True(t, actual)
}

func TestRawPrices_UpdatePriceOrAppendIfMissing_appendNewRawPrice(t *testing.T) {
	rawPrices := RawPrices{}
	actual, updated := rawPrices.UpdatePriceOrAppendIfMissing(TestRawPrice)
	assert.True(t, updated)
	assert.Equal(t, TestRawPrice, actual[0])
}

func TestRawPrices_UpdatePriceOrAppendIfMissing_priceAlreadyInserted(t *testing.T) {
	rawPrices := RawPrices{TestRawPrice}
	actual, updated := rawPrices.UpdatePriceOrAppendIfMissing(TestRawPrice)
	assert.False(t, updated)
	assert.Equal(t, rawPrices, actual)
}

func TestRawPrices_UpdatePriceOrAppendIfMissing_updatedPrice(t *testing.T) {
	curPrice := CurrentPrice{
		AssetName: TestPriceInfo.AssetName,
		Price:     sdk.NewDec(200),
		Expiry:    sdk.NewInt(6000),
	}
	rawPrice := RawPrice{
		Oracle:    TestRawPrice.Oracle,
		PriceInfo: curPrice,
	}
	rawPrices := RawPrices{TestRawPrice}
	actual, updated := rawPrices.UpdatePriceOrAppendIfMissing(rawPrice)
	assert.Equal(t, rawPrice, actual[0])
	assert.True(t, updated)
}
