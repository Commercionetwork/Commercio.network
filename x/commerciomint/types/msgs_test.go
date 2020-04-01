package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

//
//import (
//	"testing"
//	"time"
//
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/stretchr/testify/assert"
//)
//
//var testMsgOpenCdp = MsgOpenCdp{}
//var testMsgCloseCdp = MsgCloseCdp{
//	Signer:    testOwner,
//}
//
//func TestMsgOpenCdp_Route(t *testing.T) {
//	actual := testMsgOpenCdp.Route()
//	require.Equal(t, RouterKey, actual)
//}
//
//func TestMsgOpenCdp_Type(t *testing.T) {
//	actual := testMsgOpenCdp.Type()
//	require.Equal(t, MsgTypeOpenCdp, actual)
//}
//
//func TestMsgOpenCdp_ValidateBasic_Valid(t *testing.T) {
//	actual := testMsgOpenCdp.ValidateBasic()
//	require.Nil(t, actual)
//}
//
//func TestMsgOpenCdp_ValidateBasic_InvalidOwnerAddr(t *testing.T) {
//	invalidMsg := MsgOpenCdp(CdpRequest{
//		Signer:          nil,
//		DepositedAmount: nil,
//		Timestamp:       time.Time{},
//	})
//	err := invalidMsg.ValidateBasic()
//	require.Error(t, err)
//	require.Equal(t, sdkErr.Wrap(sdkErr.ErrInvalidAddress, invalidMsg.Depositor.String()), err)
//}
//
//func TestMsgOpenCdp_ValidateBasic_InvalidDepositedAmount(t *testing.T) {
//	invalidMsg := MsgOpenCdp(CdpRequest{
//		Signer:          testOwner,
//		DepositedAmount: nil,
//		Timestamp:       time.Time{},
//	})
//	err := invalidMsg.ValidateBasic()
//	require.Error(t, err)
//	require.Equal(t, sdk.ErrInvalidCoins(invalidMsg.DepositedAmount.String()), err)
//}
//
//func TestMsgOpenCdp_ValidateBasic_InvalidTimestamp(t *testing.T) {
//	invalidMsg := MsgOpenCdp(CdpRequest{
//		Signer:          testOwner,
//		DepositedAmount: TestDepositedAmount,
//		Timestamp:       time.Time{},
//	})
//	err := invalidMsg.ValidateBasic()
//	require.Error(t, err)
//	require.Equal(t, sdkErr.Wrap(sdkErr.ErrUnknownRequest, "cdp request's timestamp is invalid"), err)
//}
//
//func TestMsgOpenCdp_GetSignBytes(t *testing.T) {
//	actual := testMsgOpenCdp.GetSignBytes()
//	expected := sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(testMsgOpenCdp))
//	require.Equal(t, expected, actual)
//}
//
//func TestMsgOpenCdp_GetSigners(t *testing.T) {
//	actual := testMsgOpenCdp.GetSigners()
//	expected := []sdk.AccAddress{testMsgOpenCdp.Depositor}
//	require.Equal(t, expected, actual)
//}
//
/////////////////////
/////MsgCloseCdp////
///////////////////
//func TestMsgCloseCdp_Route(t *testing.T) {
//	actual := testMsgCloseCdp.Route()
//	require.Equal(t, RouterKey, actual)
//}
//
//func TestMsgCloseCdp_Type(t *testing.T) {
//	actual := testMsgCloseCdp.Type()
//	require.Equal(t, MsgTypeCloseCdp, actual)
//}
//
//func TestMsgCloseCdp_ValidateBasic_Valid(t *testing.T) {
//	actual := testMsgCloseCdp.ValidateBasic()
//	require.Nil(t, actual)
//}
//
//func TestMsgCloseCdp_ValidateBasic_InvalidSigner(t *testing.T) {
//	msg := MsgCloseCdp{
//		Signer:    nil,
//		Timestamp: time.Time{},
//	}
//	actual := msg.ValidateBasic()
//	require.Equal(t, sdkErr.Wrap(sdkErr.ErrInvalidAddress, msg.Signer.String()), actual)
//	require.Error(t, actual)
//}
//
//func TestMsgCloseCdp_ValidateBasic_InvalidTimestamp(t *testing.T) {
//	msg := MsgCloseCdp{
//		Signer:    testOwner,
//		Timestamp: time.Time{},
//	}
//	actual := msg.ValidateBasic()
//	require.Equal(t, sdkErr.Wrap(sdkErr.ErrUnknownRequest, "cdp's timestamp is invalid"), actual)
//	require.Error(t, actual)
//}
//
//func TestMsgCloseCdp_GetSignBytes(t *testing.T) {
//	actual := testMsgCloseCdp.GetSignBytes()
//	expected := sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(testMsgCloseCdp))
//	require.Equal(t, expected, actual)
//}
//
//func TestMsgCloseCdp_GetSigners(t *testing.T) {
//	actual := testMsgCloseCdp.GetSigners()
//	expected := []sdk.AccAddress{testMsgCloseCdp.Signer}
//	require.Equal(t, expected, actual)
//}

func TestMsgSetCdpCollateralRate_ValidateBasic(t *testing.T) {
	type fields struct {
		Signer            sdk.AccAddress
		CdpCollateralRate sdk.Dec
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"empty signer", fields{nil, sdk.NewDec(2)}, true},
		{"ok", fields{[]byte("test"), sdk.NewDec(2)}, false},
		{"zero collateral rate", fields{[]byte("test"), sdk.NewDec(0)}, true},
		{"negative collateral rate", fields{[]byte("test"), sdk.NewDec(-1)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := NewMsgSetCdpCollateralRate(tt.fields.Signer, tt.fields.CdpCollateralRate)
			require.Equal(t, "commerciomint", msg.Route())
			require.Equal(t, "setCdpCollateralRate", msg.Type())
			require.Equal(t, tt.wantErr, msg.ValidateBasic() != nil)
		})
	}
}
