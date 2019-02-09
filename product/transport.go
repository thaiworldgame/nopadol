package product

import (
	"net/http"
	"github.com/acoshift/hrpc"
	"encoding/json"
	"fmt"
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
	mux.Handle("/search/barcode", m.Handler(SearchByBarcode(s)))
	mux.Handle("/search/itemcode", m.Handler(SearchByItemCode(s)))
	mux.Handle("/search/itemstock", m.Handler(SearchByItemStockLocation(s)))
	mux.Handle("/search/keyword", m.Handler(SearchByKeyword(s)))
	mux.Handle("/new", m.Handler(MakeNewProduct(s)))
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
	jsonEncoder(w, http.StatusOK, v)
}

func errorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	encoder := jsonEncoder

	var status = http.StatusNoContent

	fmt.Println("Error Encode = ", err)
	switch err.Error() {
	case StatusNotFound.Error():
		status = http.StatusOK
	default:
		status = http.StatusOK
	}

	encoder(w, status, &errorResponse{err.Error()})
}

//type httpError struct {
//	Message string `json:"message"`
//}
//
//func NewHttpTransport(ep Endpoint) http.Handler {
//	mux := http.NewServeMux()
//
//	errorEncoder := func(w http.ResponseWriter, err error) {
//		status := http.StatusInternalServerError
//		fmt.Println("error case =", err)
//		switch err {
//		case ErrProductNotFound:
//			status = http.StatusNotFound
//		}
//		httptransport.EncodeJSON(w, status, &httpError{Message: err.Error()})
//		fmt.Println("transport error =", err.Error())
//	}
//
//	mux.Handle("/search/barcode", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
//		var req SearchByBarcodeRequest
//		err := httptransport.DecodeJSON(r.Body, &req)
//		if err != nil {
//			errorEncoder(w, err)
//			return
//		}
//
//		resp, err := ep.SearchProductByBarCode(r.Context(), &req)
//		if err != nil {
//			errorEncoder(w, err)
//			return
//		}
//		httptransport.EncodeJSON(w, http.StatusInternalServerError, &resp)
//	}))
//
//	//fmt.Println("mux = ",mux)
//
//	return mux
//}
