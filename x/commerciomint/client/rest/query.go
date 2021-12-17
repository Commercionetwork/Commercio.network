package rest

import (
	"fmt"
	"net/http"

	"github.com/commercionetwork/commercionetwork/x/commerciomint/types"
	"github.com/cosmos/cosmos-sdk/client"
	restTypes "github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

const (
	restuser = "user"
)

func RegisterRoutes(cliCtx client.Context, r *mux.Router) {
	r.HandleFunc(
		fmt.Sprintf("/commercionetwork/{%s}/etp", restuser),
		getEtpsHandler(cliCtx)).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/commercionetwork/{%s}/owner", restuser),
		getEtpsByOwnerHandler(cliCtx)).Methods("GET")
	r.HandleFunc("/commercionetwork/etps", getAllEtpsHandler(cliCtx)).Methods("GET")
	r.HandleFunc("/commercionetwork/conversion_rate", getConversionRateHandler(cliCtx)).Methods("GET")
	r.HandleFunc("/commercionetwork/freeze_period", getFreezePeriodHandler(cliCtx)).Methods("GET")
}

// ----------------------------------
// --- Commerciomint
// ----------------------------------

func getEtpsHandler(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars[restuser]

		route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, id, types.QueryGetEtp)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			restTypes.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		restTypes.PostProcessResponse(w, cliCtx, res)
	}
}

func getEtpsByOwnerHandler(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ownerAddr := vars[restuser]

		route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, ownerAddr, types.QueryGetEtpsByOwner)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			restTypes.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		restTypes.PostProcessResponse(w, cliCtx, res)
	}
}
func getAllEtpsHandler(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryGetallEtps)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			restTypes.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		restTypes.PostProcessResponse(w, cliCtx, res)
	}
}

func getConversionRateHandler(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryConversionRateRest)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			restTypes.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return 
		}
		restTypes.PostProcessResponse(w, cliCtx, res)
	}
}

func getFreezePeriodHandler(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryFreezePeriodRest)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			restTypes.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return 
		}
		restTypes.PostProcessResponse(w, cliCtx, res)
	}
}
