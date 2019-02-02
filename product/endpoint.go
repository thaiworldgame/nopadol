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
	// unitID กรณี รับจากระบบ erp บน Web รับเป็น ID ก็ได้
	// unitID กรณี รับจากการโอนจาก BC ให้รับเป็น UnitCode แทน
	// unitID กรณี ส่งมาเป็น UnitID หรือ UnitCode อย่างใดอย่างนึง ต้อง Bind ค่า ให้ครบก่อนน่าจะดี
	type requestPrice struct {
		UnitID     int64   `json:"unit_id"`
		UnitCode   string  `json:"unit_code"`
		SalePrice1 float64 `json:"sale_price_1"`
		SalePrice2 float64 `json:"sale_price_2"`
		SaleType   int     `json:"sale_type"`
	}
	type requestBarcode struct {
		Barcode  string `json:"barcode"`
		UnitCode string `json:"unit_code"`
		UnitID   int64  `json:"unit_id"`
	}

	type RequestPacking struct {
		UnitID          int64  `json:"unit_id"`
		UnitCode        string `json:"unit_code"`
		RatePerBaseUnit int    `json:"rate_per_base_unit"`
	}
	type requestNewItem struct {
		Code        string           `json:"code"`
		Name        string           `json:"name"`
		UnitCode    string           `json:"unit_code"`
		UnitID      int64            `json:"unit_id"`
		Picture     string           `json:"picture"`
		StockType   int              `json:"stock_type"`
		Price       []requestPrice   `json:"price"`
		Barcode     []requestBarcode `json:"barcode"`
		PackingRate []RequestPacking `json:"packing_rate"`
		CompanyID   int              `json:"company_id"`
	}
	return func(ctx context.Context, req *requestNewItem) (interface{}, error) {

		var barcodes []BarcodeTemplate
		var prices []PriceTemplate
		var Rates []PackingRate

		// bind barcode template
		if len(req.Barcode) > 0 {
			fmt.Println("bind req barcocde")
			for _, value := range req.Barcode {
				bct := BarcodeTemplate{}
				bct.Barcode = value.Barcode
				bct.UnitID = value.UnitID
				barcodes = append(barcodes, bct)
			}
		}

		if len(req.Price) > 0 {
			fmt.Println("bind req price value : ", req.Price)

			// bind price template
			for _, value := range req.Price {
				pr := PriceTemplate{}
				pr.SalePrice1 = value.SalePrice1
				pr.SalePrice2 = value.SalePrice2
				pr.UnitID = value.UnitID
				pr.SaleType = value.SaleType
				pr.CompanyID = req.CompanyID
				prices = append(prices, pr)
			}
		}

		if len(req.PackingRate) > 0 {

			fmt.Println("bind req Rate")
			for _, value := range req.PackingRate {
				rt := PackingRate{}
				rt.RatePerBaseUnit = value.RatePerBaseUnit
				rt.UnitID = value.UnitID
			}
		}

		fmt.Println("request data ->", req)
		resp, err := s.StoreItem(&ProductNewRequest{
			ItemCode:    req.Code,
			ItemName:    req.Name,
			UnitID:      req.UnitID,
			StockType:   req.StockType,
			Picture:     req.Picture,
			Barcode:     barcodes,
			Price:       prices,
			PackingRate: Rates,
			CompanyID:   req.CompanyID,
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


func MakeNewBarcode(s Service) interface{}
