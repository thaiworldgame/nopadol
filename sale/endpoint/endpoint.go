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

func buildsaleorderRequest(x *sale.NewSaleOrderRequest) sale.SaleOrderTemplate{
	var subs []sale.SubsTemplate
	return sale.SaleOrderTemplate{
		DocNo:x.DocNo,
		DocDate:x.DocDate,
		ArCode:x.ArCode,
		ArName:x.ArName,
		Subs:subs,
	}
}

func buildsaleordersubRequest(x *sale.SubsSaleOrderResponse) sale.SubsTemplate {
	return sale.SubsTemplate{
		ItemCode:x.ItemCode,
		ItemName:x.ItemName,
		Qty:x.Qty,
		UnitCode:x.UnitCode,
	}
}

func (ep *endpoint) NewSaleOrder(ctx context.Context, req *sale.NewSaleOrderRequest) (*sale.NewSaleOrderResponse, error) {
	fmt.Println("endpoint docno=",req)

	Resp := buildsaleorderRequest(req)

	for _, v := range req.Subs {
		//fmt.Println(v)
		soline := buildsaleordersubRequest(v)
		Resp.Subs = append(Resp.Subs,soline)

		//fmt.Println("So line = ",soline)
	}


	id, err := ep.s.NewSaleOrder(ctx, &sale.SaleOrderTemplate{DocNo:req.DocNo,DocDate:req.DocDate,ArCode:req.ArCode,ArName:req.ArName,Subs:req.Subs})

	if err != nil {
		return nil,err
	}

	return &sale.NewSaleOrderResponse{Id:id},nil
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


func buildsaleorder(x sale.SaleOrderTemplate) sale.SearchSaleOrderResponse{
	var subs []sale.SubsSaleOrderResponse
	return sale.SearchSaleOrderResponse{
		DocNo:x.DocNo,
		DocDate:x.DocDate,
		ArCode:x.ArCode,
		ArName:x.ArName,
		Subs:subs,
	}
}

func buildsaleordersub(x sale.SubsTemplate) sale.SubsSaleOrderResponse {
	return sale.SubsSaleOrderResponse{
		ItemCode:x.ItemCode,
		ItemName:x.ItemName,
		Qty:x.Qty,
		UnitCode:x.UnitCode,
	}
}

func (ep *endpoint) Search(ctx context.Context, req *sale.SearchSaleOrderRequest) (*sale.SearchSaleOrderResponse, error) {
//func (ep *endpoint) Search(ctx context.Context, req *sale.SearchSaleOrderRequest) (interface{}, error) {

	fmt.Println("keyword endpoint = ",req.Keyword)

	sale_order, err := ep.s.Search(ctx, &sale.EntitySearch{
		Keyword:req.Keyword,
	})
	//fmt.Println("saleorder = ",sale_order.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}


	Resp := buildsaleorder(sale_order)

	for _, v := range sale_order.Subs {
		//fmt.Println(v)
		soline := buildsaleordersub(v)
		Resp.Subs = append(Resp.Subs,soline)

		//fmt.Println("So line = ",soline)
	}


	fmt.Println("DocNo =", sale_order.DocNo)
	fmt.Println("DocDate =", sale_order.DocDate)
	fmt.Println("ArCode =", sale_order.ArCode)
	fmt.Println("so =",sale_order)

	fmt.Println("Search By = ",sale.EntitySearch{}.Keyword)

	//}, nil
	return &sale.SearchSaleOrderResponse{
		DocNo: Resp.DocNo, DocDate: Resp.DocDate, ArCode: Resp.ArCode, ArName: Resp.ArName,Subs:Resp.Subs,
	},nil
	//return sale_order,nil
}
