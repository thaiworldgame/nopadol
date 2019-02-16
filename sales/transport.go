package sales

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
	mux.Handle("/quo/new", m.Handler(CreateQuotation(s)))
	mux.Handle("/quo/search/id", m.Handler(SearchQuoById(s)))
	mux.Handle("/quo/confirm", m.Handler(ConfirmQuotation(s)))
	mux.Handle("/quo/cancel", m.Handler(CancelQuotation(s)))
	mux.Handle("/sale/new", m.Handler(CreateSaleOrder(s)))
	mux.Handle("/sale/search/id", m.Handler(SearchSaleOrderById(s)))
	mux.Handle("/sale/doc/search", m.Handler(SearchDocByKeyword(s)))
	mux.Handle("/dep/new", m.Handler(CreateDeposit(s)))
	mux.Handle("/dep/search/id", m.Handler(SearchDepositById(s)))
	mux.Handle("/dep/search/keyword", m.Handler(SearchDepositByKeyword(s)))
	mux.Handle("/dep/reserve/search", m.Handler(SearchReserveToDeposit(s)))
	mux.Handle("/inv/search/id", m.Handler(SearchInvoiceById(s)))
	mux.Handle("/inv/new", m.Handler(CreateInvoice(s)))
	mux.Handle("/inv/search/keyword", m.Handler(SearchInvoiceByKeyword(s)))
	mux.Handle("/sale/search/item", m.Handler(SearchSaleByItem(s)))
	mux.Handle("/search/credit", m.Handler(SearchCredit(s)))
	mux.Handle("/inv/list", m.Handler(Invoicelist(s)))
	mux.Handle("/his/search/keyword", m.Handler(SearchHisByKeyword(s)))
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
	case StatusNotFound.Error():
		status = http.StatusOK
	case ArCodeNull.Error():
		status = http.StatusOK
	case NotHaveItem.Error():
		status = http.StatusOK
	case NotHavePayMoney.Error():
		status = http.StatusOK
	case NotHaveSumOfItem.Error():
		status = http.StatusOK
	case ItemNotHaveQty.Error():
		status = http.StatusOK
	case ItemNotHaveUnit.Error():
		status = http.StatusOK
	case MoneyOverTotalAmount.Error():
		status = http.StatusOK
	case MoneyLessThanTotalAmount.Error():
		status = http.StatusOK
	case PosNotHaveDate.Error():
		status = http.StatusOK
	case PosNotHaveChqData.Error():
		status = http.StatusOK
	case PosNotHaveCreditCardData.Error():
		status = http.StatusOK
	default:
		status = http.StatusOK
	}

	encoder(w, status, &errorResponse{err.Error()})
}
