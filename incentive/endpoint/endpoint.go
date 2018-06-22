package endpoint

import (
	"context"
	"github.com/mrtomyum/nopadol/incentive"
	"fmt"
)

// New creates new domain1 endpoint
func New(s incentive.Service) incentive.Endpoint {
	return &endpoint{s}
}

type endpoint struct {
	s incentive.Service
}

func (ep *endpoint) SearchSaleCode(ctx context.Context, req *incentive.SearchSaleCodeRequest) (*incentive.SearchSaleCodeResponse, error) {
	
	fmt.Println("keyword endpoint = ",req.Keyword)

	sale_code, err := ep.s.SearchSaleCode(ctx, &incentive.EntitySearch{
		Keyword:req.Keyword,
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	fmt.Println("Search By = ",incentive.EntitySearch{}.Keyword)
	return &incentive.SearchSaleCodeResponse{
		Incentive: sale_code,
	}, nil
}
