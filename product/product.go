package product

import "time"

type SearchByBarcodeTemplate struct {
	BarCode string `json:"bar_code"`
}

type SearchByItemCodeTemplate struct {
	ItemCode string `json:"item_code"`
}

type SearchByKeywordTemplate struct {
	Keyword string `json:"keyword"`
}

type ProductTemplate struct {
	Id          int             `json:"id"`
	ItemCode    string          `json:"item_code"`
	ItemName    string          `json:"item_name"`
	BarCode     string          `json:"bar_code"`
	UnitCode    string          `json:"unit_code"`
	SalePrice1  float64         `json:"sale_price_1"`
	SalePrice2  float64         `json:"sale_price_2"`
	Rate1       float64         `json:"rate_1"`
	PicPath1    string          `json:"pic_path_1"`
	StkQty      float64         `json:"stk_qty"`
	StockType   int64           `json:"stock_type"`
	AverageCost float64         `json:"average_cost"`
	StkLocation []StockTemplate `json:"stk_location"`
}

type ProductNewRequest struct {
	ItemCode    string            `json:"item_code"`
	ItemName    string            `json:"item_name"`
	UnitID      int64             `json:"unit_id"`
	UnitCode    string            `json:"unit_code"`
	Picture     string            `json:"picture"`
	StockType   int               `json:"stock_type"`
	Barcode     []BarcodeTemplate `json:"barcode"`
	Price       []PriceTemplate   `json:"price"`
	PackingRate []PackingRate     `json:"packing_rate"`
	CompanyID   int               `json:"company_id"`
	CreateBy    string            `json:"create_by"`
	CreateTime  time.Time         `json:"create_time"`
	EditBy      string            `json:"edit_by"`
	EditTime    time.Time         `json:"edit_time"`
}

type BarcodeNewRequest struct {
	temID        int64  `json:"item_id"`
	ItemCode     string `json:"item_code"`
	Barcode      string `json:"barcode"`
	UnitCode     string `json:"unit_code"`
	UnitID       int64  `json:"unit_id"`
	ActiveStatus int    `json:"active_status"`
}
type PackingRate struct {
	UnitID          int64 `json:"unit_id"`
	RatePerBaseUnit int   `json:"rate_per_base_unit"`
}

type StockTemplate struct {
	WHCode      string  `json:"wh_code"`
	ShelfCode   string  `json:"shelf_code"`
	Qty         float64 `json:"qty"`
	StkUnitCode string  `json:"Stk_unit_code"`
}

type SearchProductStockTemplate struct {
	Id        int     `json:"id"`
	ItemCode  string  `json:"item_code"`
	WHCode    string  `json:"wh_code"`
	ShelfCode string  `json:"shelf_code"`
	Qty       float64 `json:"qty"`
	UnitCode  string  `json:"unit_code"`
	CompanyID int     `json:"company_id"`
}

type BarcodeTemplate struct {
	Id        int64  `json:"id"`
	ItemID    int64  `json:"item_id"`
	Barcode   string `json:"barcode"`
	UnitID    int64  `json:"unit_id"`
	CompanyID int    `json:"company_id"`
}

type PriceTemplate struct {
	Id         int64   `json:"id"`
	ItemID     int64   `json:"item_id"`
	UnitID     int64   `json:"unit_id"`
	SalePrice1 float64 `json:"sale_price_1"`
	SalePrice2 float64 `json:"sale_price_2"`
	SaleType   int     `json:"sale_type"`
	CompanyID  int     `json:"company_id"`
}
