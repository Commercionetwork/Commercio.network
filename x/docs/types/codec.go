package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on wire codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgShareDocument{}, "commercio/MsgShareDocument", nil)
	cdc.RegisterConcrete(MsgSendDocumentReceipt{}, "commercio/MsgSendDocumentReceipt", nil)
	cdc.RegisterConcrete(MsgAddSupportedMetadataSchema{}, "commercio/MsgAddSupportedMetadataSchema", nil)
	cdc.RegisterConcrete(MsgAddTrustedMetadataSchemaProposer{}, "commercio/MsgAddTrustedMetadataSchemaProposer", nil)
}

var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
