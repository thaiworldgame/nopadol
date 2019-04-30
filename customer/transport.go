package customer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/acoshift/hrpc"
)

//type httpError struct {
//	Message string `json:"message"`
//}

type errorResponse struct {
	Error string `json:"error"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

}

// MakeHandler creates new vending  handler
func MakeHandler(s Service) http.Handler {
	// m := hrpc.New(hrpc.Config{
	// 	Validate:        true,
	// 	RequestDecoder:  requestDecoder,
	// 	ResponseEncoder: responseEncoder,
	// 	ErrorEncoder:    errorEncoder,
	// })
	// mux := http.NewServeMux()

	m := hrpc.Manager{
		Validate:     true,
		Decoder:      requestDecoder,
		Encoder:      responseEncoder,
		ErrorEncoder: errorEncoder,
	}
	mux := http.NewServeMux()

	mux.Handle("/search/id", m.Handler(SearchById(s)))
	mux.Handle("/search/keyword", m.Handler(SearchByKeyword(s)))

	mux.Handle("/new", m.Handler(makeNewCustomer(s)))
	mux.Handle("/update", m.Handler(makeUpdateCustomer(s)))
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
	// TODO: choose decoder from request's content type
	// right now we have only json decoder
	return jsonDecoder(r, v)
}

func responseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) {
	// TODO: choose encoder from request's accept
	// right now we have only json encoder
	jsonEncoder(w, http.StatusOK, v)
}

func errorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	// TODO: choose encoder from request's accept
	encoder := jsonEncoder

	var status = http.StatusNoContent

	fmt.Println("Error Encode = ", err)
	switch err.Error() {
	case StatusNotFound.Error():
		status = http.StatusOK
	case ErrMethodNotAllowed.Error():
		status = http.StatusOK
	case ErrForbidden.Error():
		status = http.StatusOK
	default:
		status = http.StatusOK
	}

	encoder(w, status, &errorResponse{err.Error()})
}

//func NewHttpTransport(ep Endpoint) http.Handler {
//	mux := http.NewServeMux()
//
//	errorEncoder := func(w http.ResponseWriter, err error) {
//		status := http.StatusInternalServerError
//		fmt.Println(err)
//		switch err {
//		case ErrCustomerNotFound:
//			status = http.StatusNotFound
//		}
//		httptransport.EncodeJSON(w, status, &httpError{Message: err.Error()})
//	}
//
//	mux.Handle("/search/id", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		enableCors(&w)
//		var req SearchCustomerByIdRequest
//		err := httptransport.DecodeJSON(r.Body, &req)
//		if err != nil {
//			errorEncoder(w, err)
//			return
//		}
//
//		resp, err := ep.SearchCustomerById(r.Context(), &req)
//		if err != nil {
//			errorEncoder(w, err)
//			return
//		}
//
//		httptransport.EncodeJSON(w, http.StatusOK, resp)
//	}))
//
//	return mux
//}
