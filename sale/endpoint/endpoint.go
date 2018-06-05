package endpoint

import (
	"context"
	"github.com/mrtomyum/nopadol/sale"
	"fmt"
)

// New creates new domain1 endpoint
func New(s sale.Service) sale.Endpoint {
	return &endpoint{s}
}

type endpoint struct {
	s sale.Service
}

func (ep *endpoint) Create(ctx context.Context, req *sale.CreateRequest) (*sale.CreateResponse, error) {
	fmt.Println("CreateRequest = ",req.Field1)
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


func (ep *endpoint) Search(ctx context.Context, req *sale.SearchSaleRequest) (*sale.SearchSaleResponse, error) {
	fmt.Println("keyword endpoint = ",req.Keyword)

	d, err := ep.s.Search(ctx, &sale.EntitySearch{
		Keyword:req.Keyword,
	})
	if err != nil {
		return nil,err
	}

	fmt.Println("Search By = ",sale.EntitySearch{})

	return &sale.SearchSaleResponse{DocNo:d.DocNo,ArCode:d.ArCode,ArName:d.ArName}, nil
}


