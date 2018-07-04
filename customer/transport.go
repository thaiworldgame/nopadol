package customer

import (
	"net/http"
	"github.com/mrtomyum/nopadol/internal/httptransport"
)

type httpError struct {
	Message string `json:"message"`
}

type Response struct {
	Status  string `json:"status"`
	Message string  `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}


func NewHttpTransport(ep Endpoint) http.Handler {
	mux := http.NewServeMux()

	//errorEncoder := func(w http.ResponseWriter, err error){
	//	status := http.StatusInternalServerError
	//	fmt.Println(err)
	//	switch err {
	//	case ErrCustomerNotFound:
	//		status = http.StatusNotFound
	//	}
	//	httptransport.EncodeJSON(w, status, &httpError{Message:err.Error()})
	//}

	mux.Handle("/search/id", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		r.Header.Set("Access-Control-Allow-Origin", "*")
		r.Header.Set("Access-Control-Allow-Headers", "Content-Type,Token")
		r.Header.Set("Content_Type", "application/json")
		r.Header.Set("Access-Control-Allow-Methods","GET, POST, PUT, DELETE")

		var req SearchCustomerByIdRequest
		err := httptransport.DecodeJSON(r.Body, &req)
		if err != nil {
			//errorEncoder(w, err)
			//return
		}

		resp, err := ep.SearchCustomerById(r.Context(), &req)
		if err != nil {
			//errorEncoder(w, err)
			//return
		}

		rs := Response{}
		if err != nil {
			rs.Status = "error"
			rs.Message = "No Content and Error :"+ err.Error()
		}else{
			rs.Status = "success"
			rs.Data = resp
		}


		httptransport.EncodeJSON(w, http.StatusOK, rs)
	}))


	////Call API Insert POS
	//mux.Handle("/search/id",http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	//	var req SearchCustomerIdRequest
	//	err := httptransport.DecodeJSON(r.Body, &req)
	//	if err != nil {
	//		errorEncoder(w, err)
	//		return
	//	}
	//	resp, err := ep.SearchCustomerById(r.Context(), req)
	//	if err != nil {
	//		errorEncoder(w, err)
	//		return
	//	}
	//	httptransport.EncodeJSON(w, http.StatusOK, &resp)
	//}))


	return mux
}