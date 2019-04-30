package endpoint

//import (
//	"context"
//	"github.com/mrtomyum/nopadol/product"
//	"fmt"
//)
//
//func New(s product.Service) product.Endpoint {
//	return &endpoint{s}
//}
//
//type endpoint struct {
//	s product.Service
//}
//
//func (ep *endpoint) SearchProductByBarCode(ctx context.Context, req *product.SearchByBarcodeRequest) (*product.SearchProductResponse, error) {
//	product, err := ep.s.SearchProductByBarcode(ctx, &product.SearchByBarcodeTemplate{
//		BarCode: req.BarCode,
//	})
//	if err != nil {
//		fmt.Println("error = ", err.Error())
//		return nil, err
//	}
//
//	resp := map_product_response(product)
//
//	return &resp, nil
//}
//
//func map_product_response(x product.ProductTemplate) product.SearchProductResponse {
//	return product.SearchProductResponse{
//		Id:       x.Id,
//		BarCode:  x.BarCode,
//		ItemCode: x.ItemCode,
//		ItemName: x.ItemName,
//		Price:    x.Price,
//		UnitCode: x.UnitCode,
//		Rate1:    x.Rate1,
//		PicPath:  x.PicPath,
//	}
//}
