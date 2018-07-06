package product

import (
	"net/http"
	"github.com/mrtomyum/nopadol/internal/httptransport"
	"fmt"
)

type httpError struct {
	Message string `json:"message"`
}

func NewHttpTransport(ep Endpoint) http.Handler {
	mux := http.NewServeMux()

	errorEncoder := func(w http.ResponseWriter, err error) {
		status := http.StatusInternalServerError
		fmt.Println("error case =", err)
		switch err {
		case ErrProductNotFound:
			status = http.StatusNotFound
		}
		httptransport.EncodeJSON(w, status, &httpError{Message: err.Error()})
		fmt.Println("transport error =", err.Error())
	}

	mux.Handle("/search/barcode", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var req SearchByBarcodeRequest
		err := httptransport.DecodeJSON(r.Body, &req)
		if err != nil {
			errorEncoder(w, err)
			return
		}

		resp, err := ep.SearchProductByBarCode(r.Context(), &req)
		if err != nil {
			errorEncoder(w, err)
			return
		}
		httptransport.EncodeJSON(w, http.StatusInternalServerError, &resp)
	}))

	//fmt.Println("mux = ",mux)

	return mux
}
