package rest

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	// sdk "github.com/enigmampc/cosmos-sdk/types"
	"github.com/enigmampc/cosmos-sdk/types/rest"
	ra "github.com/enigmampc/SecretNetwork/x/registration/remote_attestation"

	"github.com/enigmampc/SecretNetwork/x/registration/internal/keeper"
	"github.com/enigmampc/SecretNetwork/x/registration/internal/types"

	"github.com/enigmampc/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/reg/code", listCodesHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/reg/consensus-io-exch-pubkey", ioPubkeyHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/reg/consensus-seed-exch-pubkey", seedPubkeyHandlerFn(cliCtx)).Methods("GET")
}

func listCodesHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, keeper.QueryEncryptedSeed)
		res, height, err := cliCtx.Query(route)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, json.RawMessage(res))
	}
}

func ioPubkeyHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, keeper.QueryMasterCertificate)
		res, height, err := cliCtx.Query(route)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		cliCtx = cliCtx.WithHeight(height)

		var certs types.GenesisState

		err = json.Unmarshal(res, &certs)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		ioExchPubkey, err := ra.VerifyRaCert(certs.IoMasterCertificate)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res = []byte(fmt.Sprintf(`{"ioExchPubkey":"%s"}`, base64.StdEncoding.EncodeToString(ioExchPubkey)))

		rest.PostProcessResponse(w, cliCtx, json.RawMessage(res))
	}
}

func seedPubkeyHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, keeper.QueryMasterCertificate)
		res, height, err := cliCtx.Query(route)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		cliCtx = cliCtx.WithHeight(height)

		var certs types.GenesisState

		err = json.Unmarshal(res, &certs)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		nodeExchPubkey, err := ra.VerifyRaCert(certs.NodeExchMasterCertificate)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res = []byte(fmt.Sprintf(`{"nodeExchPubkey":"%s"}`, base64.StdEncoding.EncodeToString(nodeExchPubkey)))

		rest.PostProcessResponse(w, cliCtx, nodeExchPubkey)
	}
}

//
//func queryCodeHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		codeID, err := strconv.ParseUint(mux.Vars(r)["codeID"], 10, 64)
//		if err != nil {
//			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
//			return
//		}
//
//		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
//		if !ok {
//			return
//		}
//
//		route := fmt.Sprintf("custom/%s/%s/%d", types.QuerierRoute, keeper.QueryGetCode, codeID)
//		res, height, err := cliCtx.Query(route)
//		if err != nil {
//			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
//			return
//		}
//		if len(res) == 0 {
//			rest.WriteErrorResponse(w, http.StatusNotFound, "contract not found")
//			return
//		}
//
//		cliCtx = cliCtx.WithHeight(height)
//		rest.PostProcessResponse(w, cliCtx, json.RawMessage(res))
//	}
//}

//type smartResponse struct {
//	Smart []byte `json:"smart"`
//}

type argumentDecoder struct {
	// dec is the default decoder
	dec      func(string) ([]byte, error)
	encoding string
}

func newArgDecoder(def func(string) ([]byte, error)) *argumentDecoder {
	return &argumentDecoder{dec: def}
}

func (a *argumentDecoder) DecodeString(s string) ([]byte, error) {

	switch a.encoding {
	case "hex":
		return hex.DecodeString(s)
	case "base64":
		return base64.StdEncoding.DecodeString(s)
	default:
		return a.dec(s)
	}
}
