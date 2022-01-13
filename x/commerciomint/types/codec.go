package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgMintCCC{}, "commerciomint/MintCCC", nil)
	cdc.RegisterConcrete(&MsgBurnCCC{}, "commerciomint/BurnCCC", nil)
	cdc.RegisterConcrete(&MsgSetCCCConversionRate{}, "commerciomint/MsgSetCCCConversionRate", nil)
	cdc.RegisterConcrete(&MsgSetCCCFreezePeriod{}, "commerciomint/MsgSetCCCFreezePeriod", nil)
	cdc.RegisterConcrete(&MsgSetParams{}, "commerciomint/MsgSetParams", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintCCC{},
		&MsgBurnCCC{},
		&MsgSetCCCConversionRate{},
		&MsgSetCCCFreezePeriod{},
		&MsgSetParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	//amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
