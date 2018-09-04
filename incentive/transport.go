package incentive

import (
	"net/http"

	"github.com/mrtomyum/nopadol/internal/httptransport"
)

type httpError struct {
	Message string `json:"message"`
}

// NewHTTPTransport creates new HTTP transport for domain1
func NewHTTPTransport(ep Endpoint) http.Handler {
	mux := http.NewServeMux()

	errorEncoder := func(w http.ResponseWriter, err error) {
		status := http.StatusInternalServerError
		switch err {
		case ErrEntity1NotFound:
			status = http.StatusNotFound
		}
		httptransport.EncodeJSON(w, status, &httpError{Message: err.Error()})
	}

	mux.Handle("/searchsale", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req SearchSaleCodeRequest
		err := httptransport.DecodeJSON(r.Body, &req)
		if err != nil {
			errorEncoder(w, err)
			return
		}
		resp, err := ep.SearchSaleCode(r.Context(), &req)
		if err != nil {
			errorEncoder(w, err)
			return
		}
		httptransport.EncodeJSON(w, http.StatusOK, &resp)
	}))
	return mux
}

