package sync

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/acoshift/hrpc"
)

type errorResponse struct {
	Error string `json:"error"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func MakeHandler(s Service) http.Handler {
	m := hrpc.New(hrpc.Config{
		Validate:        true,
		RequestDecoder:  requestDecoder,
		ResponseEncoder: responseEncoder,
		ErrorEncoder:    errorEncoder,
	})
	mux := http.NewServeMux()
	mux.Handle("/quotation/new", m.Handler(NewQuotation(s)))
	mux.Handle("/saleorder/new", m.Handler(NewSaleOrder(s)))
	mux.Handle("/done", m.Handler(ConfirmTransfer(s)))

	return mustLogin()(mux)

}

func mustLogin() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			enableCors(&w)
			h.ServeHTTP(w, r)
		})
	}
}

func jsonDecoder(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func jsonEncoder(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func requestDecoder(r *http.Request, v interface{}) error {
	if r.Method != http.MethodPost {
		return ErrMethodNotAllowed
	}

	return jsonDecoder(r, v)
}

func responseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) {
	fmt.Println("v =", v)
	jsonEncoder(w, http.StatusOK, v)
}

func errorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	encoder := jsonEncoder

	var status = http.StatusOK

	fmt.Println("Error Encode = ", err.Error())
	switch err.Error() {

	default:
		status = http.StatusOK
	}

	encoder(w, status, &errorResponse{err.Error()})
}
