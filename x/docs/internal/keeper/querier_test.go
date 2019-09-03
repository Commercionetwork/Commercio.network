package keeper

import (
	"testing"

	"github.com/commercionetwork/commercionetwork/x/docs/internal/types"
	"github.com/stretchr/testify/assert"
	abci "github.com/tendermint/tendermint/abci/types"
)

var querier = NewQuerier(TestUtils.DocsKeeper)
var request abci.RequestQuery
var documents = types.Documents{TestingDocument}

func Test_queryGetReceivedDocuments(t *testing.T) {
	// Setup the store
	metadataStore := TestUtils.Ctx.KVStore(TestUtils.DocsKeeper.StoreKey)
	metadataStore.Set(
		[]byte(types.ReceivedDocumentsPrefix+TestingDocument.Recipient.String()),
		TestUtils.Cdc.MustMarshalBinaryBare(&documents),
	)

	// Compose the path
	path := []string{"received", TestingDocument.Recipient.String()}

	// Get the returned documents
	var actual types.Documents
	actualBz, _ := querier(TestUtils.Ctx, path, request)
	TestUtils.Cdc.MustUnmarshalJSON(actualBz, &actual)

	assert.Equal(t, documents, actual)
}

func Test_queryGetSentDocuments(t *testing.T) {
	// Setup the store
	metadataStore := TestUtils.Ctx.KVStore(TestUtils.DocsKeeper.StoreKey)
	metadataStore.Set(
		[]byte(types.SentDocumentsPrefix+TestingDocument.Sender.String()),
		TestUtils.Cdc.MustMarshalBinaryBare(&documents),
	)

	// Compose the path
	path := []string{"sent", TestingDocument.Sender.String()}

	// Get the returned documents
	var actual types.Documents
	actualBz, _ := querier(TestUtils.Ctx, path, request)
	TestUtils.Cdc.MustUnmarshalJSON(actualBz, &actual)

	assert.Equal(t, documents, actual)
}

// ----------------------------------
// --- DocumentReceipt
// ----------------------------------

func Test_GetUserReceivedReceipts(t *testing.T) {

	receiptStore := TestUtils.Ctx.KVStore(TestUtils.DocsKeeper.StoreKey)

	//cleanup the store
	receiptStore.Delete([]byte(types.DocumentReceiptPrefix + TestingDocumentReceipt.Uuid + TestingSender2.String()))

	//Setup the store
	receiptStore.Set([]byte(types.DocumentReceiptPrefix+TestingDocumentReceipt.Uuid+TestingDocumentReceipt.Sender.String()),
		TestUtils.Cdc.MustMarshalBinaryBare(&TestingDocumentReceipt))

	var expected = types.DocumentReceipts{TestingDocumentReceipt}

	// Compose the path
	path := []string{"receipts", TestingDocumentReceipt.Recipient.String(), ""}

	//Get the returned receipts
	var actual types.DocumentReceipts
	actualBz, _ := querier(TestUtils.Ctx, path, request)
	TestUtils.Cdc.MustUnmarshalJSON(actualBz, &actual)

	assert.Equal(t, expected, actual)
}

func Test_GetUserReceivedReceiptsForDocument(t *testing.T) {

	//Setup the store
	receiptStore := TestUtils.Ctx.KVStore(TestUtils.DocsKeeper.StoreKey)
	receiptStore.Set([]byte(types.DocumentReceiptPrefix+TestingDocumentReceipt.Uuid+TestingDocumentReceipt.Sender.String()),
		TestUtils.Cdc.MustMarshalBinaryBare(&TestingDocumentReceipt))

	// Compose the path
	path := []string{"receipts", TestingDocumentReceipt.Recipient.String(), TestingDocumentReceipt.Uuid}

	//Get the returned receipts
	var actual types.DocumentReceipts
	actualBz, _ := querier(TestUtils.Ctx, path, request)
	TestUtils.Cdc.MustUnmarshalJSON(actualBz, &actual)

	var expected = types.DocumentReceipts{TestingDocumentReceipt}
	assert.Equal(t, expected, actual)
}
