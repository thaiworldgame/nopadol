package sale

import "context"

// Endpoint is the sale endpoint
type Endpoint interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Search(context.Context, *SearchSaleRequest)(*SearchSaleResponse, error)
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

	SearchSaleRequest struct {
		Keyword string `json:"keyword"`
	}

	SearchSaleResponse struct {
		DocNo string `json:"doc_no"`
		ArCode string `json:"ar_code"`
		ArName string `json:"ar_name"`
	}
)