package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// -------------------
// --- MsgSetPrice
// -------------------

type MsgSetPrice struct {
	Oracle sdk.AccAddress `json:"oracle"`
	Price  Price          `json:"price"`
}

func NewMsgSetPrice(price Price, oracle sdk.AccAddress) MsgSetPrice {
	return MsgSetPrice{
		Oracle: oracle,
		Price:  price,
	}
}

// Route Implements Msg.
func (msg MsgSetPrice) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgSetPrice) Type() string { return MsgTypeSetPrice }

// ValidateBasic Implements Msg.
func (msg MsgSetPrice) ValidateBasic() sdk.Error {
	if msg.Oracle.Empty() {
		return sdk.ErrInvalidAddress(msg.Oracle.String())
	}
	if msg.Price.Value.IsNegative() {
		return sdk.ErrUnknownRequest("Token's price cannot be zero or negative")
	}
	if len(strings.TrimSpace(msg.Price.AssetName)) == 0 {
		return sdk.ErrUnknownRequest("Cannot set price for unnamed token")
	}
	if msg.Price.Expiry.IsZero() || msg.Price.Expiry.IsNegative() {
		return sdk.ErrUnknownRequest("Cannot set price with an expire height of zero or negative")
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSetPrice) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgSetPrice) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Oracle}
}

// -------------------
// --- MsgAddOracle
// -------------------

type MsgAddOracle struct {
	Signer sdk.AccAddress `json:"signer"`
	Oracle sdk.AccAddress `json:"oracle"`
}

func NewMsgAddOracle(signer sdk.AccAddress, oracle sdk.AccAddress) MsgAddOracle {
	return MsgAddOracle{
		Signer: signer,
		Oracle: oracle,
	}
}

// Route Implements Msg.
func (msg MsgAddOracle) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgAddOracle) Type() string { return MsgTypeAddOracle }

// ValidateBasic Implements Msg.
func (msg MsgAddOracle) ValidateBasic() sdk.Error {
	if msg.Signer.Empty() {
		return sdk.ErrInvalidAddress(msg.Signer.String())
	}
	if msg.Oracle.Empty() {
		return sdk.ErrInvalidAddress(msg.Oracle.String())
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgAddOracle) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgAddOracle) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}
