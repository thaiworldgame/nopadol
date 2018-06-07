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

func (ep *endpoint) NewSO(ctx context.Context, req *sale.NewSORequest) (*sale.NewSOResponse, error) {
	id, err := ep.s.NewSO(ctx, &sale.SaleOrder{})
	if err != nil {
		return nil,err
	}

	return &sale.NewSOResponse{Id:id},nil
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

	sale_order, err := ep.s.Search(ctx, &sale.EntitySearch{
		Keyword:req.Keyword,
	})
	//fmt.Println("saleorder = ",sale_order.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}

	fmt.Println("Search By = ",sale.EntitySearch{}.Keyword)



	return &sale.SearchSaleResponse{
		Sale: sale_order,
		//DocNo:so.DocNo,
		//ArCode:so.ArCode,
		//ArName:so.ArName,
		//Subs:so.Subs,//{so.Subs.ItemCode,so.Subs.ItemName,so.Subs.Qty}
	}, nil
}
