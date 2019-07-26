package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"net/http"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(
		fmt.Sprintf("/%s/documents/{%s}", storeName, restName),
		getDocumentMetadataHandler(cliCtx, storeName)).
		Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/%s/documents/{%s}/readers", storeName, restName),
		getDocumentReadersHandler(cliCtx, storeName)).
		Methods("GET")
}

// ----------------------------------
// --- GetDocumentMetadata
// ----------------------------------

func getDocumentMetadataHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramType := vars[restName]

		route := fmt.Sprintf("custom/%s/metadata/%s", storeName, paramType)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// ----------------------------------
// --- GetDocumentReaders
// ----------------------------------

func getDocumentReadersHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramType := vars[restName]

		route := fmt.Sprintf("custom/%s/readers/%s", storeName, paramType)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
