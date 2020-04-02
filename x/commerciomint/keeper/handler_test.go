package keeper

import (
	"fmt"
	"testing"

	sdkErr "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/commercionetwork/commercionetwork/x/commerciomint/types"
	"github.com/commercionetwork/commercionetwork/x/pricefeed"
)

var testMsgOpenCdp = types.NewMsgOpenCdp(testCdp.Owner, testCdp.DepositedAmount)
var testMsgCloseCdp = types.NewMsgCloseCdp(testCdp.Owner, testCdp.Timestamp)

func TestHandler_handleMsgOpenCdp(t *testing.T) {
	ctx, bk, pfk, _, k := SetupTestInput()
	handler := NewHandler(k)

	// Test setup
	_, _ = bk.AddCoins(ctx, testCdp.Owner, testCdp.DepositedAmount)
	balance := bk.GetCoins(ctx, testCdpOwner)

	// Check balance
	require.Equal(t, "100ucommercio", balance.String())

	// Set credits denom and push a price to pricefeed
	k.SetCreditsDenom(ctx, "uccc")
	pfk.SetCurrentPrice(ctx, pricefeed.NewPrice("ucommercio", sdk.NewDec(10), sdk.NewInt(1000)))

	actual, err := handler(ctx, testMsgOpenCdp)
	require.NoError(t, err)
	require.Equal(t, &sdk.Result{Log: "Cdp opened successfully"}, actual)

	// Check final balance
	balance = bk.GetCoins(ctx, testCdpOwner)
	require.Equal(t, "500uccc", balance.String())
}

func TestHandler_handleMsgCloseCdp(t *testing.T) {
	ctx, bk, _, _, k := SetupTestInput()
	handler := NewHandler(k)

	_, _ = bk.AddCoins(ctx, k.supplyKeeper.GetModuleAddress(types.ModuleName), testCdp.DepositedAmount)
	_ = bk.SetCoins(ctx, testCdp.Owner, testCdp.CreditsAmount)
	k.AddCdp(ctx, testCdp)

	expected := &sdk.Result{Log: "Cdp closed successfully"}
	actual, err := handler(ctx, testMsgCloseCdp)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestHandler_InvalidMsg(t *testing.T) {
	ctx, _, _, _, k := SetupTestInput()
	handler := NewHandler(k)

	invalidMsg := sdk.NewTestMsg()
	errMsg := fmt.Sprintf("Unrecognized %s message type: %v", types.ModuleName, invalidMsg.Type())
	expected := sdkErr.Wrap(sdkErr.ErrUnknownRequest, errMsg)

	_, err := handler(ctx, invalidMsg)
	require.Error(t, err)
	require.Equal(t, expected.Error(), err.Error())
}
