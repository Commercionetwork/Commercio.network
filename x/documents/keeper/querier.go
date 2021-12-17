package keeper

import (
	"fmt"

	sdkErr "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/commercionetwork/commercionetwork/x/documents/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryReceivedDocuments:
			return queryGetReceivedDocuments(ctx, path[1:], k, legacyQuerierCdc)
		case types.QuerySentDocuments:
			return queryGetSentDocuments(ctx, path[1:], k, legacyQuerierCdc)
		case types.QueryReceivedReceipts:
			return queryGetReceivedDocsReceipts(ctx, path[1:], k, legacyQuerierCdc)
		case types.QuerySentReceipts:
			return queryGetSentDocsReceipts(ctx, path[1:], k, legacyQuerierCdc)
		default:
			return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest, fmt.Sprintf("Unknown %s query endpoint: %s", types.ModuleName, path[0]))
		}
	}
}

// ----------------------------------
// --- Documents
// ----------------------------------

func queryGetReceivedDocuments(ctx sdk.Context, path []string, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	addr := path[0]
	address, _ := sdk.AccAddressFromBech32(addr)

	ri := k.UserReceivedDocumentsIterator(ctx, address)
	defer ri.Close()

	receivedResult := []types.Document{}
	for ; ri.Valid(); ri.Next() {
		documentUUID := types.StringProto{}
		k.cdc.MustUnmarshalBinaryBare(ri.Value(), &documentUUID)

		document, err := k.GetDocumentByID(ctx, documentUUID.StProto)
		if err != nil {
			return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest,
				fmt.Sprintf(
					"could not find document with UUID %s even though the user has an associated received document",
					documentUUID,
				),
			)
		}

		receivedResult = append(receivedResult, document)
	}

	bz, err2 := codec.MarshalJSONIndent(legacyQuerierCdc, receivedResult)
	if err2 != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest, "Could not marshal result to JSON")
	}

	return bz, nil
}

func queryGetSentDocuments(ctx sdk.Context, path []string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	addr := path[0]
	address, _ := sdk.AccAddressFromBech32(addr)

	usdi := keeper.UserSentDocumentsIterator(ctx, address)
	defer usdi.Close()

	receivedResult := []types.Document{}
	for ; usdi.Valid(); usdi.Next() {
		documentUUID := types.StringProto{}
		keeper.cdc.MustUnmarshalBinaryBare(usdi.Value(), &documentUUID)

		document, err := keeper.GetDocumentByID(ctx, documentUUID.StProto)
		if err != nil {
			return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest,
				fmt.Sprintf(
					"could not find document with UUID %s even though the user has an associated received document",
					documentUUID,
				),
			)
		}

		receivedResult = append(receivedResult, document)
	}

	bz, err2 := codec.MarshalJSONIndent(legacyQuerierCdc, receivedResult)
	if err2 != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest, "Could not marshal result to JSON")
	}

	return bz, nil
}

// ----------------------------------
// --- Documents receipts
// ----------------------------------

func queryGetReceivedDocsReceipts(ctx sdk.Context, path []string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	addr := path[0]
	address, _ := sdk.AccAddressFromBech32(addr)

	var uuid string
	if len(path) == 2 {
		uuid = path[1]
	}

	receipts := []types.DocumentReceipt{}

	ri := keeper.UserReceivedReceiptsIterator(ctx, address)
	defer ri.Close()

	for ; ri.Valid(); ri.Next() {
		rid := types.StringProto{}
		keeper.cdc.MustUnmarshalBinaryBare(ri.Value(), &rid)

		newReceipt, err := keeper.GetReceiptByID(ctx, rid.StProto)
		if err != nil {
			return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest,
				fmt.Sprintf(
					"could not find document receipt with UUID %s even though the user has an associated received document with it",
					rid,
				),
			)
		}

		if uuid == "" {
			receipts = append(receipts, newReceipt)
			continue
		}

		if newReceipt.DocumentUUID == uuid {
			receipts = append(receipts, newReceipt)
		}
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &receipts)

	if err != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest, "Could not marshal result to JSON")
	}

	return bz, nil
}

func queryGetSentDocsReceipts(ctx sdk.Context, path []string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	addr := path[0]
	address, err := sdk.AccAddressFromBech32(addr)

	if err != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrInvalidAddress, addr)
	}

	receipts := []types.DocumentReceipt{}

	ri := keeper.UserSentReceiptsIterator(ctx, address)
	defer ri.Close()

	for ; ri.Valid(); ri.Next() {
		rid := types.StringProto{}
		keeper.cdc.MustUnmarshalBinaryBare(ri.Value(), &rid)

		newReceipt, err := keeper.GetReceiptByID(ctx, rid.StProto)
		if err != nil {
			return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest,
				fmt.Sprintf(
					"could not find document receipt with UUID %s even though the user has an associated received document with it",
					rid,
				),
			)
		}

		receipts = append(receipts, newReceipt)
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &receipts)

	if err != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest, "Could not marshal result to JSON")
	}

	return bz, nil
}
