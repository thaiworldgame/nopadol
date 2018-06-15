package sale

import "context"

// Endpoint is the sale endpoint
type Endpoint interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	//Search(context.Context, *SearchSaleOrderRequest) (*SearchSaleOrderResponse, error)
	NewSaleOrder(context.Context, NewSaleOrderRequest) (*NewSaleOrderResponse, error)
	Search(context.Context, *SearchSaleOrderRequest) (*SearchSaleOrderResponse, error)
}

// Create
type (
	// CreateRequest is the request for create endpoint
	CreateRequest struct {
		Field1 string `json:"field1"`
	}

	// CreateResponse is the response for create endpoint
	CreateResponse struct {
		ID string `json:"id"`
	}

	// CreateRequest is the request for create endpoint
	//NewSORequest struct {
	//	//Sale SaleOrder `json:"sale"`
	//	DocNo   string `json:"doc_no"`
	//	DocDate string `json:"doc_date"`
	//	ArCode  string `json:"ar_code"`
	//	ArName  string `json:"ar_name"`
	//	Subs []*struct {
	//		ItemCode string  `json:"item_code"`
	//		ItemName string  `json:"item_name"`
	//		Qty      float64 `json:"qty"`
	//		UnitCode string  `json:"unit_code"`
	//	} `json:"subs"`
	//}
	//
	//// CreateResponse is the response for create endpoint
	//NewSOResponse struct {
	//	SOID int64 `json:"soid"`
	//}

	NewSaleOrderRequest struct {
		DocNo   string                  `json:"doc_no"`
		DocDate string                  `json:"doc_date"`
		ArCode  string                  `json:"ar_code"`
		ArName  string                  `json:"ar_name"`
		Subs    []NewSubsSaleOrderRequest `json:"subs"`
	}

	NewSubsSaleOrderRequest struct {
		ItemCode string  `json:"item_code"`
		ItemName string  `json:"item_name"`
		Qty      float64 `json:"qty"`
		UnitCode string  `json:"unit_code"`
	}

	NewSaleOrderResponse struct {
		Id int64 `json:"doc_no"`
	}

	SearchSaleOrderRequest struct {
		Keyword string `json:"keyword"`
	}

	SearchSaleOrderResponse struct {
		DocNo   string                  `json:"doc_no"`
		DocDate string                  `json:"doc_date"`
		ArCode  string                  `json:"ar_code"`
		ArName  string                  `json:"ar_name"`
		Subs    []SubsSaleOrderResponse `json:"subs"`
	}

	SubsSaleOrderResponse struct {
		ItemCode string  `json:"item_code"`
		ItemName string  `json:"item_name"`
		Qty      float64 `json:"qty"`
		UnitCode string  `json:"unit_code"`
	}
)
