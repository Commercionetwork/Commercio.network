package types

import sdk "github.com/cosmos/cosmos-sdk/types"

var sender, _ = sdk.AccAddressFromBech32("cosmos1lwmppctrr6ssnrmuyzu554dzf50apkfvd53jx0")
var recipient, _ = sdk.AccAddressFromBech32("cosmos1v0yk4hs2nry020ufmu9yhpm39s4scdhhtecvtr")

var ValidDocument = Document{
	UUID:       "d83422c6-6e79-4a99-9767-fcae46dfa371",
	ContentURI: "https://example.com/document",
	Metadata: &DocumentMetadata{
		ContentURI: "https://example.com/document/metadata",
		Schema: &DocumentMetadataSchema{
			URI:     "https://example.com/document/metadata/schema",
			Version: "1.0.0",
		},
	},
	Checksum: &DocumentChecksum{
		Value:     "93dfcaf3d923ec47edb8580667473987",
		Algorithm: "md5",
	},
	EncryptionData: &DocumentEncryptionData{
		Keys:          []*DocumentEncryptionKey{{Recipient: recipient.String(), Value: "6F7468657276616C7565"}},
		EncryptedData: []string{"content", "content_uri", "metadata.content_uri", "metadata.schema.uri"},
	},
	DoSign: &DocumentDoSign{
		StorageURI:     "https://example.com/document/storage",
		SignerInstance: "SignerInstance",
		SdnData: SdnData{
			SdnDataCommonName,
			SdnDataSurname,
			SdnDataSurname,
			SdnDataGivenName,
			SdnDataOrganization,
			SdnDataCountry,
		},
		VcrID:              "VcrID",
		CertificateProfile: "CertificateProfile",
	},
	Sender:     sender.String(),
	Recipients: []string{recipient.String()},
}

var ValidDocumentReceipt = DocumentReceipt{
	UUID:         "8db853ac-5265-4da6-a07a-c52ac8099385",
	Sender:       sender.String(),
	Recipient:    recipient.String(),
	TxHash:       "txHash",
	DocumentUUID: ValidDocument.UUID,
	Proof:        "proof",
}