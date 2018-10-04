package sales

import (
	"net/http"
	"encoding/json"
	"fmt"
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
	mux.Handle("/quo/new",m.Handler(CreateQuo(s)))
	mux.Handle("/quo/search/id", m.Handler(SearchQuoById(s)))
	mux.Handle("/sale/new", m.Handler(CreateSale(s)))
	mux.Handle("/sale/search/id", m.Handler(SearchSaleById(s)))
	mux.Handle("/sale/doc/search",m.Handler(SearchDocByKeyword(s)))
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

	var status = http.StatusNoContent

	fmt.Println("Error Encode = ", err.Error())
	switch err.Error() {
	case StatusNotFound.Error():
		status = http.StatusNotFound
	case ArCodeNull.Error():
		status = http.StatusNotFound
	case NotHaveItem.Error():
		status = http.StatusNotFound
	case NotHavePayMoney.Error():
		status = http.StatusNotFound
	case NotHaveSumOfItem.Error():
		status = http.StatusNotFound
	case ItemNotHaveQty.Error():
		status = http.StatusNotFound
	case ItemNotHaveUnit.Error():
		status = http.StatusNotFound
	case MoneyOverTotalAmount.Error():
		status = http.StatusNotFound
	case MoneyLessThanTotalAmount.Error():
		status = http.StatusNotFound
	case PosNotHaveDate.Error():
		status = http.StatusNotFound
	case PosNotHaveChqData.Error():
		status = http.StatusNotFound
	case PosNotHaveCreditCardData.Error():
		status = http.StatusNotFound
	default:
		status = http.StatusForbidden
	}

	encoder(w, status, &errorResponse{err.Error()})
}
