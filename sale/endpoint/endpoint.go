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


//Sale Order ///////////////////////////////////////////////////////////////////////////////////////////////////////////

func (ep *endpoint) Search(ctx context.Context, req *sale.SearchSaleOrderRequest) (*sale.SearchSaleOrderResponse, error) {

	fmt.Println("keyword endpoint = ",req.Keyword)

	sale_order, err := ep.s.Search(ctx, &sale.EntitySearch{
		Keyword:req.Keyword,
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}

	Resp := buildsaleorder(sale_order)

	for _, v := range sale_order.Subs {
		soline := buildsaleordersub(v)
		Resp.Subs = append(Resp.Subs,soline)
	}

	return &sale.SearchSaleOrderResponse{
		DocNo: Resp.DocNo, DocDate: Resp.DocDate, ArCode: Resp.ArCode, ArName: Resp.ArName,Subs:Resp.Subs,
	},nil
}

func (ep *endpoint) NewSaleOrder(ctx context.Context, req sale.NewSaleOrderRequest) (*sale.NewSaleOrderResponse, error) {
	fmt.Println("endpoint docno=",req)

	Resp := buildsaleorderRequest(req)

	for _, v := range req.Subs{
		soline := buildsaleordersubRequest(v)
		Resp.Subs = append(Resp.Subs,soline)
	}

	id, err := ep.s.NewSaleOrder(ctx, &sale.SaleOrderTemplate{DocNo:req.DocNo,DocDate:req.DocDate,ArCode:req.ArCode,ArName:req.ArName,Subs:Resp.Subs})

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

func buildsaleorderRequest(x sale.NewSaleOrderRequest) sale.SaleOrderTemplate{
	var subs []sale.SubsTemplate
	return sale.SaleOrderTemplate{
		DocNo:x.DocNo,
		DocDate:x.DocDate,
		ArCode:x.ArCode,
		ArName:x.ArName,
		Subs:subs,
	}
}

func buildsaleordersubRequest(x sale.NewSubsSaleOrderRequest) sale.SubsTemplate {
	return sale.SubsTemplate{
		ItemCode:x.ItemCode,
		ItemName:x.ItemName,
		Qty:x.Qty,
		UnitCode:x.UnitCode,
	}
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

//Pos //////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//func(ep *endpoint) NewPos(ctx context.Context, req sale.NewPosRequest)(*sale.NewResponse, error){
//	err := ep.s.NewPos(ctx, &sale.NewPosRequest{})
//	if err != nil {
//		return nil, nil
//	}
//	return &sale.NewResponse{Id:0},nil
//}