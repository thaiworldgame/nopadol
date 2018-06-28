package employee

import (
	"net/http"
	"github.com/mrtomyum/nopadol/internal/httptransport"
)

type httpError struct {
	Message string `json:"message"`
}

func NewHttpTransport(ep Endpoint) http.Handler {
	mux := http.NewServeMux()

	errorEncoder := func(w http.ResponseWriter, err error) {
		status := http.StatusInternalServerError
		switch err {
		case ErrEmployeeNotFound:
			status = http.StatusNotFound
		}
		httptransport.EncodeJSON(w, status, &httpError{Message: err.Error()})
	}

	mux.Handle("/search/id", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req SearchEmployeeByIdRequest
		err := httptransport.DecodeJSON(r.Body, &req)
		if err != nil {
			errorEncoder(w, err)
			return
		}
		resp, err := ep.SearchEmployeeById(r.Context(), &req)
		if err != nil {
			errorEncoder(w, err)
			return
		}
		httptransport.EncodeJSON(w, http.StatusOK, &resp)
	}))

	return mux
}
