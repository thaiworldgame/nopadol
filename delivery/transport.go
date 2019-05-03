package delivery

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/acoshift/hrpc"
)

var (
	errMethodNotAllowed = errors.New("delivery : method not allowed")
	errForbidden        = errors.New("delivery : forbidden")
	errBadRequest       = errors.New("Sale: bad request body")
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

type errorResponse struct {
	Error string `json:"error"`
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

	mux.Handle("/report", m.Handler(makeReportDoData(s)))
	return mustLogin()(mux)
}

func mustLogin() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//clientID := auth.GetClientID(r.Context())
			//if clientID < 0 {
			//	errorEncoder(w, r, errForbidden)
			//	return
			//}
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
	//w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func requestDecoder(r *http.Request, v interface{}) error {
	if r.Method != http.MethodPost {
		return errMethodNotAllowed
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

	switch err {
	case errMethodNotAllowed:
		status = http.StatusMethodNotAllowed
	case errForbidden:
		status = http.StatusForbidden
	}

	encoder(w, status, &errorResponse{err.Error()})
}
