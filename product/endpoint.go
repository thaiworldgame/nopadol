package product

import (
	"context"
	"fmt"
)

//type Endpoint interface {
//	SearchProductByBarCode(context.Context, *SearchByBarcodeRequest) (*SearchProductResponse, error)
//}

type (
	SearchByBarcodeRequest struct {
		BarCode string `json:"bar_code"`
	}

	SearchByItemCodeRequest struct {
		ItemCode string `json:"item_code"`
	}

	SearchByKeywordRequest struct {
		Keyword string `json:"keyword"`
	}

	SearchProductResponse struct {
		Id            int     `json:"id"`
		ItemCode      string  `json:"item_code"`
		ItemName      string  `json:"item_name"`
		BarCode       string  `json:"bar_code"`
		UnitCode      string  `json:"unit_code"`
		SalePrice1    float64 `json:"sale_price_1"`
		SalePrice2    float64 `json:"sale_price_2"`
		Rate1         float64 `json:"rate_1"`
		PicPath1      string  `json:"pic_path_1"`
		StkQty      float64 `json:"stk_qty"`
		StkLocation []Stock `json:"stk_location"`
	}

	Stock struct {
		WHCode      string  `json:"wh_code"`
		ShelfCode   string  `json:"shelf_code"`
		Qty         float64 `json:"qty"`
		StkUnitCode string  `json:"Stk_unit_code"`
	}

	SearchProductStock struct {
		Id        int     `json:"id"`
		ItemCode  string  `json:"item_code"`
		WHCode    string  `json:"wh_code"`
		ShelfCode string  `json:"shelf_code"`
		Qty       float64 `json:"qty"`
		UnitCode  string  `json:"unit_code"`
	}
)

func SearchByBarcode(s Service) interface{} {
	return func(ctx context.Context, req *SearchByBarcodeRequest) (interface{}, error) {
		resp, err := s.SearchByBarcode(&SearchByBarcodeTemplate{BarCode: req.BarCode})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func SearchByItemCode(s Service) interface{} {
	return func(ctx context.Context, req *SearchByItemCodeRequest) (interface{}, error) {
		resp, err := s.SearchByItemCode(&SearchByItemCodeTemplate{ItemCode: req.ItemCode})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func SearchByItemStockLocation(s Service) interface{} {
	return func(ctx context.Context, req *SearchByItemCodeRequest) (interface{}, error) {
		resp, err := s.SearchByItemStockLocation(&SearchByItemCodeTemplate{ItemCode: req.ItemCode})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func SearchByKeyword(s Service) interface{} {
	return func(ctx context.Context, req *SearchByKeywordRequest) (interface{}, error) {
		resp, err := s.SearchByKeyword(&SearchByKeywordTemplate{Keyword: req.Keyword})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}
