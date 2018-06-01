package endpoint

import (
	"context"
	"github.com/mrtomyum/nopadol/sale"
)

// New creates new domain1 endpoint
func New(s sale.Service) sale.Endpoint {
	return &endpoint{s}
}

type endpoint struct {
	s sale.Service
}

func (ep *endpoint) Create(ctx context.Context, req *sale.CreateRequest) (*sale.CreateResponse, error) {
	id, err := ep.s.Create(ctx, &sale.Entity1{
		Field2: sale.Entity2{
			Field1: req.Field1,
		},
	})
	if err != nil {
		return nil, err
	}
	return &sale.CreateResponse{ID: id}, nil
}


