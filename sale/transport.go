package sale

import (
	"net/http"

	"github.com/mrtomyum/nopadol/internal/httptransport"
)

type httpError struct {
	Message string `json:"message"`
}

// NewHTTPTransport creates new HTTP transport for domain Sale
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

	//Call API Insert POS
	//mux.Handle("/pos/new",http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	//	var req NewPosRequest
	//	err := httptransport.DecodeJSON(r.Body, &req)
	//	if err != nil {
	//		errorEncoder(w, err)
	//		return
	//	}
	//	resp, err := ep.NewPos(r.Context(), req)
	//	if err != nil {
	//		errorEncoder(w, err)
	//		return
	//	}
	//	httptransport.EncodeJSON(w, http.StatusOK, &resp)
	//}))

	mux.Handle("/search", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req SearchSaleOrderRequest
		err := httptransport.DecodeJSON(r.Body, &req)
		if err != nil {
			errorEncoder(w, err)
			return
		}
		resp, err := ep.Search(r.Context(), &req)
		if err != nil {
			errorEncoder(w, err)
			return
		}
		httptransport.EncodeJSON(w, http.StatusOK, &resp)
	}))
	//
	//mux.Handle("/create", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	var req CreateRequest
	//	err := httptransport.DecodeJSON(r.Body, &req)
	//	if err != nil {
	//		errorEncoder(w, err)
	//		return
	//	}
	//	resp, err := ep.Create(r.Context(), &req)
	//	if err != nil {
	//		errorEncoder(w, err)
	//		return
	//	}
	//	httptransport.EncodeJSON(w, http.StatusOK, &resp)
	//}))
	//

	mux.Handle("/so/new",http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var req NewSaleOrderRequest
		err := httptransport.DecodeJSON(r.Body, &req)
		if err != nil {
			errorEncoder(w, err)
			return
		}
		resp, err := ep.NewSaleOrder(r.Context(), req)
		if err != nil {
			errorEncoder(w, err)
			return
		}
		httptransport.EncodeJSON(w, http.StatusOK, resp)
	}))

	// mux.Handle("/new", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	var req NewRequest
	// 	err := httptransport.DecodeJSON(r.Body, &req)
	// 	if err != nil {
	// 		errorEncoder(w, err)
	// 		return
	// 	}
	// 	resp, err := ep.New(r.Context(), &req)
	// 	if err != nil {
	// 		errorEncoder(w, err)
	// 		return
	// 	}
	// 	httptransport.EncodeJSON(w, http.StatusOK, resp)
	// }))

	// or use https://github.com/acoshift/hrpc for RPC-HTTP style API
	// mux.Handle("/create", m.Handler(ep.Create))

	return mux
}
