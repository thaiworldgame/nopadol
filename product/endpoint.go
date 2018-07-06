package product

import "context"

type Endpoint interface {
	SearchProductByBarCode(context.Context, *SearchByBarcodeRequest) (*SearchProductResponse, error)
}

type (
	SearchByBarcodeRequest struct {
		BarCode string `json:"bar_code"`
	}

	SearchProductResponse struct {
		Id       int     `json:"id"`
		ItemCode string  `json:"item_code"`
		ItemName string  `json:"item_name"`
		BarCode  string  `json:"bar_code"`
		UnitCode string  `json:"unit_code"`
		Price    float64 `json:"price"`
		Rate1    float64 `json:"rate_1"`
		PicPath  string  `json:"pic_path"`
	}
)
