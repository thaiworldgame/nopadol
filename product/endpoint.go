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
		Id          int     `json:"id"`
		ItemCode    string  `json:"item_code"`
		ItemName    string  `json:"item_name"`
		BarCode     string  `json:"bar_code"`
		UnitCode    string  `json:"unit_code"`
		SalePrice1  float64 `json:"sale_price_1"`
		SalePrice2  float64 `json:"sale_price_2"`
		Rate1       float64 `json:"rate_1"`
		PicPath1    string  `json:"pic_path_1"`
		StkQty      float64 `json:"stk_qty"`
		StockType   int64   `json:"stock_type"`
		AverageCost float64 `json:"average_cost"`
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

func MakeNewProduct(s Service) interface{} {
	type request_price struct {
		UnitID     int64   `json:"unit_id"`
		SalePrice1 float64 `json:"sale_price_1"`
		SalePrice2 float64 `json:"sale_price_2"`
		SaleType   int     `json:"sale_type"`
	}
	type request_barcode struct {
		Barcode string `json:"barcode"`
		UnitID  int64  `json:"unit_id"`
	}

	type request_packingrate struct {
		UnitID          int64 `json:"unit_id"`
		RatePerBaseUnit int `json:"rate_per_base_unit"`
	}
	type request struct {
		Code        string                `json:"code"`
		Name        string                `json:"name"`
		UnitID      int64                 `json:"unit_code"`
		Picture     string                `json:"picture"`
		StockType   int                   `json:"stock_type"`
		Price       []request_price       `json:"price"`
		Barcode     []request_barcode     `json:"barcode"`
		PackingRate []request_packingrate `json:"packing_rate"`
	}
	return func(ctx context.Context, req *request) (interface{}, error) {

		var barcodes []BarcodeTemplate
		// bind barcode template
		for _, value := range req.Barcode {
			bct := BarcodeTemplate{}
			bct.Barcode = value.Barcode
			bct.UnitID = value.UnitID
			barcodes = append(barcodes, bct)
		}

		// bind price template
		var prices []PriceTemplate
		for _, value := range req.Price {
			pr := PriceTemplate{}
			pr.SalePrice1 = value.SalePrice1
			pr.SalePrice2 = value.SalePrice2
			pr.UnitID = value.UnitID
			pr.SaleType = value.SaleType
			prices = append(prices, pr)
		}

		var Rates []PackingRate
		for _, value := range req.PackingRate {
			rt := PackingRate{}
			rt.RatePerBaseUnit = value.RatePerBaseUnit
			rt.UnitID = value.UnitID
		}

		resp, err := s.StoreItem(&ProductNewRequest{
			ItemCode:    req.Code,
			ItemName:    req.Name,
			UnitID:      req.UnitID,
			StockType:   req.StockType,
			Picture:     req.Picture,
			Barcode:     barcodes,
			Price:       prices,
			PackingRate: Rates,
		})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"result": resp,
		}, nil
	}
}
