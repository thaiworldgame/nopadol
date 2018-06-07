package sale

import "context"

// Endpoint is the sale endpoint
type Endpoint interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Search(context.Context, *SearchSaleRequest) (*SearchSaleResponse, error)
	NewSO(context.Context, *NewSORequest) (*NewSOResponse, error)
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

	NewSORequest struct {
		Sale SaleOrder
	}

	NewSOResponse struct {
		Id int64 `json:"id"`
	}

	SearchSaleRequest struct {
		Keyword string `json:"keyword"`
	}

	SearchSaleResponse struct {
		Sale SaleOrder
	}
)
