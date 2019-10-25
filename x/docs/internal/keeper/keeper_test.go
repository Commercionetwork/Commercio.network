package keeper

import (
	"testing"

	ctypes "github.com/commercionetwork/commercionetwork/x/common/types"
	"github.com/commercionetwork/commercionetwork/x/docs/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------
// --- Metadata schemes
// ----------------------------------

func TestKeeper_AddSupportedMetadataScheme_EmptyList(t *testing.T) {
	cdc, ctx, k := SetupTestInput()
	//Setup the store
	store := ctx.KVStore(k.StoreKey)

	schema := types.MetadataSchema{Type: "schema", SchemaUri: "https://example.com/schema", Version: "1.0.0"}
	k.AddSupportedMetadataScheme(ctx, schema)

	var stored types.MetadataSchemes
	storedBz := store.Get([]byte(types.SupportedMetadataSchemesStoreKey))
	cdc.MustUnmarshalBinaryBare(storedBz, &stored)

	assert.Equal(t, 1, len(stored))
	assert.Contains(t, stored, schema)
}

func TestKeeper_AddSupportedMetadataScheme_ExistingList(t *testing.T) {
	cdc, ctx, k := SetupTestInput()
	//Setup the store
	store := ctx.KVStore(k.StoreKey)

	existingSchema := types.MetadataSchema{Type: "schema", SchemaUri: "https://example.com/newSchema", Version: "1.0.0"}
	existing := []types.MetadataSchema{existingSchema}
	existingBz := cdc.MustMarshalBinaryBare(&existing)
	store.Set([]byte(types.SupportedMetadataSchemesStoreKey), existingBz)

	newSchema := types.MetadataSchema{Type: "schema2", SchemaUri: "https://example.com/schema2", Version: "2.0.0"}
	k.AddSupportedMetadataScheme(ctx, newSchema)

	var stored types.MetadataSchemes
	storedBz := store.Get([]byte(types.SupportedMetadataSchemesStoreKey))
	cdc.MustUnmarshalBinaryBare(storedBz, &stored)

	assert.Equal(t, 2, len(stored))
	assert.Contains(t, stored, existingSchema)
	assert.Contains(t, stored, newSchema)
}

func TestKeeper_IsMetadataSchemeTypeSupported_EmptyList(t *testing.T) {
	_, ctx, k := SetupTestInput()

	assert.False(t, k.IsMetadataSchemeTypeSupported(ctx, "schema"))
	assert.False(t, k.IsMetadataSchemeTypeSupported(ctx, "schema2"))
	assert.False(t, k.IsMetadataSchemeTypeSupported(ctx, "non-existent"))
}

func TestKeeper_IsMetadataSchemeTypeSupported_ExistingList(t *testing.T) {
	cdc, ctx, k := SetupTestInput()
	store := ctx.KVStore(k.StoreKey)

	existingSchema := types.MetadataSchema{Type: "schema", SchemaUri: "https://example.com/newSchema", Version: "1.0.0"}
	existing := []types.MetadataSchema{existingSchema}
	existingBz := cdc.MustMarshalBinaryBare(&existing)
	store.Set([]byte(types.SupportedMetadataSchemesStoreKey), existingBz)

	assert.True(t, k.IsMetadataSchemeTypeSupported(ctx, "schema"))
	assert.False(t, k.IsMetadataSchemeTypeSupported(ctx, "schema2"))
	assert.False(t, k.IsMetadataSchemeTypeSupported(ctx, "any-schema"))
}

func TestKeeper_GetSupportedMetadataSchemes_EmptyList(t *testing.T) {
	_, ctx, k := SetupTestInput()

	result := k.GetSupportedMetadataSchemes(ctx)

	assert.Empty(t, result)
}

func TestKeeper_GetSupportedMetadataSchemes_ExistingList(t *testing.T) {
	cdc, ctx, k := SetupTestInput()
	store := ctx.KVStore(k.StoreKey)

	existingSchema := types.MetadataSchema{Type: "schema", SchemaUri: "https://example.com/newSchema", Version: "1.0.0"}
	existing := types.MetadataSchemes{existingSchema}
	existingBz := cdc.MustMarshalBinaryBare(&existing)
	store.Set([]byte(types.SupportedMetadataSchemesStoreKey), existingBz)

	actual := k.GetSupportedMetadataSchemes(ctx)

	assert.Equal(t, existing, actual)
}

// ----------------------------------
// --- Metadata schema proposers
// ----------------------------------

func TestKeeper_AddTrustedSchemaProposer_EmptyList(t *testing.T) {
	cdc, ctx, k := SetupTestInput()
	store := ctx.KVStore(k.StoreKey)

	k.AddTrustedSchemaProposer(ctx, TestingSender)

	var proposers []sdk.AccAddress
	proposersBz := store.Get([]byte(types.MetadataSchemaProposersStoreKey))
	cdc.MustUnmarshalBinaryBare(proposersBz, &proposers)

	assert.Equal(t, 1, len(proposers))
	assert.Contains(t, proposers, TestingSender)
}

func TestKeeper_AddTrustedSchemaProposer_ExistingList(t *testing.T) {
	cdc, ctx, k := SetupTestInput()
	store := ctx.KVStore(k.StoreKey)

	existing := []sdk.AccAddress{TestingSender}
	proposersBz := cdc.MustMarshalBinaryBare(&existing)
	store.Set([]byte(types.MetadataSchemaProposersStoreKey), proposersBz)

	k.AddTrustedSchemaProposer(ctx, TestingSender2)

	var stored []sdk.AccAddress
	storedBz := store.Get([]byte(types.MetadataSchemaProposersStoreKey))
	cdc.MustUnmarshalBinaryBare(storedBz, &stored)

	assert.Equal(t, 2, len(stored))
	assert.Contains(t, stored, TestingSender)
	assert.Contains(t, stored, TestingSender2)
}

func TestKeeper_IsTrustedSchemaProposer_EmptyList(t *testing.T) {
	_, ctx, k := SetupTestInput()

	assert.False(t, k.IsTrustedSchemaProposer(ctx, TestingSender))
	assert.False(t, k.IsTrustedSchemaProposer(ctx, TestingSender2))
}

func TestKeeper_IsTrustedSchemaProposerExistingList(t *testing.T) {
	cdc, ctx, k := SetupTestInput()
	store := ctx.KVStore(k.StoreKey)

	existing := []sdk.AccAddress{TestingSender}
	proposersBz := cdc.MustMarshalBinaryBare(&existing)
	store.Set([]byte(types.MetadataSchemaProposersStoreKey), proposersBz)

	assert.True(t, k.IsTrustedSchemaProposer(ctx, TestingSender))
	assert.False(t, k.IsTrustedSchemaProposer(ctx, TestingSender2))
}

func TestKeeper_GetTrustedSchemaProposers_EmptyList(t *testing.T) {
	_, ctx, k := SetupTestInput()

	stored := k.GetTrustedSchemaProposers(ctx)

	assert.Empty(t, stored)
}

func TestKeeper_GetTrustedSchemaProposers_ExistingList(t *testing.T) {
	cdc, ctx, k := SetupTestInput()
	store := ctx.KVStore(k.StoreKey)

	existing := ctypes.Addresses{TestingSender}
	proposersBz := cdc.MustMarshalBinaryBare(&existing)
	store.Set([]byte(types.MetadataSchemaProposersStoreKey), proposersBz)

	stored := k.GetTrustedSchemaProposers(ctx)

	assert.Equal(t, existing, stored)
}

// ----------------------------------
// --- Documents
// ----------------------------------

func TestKeeper_ShareDocument_EmptyList(t *testing.T) {
	cdc, ctx, k := SetupTestInput()
	store := ctx.KVStore(k.StoreKey)

	err := k.SaveDocument(ctx, TestingDocument)
	assert.Nil(t, err)

	docsBz := store.Get(k.getDocumentStoreKey(TestingDocument.Uuid))
	sentDocsBz := store.Get(k.getSentDocumentsIdsStoreKey(TestingSender))
	receivedDocsBz := store.Get(k.getReceivedDocumentsIdsStoreKey(TestingRecipient))

	var stored types.Document
	cdc.MustUnmarshalBinaryBare(docsBz, &stored)
	assert.Equal(t, stored, TestingDocument)

	var sentDocs, receivedDocs types.DocumentIds
	cdc.MustUnmarshalBinaryBare(sentDocsBz, &sentDocs)
	cdc.MustUnmarshalBinaryBare(receivedDocsBz, &receivedDocs)

	assert.Equal(t, 1, len(sentDocs))
	assert.Contains(t, sentDocs, TestingDocument.Uuid)

	assert.Equal(t, 1, len(receivedDocs))
	assert.Contains(t, receivedDocs, TestingDocument.Uuid)
}

func TestKeeper_ShareDocument_ExistingDocument(t *testing.T) {
	cdc, ctx, k := SetupTestInput()
	store := ctx.KVStore(k.StoreKey)

	store.Set(k.getDocumentStoreKey(TestingDocument.Uuid), cdc.MustMarshalBinaryBare(TestingDocument))

	documentsIds := types.DocumentIds{TestingDocument.Uuid}
	store.Set(k.getSentDocumentsIdsStoreKey(TestingSender), cdc.MustMarshalBinaryBare(&documentsIds))
	store.Set(k.getReceivedDocumentsIdsStoreKey(TestingRecipient), cdc.MustMarshalBinaryBare(&documentsIds))

	err := k.SaveDocument(ctx, TestingDocument)
	assert.NotNil(t, err)

	sentDocsBz := store.Get(k.getSentDocumentsIdsStoreKey(TestingSender))
	receivedDocsBz := store.Get(k.getReceivedDocumentsIdsStoreKey(TestingRecipient))

	var sentDocs, receivedDocs types.DocumentIds
	cdc.MustUnmarshalBinaryBare(sentDocsBz, &sentDocs)
	cdc.MustUnmarshalBinaryBare(receivedDocsBz, &receivedDocs)

	assert.Equal(t, 1, len(sentDocs))
	assert.Contains(t, sentDocs, TestingDocument.Uuid)

	assert.Equal(t, 1, len(receivedDocs))
	assert.Contains(t, receivedDocs, TestingDocument.Uuid)
}

func TestKeeper_ShareDocument_ExistingDocument_DifferentRecipient(t *testing.T) {
	cdc, ctx, k := SetupTestInput()
	documentsIds := types.DocumentIds{TestingDocument.Uuid}

	store := ctx.KVStore(k.StoreKey)
	store.Set(k.getSentDocumentsIdsStoreKey(TestingSender), cdc.MustMarshalBinaryBare(&documentsIds))
	store.Set(k.getReceivedDocumentsIdsStoreKey(TestingRecipient), cdc.MustMarshalBinaryBare(&documentsIds))

	newRecipient, _ := sdk.AccAddressFromBech32("cosmos1h2z8u9294gtqmxlrnlyfueqysng3krh009fum7")
	newDocument := types.Document{
		Uuid:       TestingDocument.Uuid,
		ContentUri: TestingDocument.ContentUri,
		Metadata:   TestingDocument.Metadata,
		Checksum:   TestingDocument.Checksum,
		Sender:     TestingDocument.Sender,
		Recipients: ctypes.Addresses{newRecipient},
	}
	err := k.SaveDocument(ctx, newDocument)
	assert.Nil(t, err)

	sentDocsBz := store.Get(k.getSentDocumentsIdsStoreKey(TestingSender))
	receivedDocsBz := store.Get(k.getReceivedDocumentsIdsStoreKey(TestingRecipient))
	newReceivedDocsBz := store.Get(k.getReceivedDocumentsIdsStoreKey(newRecipient))

	var sentDocs, receivedDocs, newReceivedDocs types.DocumentIds
	cdc.MustUnmarshalBinaryBare(sentDocsBz, &sentDocs)
	cdc.MustUnmarshalBinaryBare(receivedDocsBz, &receivedDocs)
	cdc.MustUnmarshalBinaryBare(newReceivedDocsBz, &newReceivedDocs)

	assert.Equal(t, 1, len(sentDocs))
	assert.Contains(t, sentDocs, TestingDocument.Uuid)

	assert.Equal(t, 1, len(receivedDocs))
	assert.Contains(t, receivedDocs, TestingDocument.Uuid)

	assert.Equal(t, 1, len(newReceivedDocs))
	assert.Contains(t, newReceivedDocs, newDocument.Uuid)
}

func TestKeeper_ShareDocument_ExistingDocument_DifferentUuid(t *testing.T) {
	cdc, ctx, k := SetupTestInput()
	documentsIds := types.DocumentIds{TestingDocument.Uuid}

	store := ctx.KVStore(k.StoreKey)
	store.Set(k.getSentDocumentsIdsStoreKey(TestingSender), cdc.MustMarshalBinaryBare(&documentsIds))
	store.Set(k.getReceivedDocumentsIdsStoreKey(TestingRecipient), cdc.MustMarshalBinaryBare(&documentsIds))

	newDocument := types.Document{
		Uuid:       TestingDocument.Uuid + "new",
		ContentUri: TestingDocument.ContentUri,
		Metadata:   TestingDocument.Metadata,
		Checksum:   TestingDocument.Checksum,
		Recipients: TestingDocument.Recipients,
		Sender:     TestingDocument.Sender,
	}
	err := k.SaveDocument(ctx, newDocument)
	assert.Nil(t, err)

	sentDocsBz := store.Get(k.getSentDocumentsIdsStoreKey(TestingSender))
	receivedDocsBz := store.Get(k.getReceivedDocumentsIdsStoreKey(TestingRecipient))

	var sentDocs, receivedDocs types.DocumentIds
	cdc.MustUnmarshalBinaryBare(sentDocsBz, &sentDocs)
	cdc.MustUnmarshalBinaryBare(receivedDocsBz, &receivedDocs)

	assert.Equal(t, 2, len(sentDocs))
	assert.Contains(t, sentDocs, TestingDocument.Uuid)

	assert.Equal(t, 2, len(receivedDocs))
	assert.Contains(t, receivedDocs, TestingDocument.Uuid)
}

func TestKeeper_GetDocumentById_NonExisting(t *testing.T) {
	_, ctx, k := SetupTestInput()
	_, found := k.GetDocumentById(ctx, "non-existing")
	assert.False(t, found)
}

func TestKeeper_GetDocumentById_Existing(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	store := ctx.KVStore(k.StoreKey)
	store.Set(k.getDocumentStoreKey(TestingDocument.Uuid), cdc.MustMarshalBinaryBare(&TestingDocument))

	doc, found := k.GetDocumentById(ctx, TestingDocument.Uuid)
	assert.True(t, found)
	assert.Equal(t, TestingDocument, doc)
}

func TestKeeper_GetUserReceivedDocuments_EmptyList(t *testing.T) {
	_, ctx, k := SetupTestInput()

	receivedDocs, err := k.GetUserReceivedDocuments(ctx, TestingRecipient)
	assert.Nil(t, err)

	assert.Equal(t, 0, len(receivedDocs))
}

func TestKeeper_GetUserReceivedDocuments_NonEmptyList(t *testing.T) {
	cdc, ctx, k := SetupTestInput()
	store := ctx.KVStore(k.StoreKey)

	documents := types.Documents{TestingDocument}
	documentIds := types.DocumentIds{TestingDocument.Uuid}

	store.Set(k.getDocumentStoreKey(TestingDocument.Uuid), cdc.MustMarshalBinaryBare(TestingDocument))
	store.Set(k.getReceivedDocumentsIdsStoreKey(TestingRecipient), cdc.MustMarshalBinaryBare(&documentIds))

	receivedDocs, err := k.GetUserReceivedDocuments(ctx, TestingRecipient)
	assert.Nil(t, err)

	assert.Equal(t, 1, len(receivedDocs))
	assert.Equal(t, documents, receivedDocs)
}

func TestKeeper_GetUserSentDocuments_EmptyList(t *testing.T) {
	_, ctx, k := SetupTestInput()

	sentDocuments, err := k.GetUserSentDocuments(ctx, TestingSender)
	assert.Nil(t, err)

	assert.Equal(t, 0, len(sentDocuments))
}

func TestKeeper_GetUserSentDocuments_NonEmptyList(t *testing.T) {
	cdc, ctx, k := SetupTestInput()
	store := ctx.KVStore(k.StoreKey)

	documents := types.Documents{TestingDocument}
	documentIds := types.DocumentIds{TestingDocument.Uuid}
	store.Set(k.getDocumentStoreKey(TestingDocument.Uuid), cdc.MustMarshalBinaryBare(TestingDocument))
	store.Set(k.getSentDocumentsIdsStoreKey(TestingSender), cdc.MustMarshalBinaryBare(&documentIds))

	sentDocuments, err := k.GetUserSentDocuments(ctx, TestingSender)
	assert.Nil(t, err)

	assert.Equal(t, 1, len(sentDocuments))
	assert.Equal(t, documents, sentDocuments)
}

func TestKeeper_GetDocuments_EmptyList(t *testing.T) {
	_, ctx, k := SetupTestInput()
	documents := k.GetDocuments(ctx)
	assert.Empty(t, documents)
}

func TestKeeper_GetDocuments_ExistingList(t *testing.T) {
	_, ctx, k := SetupTestInput()

	doc1 := TestingDocument
	doc2 := types.Document{
		Uuid:           "uuid-2",
		Sender:         TestingDocument.Sender,
		Recipients:     TestingDocument.Recipients,
		Metadata:       TestingDocument.Metadata,
		ContentUri:     TestingDocument.ContentUri,
		Checksum:       TestingDocument.Checksum,
		EncryptionData: TestingDocument.EncryptionData,
	}
	assert.NoError(t, k.SaveDocument(ctx, doc1))
	assert.NoError(t, k.SaveDocument(ctx, doc2))

	docs := k.GetDocuments(ctx)
	assert.Len(t, docs, 2)
	assert.Contains(t, docs, doc1)
	assert.Contains(t, docs, doc2)
}

// ----------------------------------
// --- Document receipts
// ----------------------------------

func TestKeeper_SaveDocumentReceipt_EmptyList(t *testing.T) {
	cdc, ctx, k := SetupTestInput()
	store := ctx.KVStore(k.StoreKey)

	assert.NoError(t, k.SaveReceipt(ctx, TestingDocumentReceipt))

	var stored types.DocumentReceiptsIds
	docReceiptBz := store.Get(k.getSentReceiptsIdsStoreKey(TestingDocumentReceipt.Sender))
	cdc.MustUnmarshalBinaryBare(docReceiptBz, &stored)

	assert.Equal(t, 1, len(stored))
	assert.Contains(t, stored, TestingDocumentReceipt.Uuid)
}

func TestKeeper_SaveDocumentReceipt_ExistingReceipt(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	var existing = types.DocumentReceiptsIds{TestingDocumentReceipt.Uuid}

	store := ctx.KVStore(k.StoreKey)
	store.Set(k.getSentReceiptsIdsStoreKey(TestingDocumentReceipt.Sender), cdc.MustMarshalBinaryBare(&existing))

	assert.NoError(t, k.SaveReceipt(ctx, TestingDocumentReceipt))

	var stored types.DocumentReceiptsIds
	docReceiptBz := store.Get(k.getSentReceiptsIdsStoreKey(TestingDocumentReceipt.Sender))
	cdc.MustUnmarshalBinaryBare(docReceiptBz, &stored)

	assert.Equal(t, 1, len(stored))
	assert.Contains(t, stored, TestingDocumentReceipt.Uuid)
}

func TestKeeper_SaveDocumentReceipt_ExistingReceipt_DifferentUuid(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	var existing = types.DocumentReceiptsIds{TestingDocumentReceipt.Uuid}

	store := ctx.KVStore(k.StoreKey)
	store.Set(k.getSentReceiptsIdsStoreKey(TestingDocumentReceipt.Sender), cdc.MustMarshalBinaryBare(&existing))

	newReceipt := types.DocumentReceipt{
		Uuid:         TestingDocumentReceipt.Uuid + "-new",
		Sender:       TestingDocumentReceipt.Sender,
		Recipient:    TestingDocumentReceipt.Recipient,
		TxHash:       TestingDocumentReceipt.TxHash,
		DocumentUuid: TestingDocumentReceipt.DocumentUuid,
		Proof:        TestingDocumentReceipt.Proof,
	}
	assert.NoError(t, k.SaveReceipt(ctx, newReceipt))

	var stored types.DocumentReceiptsIds
	docReceiptBz := store.Get(k.getSentReceiptsIdsStoreKey(TestingDocumentReceipt.Sender))
	cdc.MustUnmarshalBinaryBare(docReceiptBz, &stored)

	assert.Equal(t, 2, len(stored))
	assert.Contains(t, stored, TestingDocumentReceipt.Uuid)
	assert.Contains(t, stored, newReceipt.Uuid)
}

func TestKeeper_GetUserReceivedReceipts_EmptyList(t *testing.T) {
	_, ctx, k := SetupTestInput()

	receipts := k.GetUserReceivedReceipts(ctx, TestingDocumentReceipt.Recipient)

	assert.Empty(t, receipts)
}

func TestKeeper_GetUserReceivedReceipts_FilledList(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	var ids = types.DocumentReceiptsIds{TestingDocumentReceipt.Uuid}

	store := ctx.KVStore(k.StoreKey)
	store.Set(k.getReceivedReceiptsIdsStoreKey(TestingDocumentReceipt.Recipient), cdc.MustMarshalBinaryBare(&ids))
	store.Set(k.getReceiptStoreKey(TestingDocumentReceipt.Uuid), cdc.MustMarshalBinaryBare(&TestingDocumentReceipt))

	actualReceipts := k.GetUserReceivedReceipts(ctx, TestingDocumentReceipt.Recipient)

	expected := types.DocumentReceipts{TestingDocumentReceipt}
	assert.Equal(t, expected, actualReceipts)
}

func TestKeeper_GetUserReceivedReceiptsForDocument_UuidNotFound(t *testing.T) {
	_, ctx, k := SetupTestInput()
	receipt := k.GetUserReceivedReceiptsForDocument(ctx, TestingDocumentReceipt.Recipient, "111")
	assert.Empty(t, receipt)
}

func TestKeeper_GetUserReceivedReceiptsForDocument_UuidFound(t *testing.T) {
	cdc, ctx, k := SetupTestInput()
	var TestingDocumentReceipt2 = types.DocumentReceipt{
		Sender:       TestingSender2,
		Recipient:    TestingDocumentReceipt.Recipient,
		TxHash:       TestingDocumentReceipt.TxHash,
		DocumentUuid: TestingDocumentReceipt.DocumentUuid,
		Proof:        TestingDocumentReceipt.Proof,
	}

	stored := types.DocumentReceiptsIds{TestingDocumentReceipt.Uuid, TestingDocumentReceipt2.Uuid}

	store := ctx.KVStore(k.StoreKey)
	store.Set(k.getReceivedReceiptsIdsStoreKey(TestingDocumentReceipt.Recipient), cdc.MustMarshalBinaryBare(&stored))
	store.Set(k.getReceiptStoreKey(TestingDocumentReceipt.Uuid), cdc.MustMarshalBinaryBare(&TestingDocumentReceipt))
	store.Set(k.getReceiptStoreKey(TestingDocumentReceipt2.Uuid), cdc.MustMarshalBinaryBare(&TestingDocumentReceipt2))

	actual := k.GetUserReceivedReceiptsForDocument(ctx, TestingDocumentReceipt.Recipient, TestingDocumentReceipt.DocumentUuid)
	expected := types.DocumentReceipts{TestingDocumentReceipt, TestingDocumentReceipt2}
	assert.Equal(t, expected, actual)
}
