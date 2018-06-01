package sale

import "context"

// Endpoint is the sale endpoint
type Endpoint interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
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
)