package product

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
	Picture     string            `json:"picture"`
	StockType   int               `json:"stock_type"`
	Barcode     []BarcodeTemplate `json:"barcode"`
	Price       []PriceTemplate   `json:"price"`
	PackingRate []PackingRate     `json:"packing_rate"`
}

type PackingRate struct {
	UnitID          int64 `json:"unit_id"`
	RatePerBaseUnit int `json:"rate_per_base_unit"`
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
}

type BarcodeTemplate struct {
	Id      int64  `json:"id"`
	ItemID  int64  `json:"item_id"`
	Barcode string `json:"barcode"`
	UnitID  int64  `json:"unit_id"`
}

type PriceTemplate struct {
	Id         int64   `json:"id"`
	ItemID     int64   `json:"item_id"`
	UnitID     int64   `json:"unit_id"`
	SalePrice1 float64 `json:"sale_price_1"`
	SalePrice2 float64 `json:"sale_price_2"`
	SaleType   int     `json:"sale_type"`
}

//type Item struct {
//	Parent *Category
//	Name   string
//	UOM    *Unit
//}
//
//type Unit struct {
//	Name  string
//	Ratio float32
//}
//
//func (i *Item) MoveCat(newCat *Category) {
//	i.Parent = newCat
//}
//
//func (i *Item) Stock(s []*Stock) []*Stock {
//	r := make([]*Stock, 0)
//	for _, v := range s {
//		if v.item == i {
//			r = append(r, v)
//		}
//	}
//	//	fmt.Printf("r=", r)
//	return r
//}
//
////func (i *Item) Calc(stocks []*Stock, tx []*Trans) {
////	for _, t := range tx {
////		if t.item == i {
////			// ถ้า Trans ไม่มี Location ใน Stock ให้เพิ่มรายการใน Stock ก่อน
////			for k, s := range stocks {
////				if s.item == i {
////					if s.loc == t.locIn {
////						s.bal += t.qty
////						fmt.Printf("Item %v Loc: %v qty: %v bal: %v\n", s.item.Name, s.loc.Code, t.qty, s.bal)
////					} else {
////
////					}
////				} else {
////					stocks = append(stocks, &Stock{item: i, loc: t.locIn, bal: t.qty})
////					fmt.Println(">>Append Stock from Loc In:", stocks[k].item.Name, stocks[k].loc.Code, stocks[k].bal)
////				}
////				if s.item == i && s.loc == t.locOut {
////					s.bal -= t.qty
////					fmt.Printf("Item %v Loc: %v qty:-%v bal: %v\n", s.item.Name, s.loc.Code, t.qty, s.bal)
////				} else {
////					stocks = append(stocks, &Stock{item: i, loc: t.locOut, bal: t.qty})
////					fmt.Println(">>Append Stock from Loc Out:", stocks[k].item.Name, stocks[k].loc.Code, stocks[k].bal)
////				}
////			}
////		}
////	}
////}
//
//func (i *Item) Calc(stocks []*Stock, tx []*Tran) []*Stock {
//	for _, t := range tx {
//		if t.item == i {
//			countIn := 0
//			countOut := 0
//			for _, s := range stocks {
//				if s.loc == t.locIn {
//					countIn++
//				} else if s.loc == t.locOut {
//					countOut++
//				}
//			}
//			if countIn == 0 {
//				stocks = append(stocks, &Stock{item: i, loc: t.locIn, bal: t.qty})
//			}
//			if countOut == 0 {
//				stocks = append(stocks, &Stock{item: i, loc: t.locOut, bal: -t.qty})
//			}
//		}
//	}
//	return stocks
//}
//
//// ==== Category ==== //
//type Category struct {
//	Parent *Category
//	Name   string
//}
//
//func (c *Category) New() {
//	c.Name = "New Cat"
//}
//
//func FindByName(cats []*Category, n string) *Category {
//	for i, c := range cats {
//		if strings.Contains(c.Name, n) {
//			return cats[i]
//		}
//	}
//	return nil
//}
//
//func (c *Category) MoveCat(newCat *Category) {
//	// ต้องเพิ่มการป้องกันการอ้าง Child มาเป็น Parent ของโหนด
//	// โดยจะต้องทวนสอบย้อนกลับว่าโหนดแม่ที่ย้ายมาจะไม่อยู่ใต้โหนดของลูกตัวเอง
//	c.Parent = newCat
//}
//
//// ==== Location ==== //
//type LocType int
//
//const (
//	ROOT LocType = iota
//	BUY
//	STORE
//	VEHICLE
//	MACHINE
//	COLUMN
//	SALE
//)
//
//type Location struct {
//	Parent *Location
//	LocType
//	Code string
//}
//
//// ==== Stock ==== //
//type Stock struct {
//	item  *Item
//	loc   *Location
//	bal   int64
//	digit int8 // ขนาดของทศนิยมที่จะใช้กับสินค้านี้
//}
//
//func (s *Stock) Move(toLoc *Location) {
//	s.loc = toLoc
//}
//
//func (s *Stock) Calc(trans []*Tran) error {
//	for _, t := range trans {
//		if t.item == s.item {
//			if t.locOut == s.loc {
//				s.bal -= t.qty
//			} else if t.locIn == s.loc {
//				s.bal += t.qty
//			} else {
//				return errors.New("No Item Stock in Transaction")
//			}
//		}
//	}
//	return nil
//}
//
//// ==== Transaction ==== //
//type Tran struct {
//	ID     uint64
//	item   *Item
//	locOut *Location
//	locIn  *Location
//	qty    int64
//}
//
//type Trans []Tran
//
//func (slice Trans) Len() int {
//	return len(slice)
//}
//
//func (slice Trans) Less(i, j int) bool {
//	return slice[i].ID < slice[j].ID
//}
//
//func (slice Trans) Swap(i, j int) {
//	slice[i], slice[j] = slice[j], slice[i]
//}
