package pos

import (
	"net/http"
	"fmt"
	"github.com/mrtomyum/nopadol/internal/httptransport"
)

type httpError struct {
	Message string `json:"message"`
}

func NewHttpTransport(ep Endpoint) http.Handler{
	mux := http.NewServeMux()

	errEncoder := func (w http.ResponseWriter, err error){
		status := http.StatusInternalServerError
		fmt.Println("error case= ",err)
		switch err {
		case ErrPosNotFound:
			status = http.StatusNotFound
		}
		httptransport.EncodeJSON(w, status, &httpError{Message: err.Error()})
		fmt.Println("transport error =", err.Error())
	}

	mux.Handle("/new", http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		var req NewPosRequest
		err := httptransport.DecodeJSON(r.Body, &req)
		if err != nil {
			errEncoder(w, err)
			return
		}

		resp, err := ep.NewPos(r.Context(), req)
		if err != nil {
			errEncoder(w, err)
			return
		}
		httptransport.EncodeJSON(w, http.StatusOK, &resp)
	}))

	return mux
}