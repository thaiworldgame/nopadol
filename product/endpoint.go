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

	SearchByKeywordRequest struct {
		Keyword string `json:"keyword"`
	}

	SearchProductResponse struct {
		Id         int     `json:"id"`
		ItemCode   string  `json:"item_code"`
		ItemName   string  `json:"item_name"`
		BarCode    string  `json:"bar_code"`
		UnitCode   string  `json:"unit_code"`
		SalePrice1 float64 `json:"sale_price_1"`
		SalePrice2 float64 `json:"sale_price_2"`
		Rate1      float64 `json:"rate_1"`
		PicPath1   string  `json:"pic_path_1"`
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
