package sync

import (
	"context"
	"fmt"
	"github.com/mrtomyum/nopadol/product"
)

func makeCreateProductEndpoint(s Service) interface{} {
	type itemUpdateRequest struct {
		product.ProductTemplate
	}
	return func(ctx context.Context, req *itemUpdateRequest) (interface{}, error) {
		resp, err := s.ProductUpdate(req)
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}
