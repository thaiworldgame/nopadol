package drivethru

import (
	"encoding/json"
	"fmt"
	"github.com/acoshift/hrpc"
	"net/http"
	"github.com/mrtomyum/nopadol/auth"
	"errors"
)

type errorResponse struct {
	Error string `json:"error"`
}
var (
	errMethodNotAllowed = errors.New("auth: method not allowed")
	errForbidden        = errors.New("auth: forbidden")
	errBadRequest       = errors.New("auth: bad request body")
	errUnauthorized     = errors.New("auth: Unauthorized")
)


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
	mux.Handle("/userlogin", m.Handler(userLogIn(s)))
	mux.Handle("/zone", m.Handler(makeListCompany(s)))
	mux.Handle("/machine", m.Handler(makeListMachine(s)))
	mux.Handle("/carbrand", m.Handler(makeSearchCarBranch(s)))
	mux.Handle("/customer", m.Handler(makeSearchCustomer(s)))
	mux.Handle("/item/search", m.Handler(makeItemSearch(s)))

	mux.Handle("/pickup/new", m.Handler(pickupNew(s)))
	mux.Handle("/pickup/manage", m.Handler(managePickup(s)))
	mux.Handle("/checkout/manage", m.Handler(manageCheckout(s)))
	mux.Handle("/queue/list", m.Handler(makeSearchListQueue(s)))

	mux.Handle("/queue/edit", m.Handler(queueEdit(s)))
	mux.Handle("/queue/status", m.Handler(queueStatus(s)))
	mux.Handle("/queue/product", m.Handler(queueProduct(s)))
	mux.Handle("/billing/done", m.Handler(billingDone(s)))


	//mux.Handle("/pickup/new",m.Handler(pickupNew(s)))

	//mux.Handle("/shift/open", m.Handler(makeShiftOpen(s)))

	mux.Handle("/shift/open", m.Handler(makeShiftOpen(s)))
	mux.Handle("/shift/close", m.Handler(makeShiftClose(s)))
	return mustLogin()(mux)
}

func mustLogin() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//companyID := auth.GetCompanyID(r.Context())
			auth.GetCompanyID(r.Context())
			// remark ชั่วคราว
			//if companyID <=  0 {
			//	errorEncoder(w, r, errForbidden)
			//	fmt.Println("error mustLogin auth.transport.go")
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
