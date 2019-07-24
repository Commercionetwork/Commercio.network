package types

/**
 * Any interface you create and any struct that implements an interface needs to be
 * declared in the RegisterCodec function.
 * In this module the Msg implementations (SetIdentity) need to be registered,
 * but your Identity query return type does not.
 */

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on wire codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetIdentity{}, "commercioid/SetIdentity", nil)
	cdc.RegisterConcrete(MsgCreateConnection{}, "commectioid/CreateConnection", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
