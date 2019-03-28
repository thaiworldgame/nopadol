package mysqldb

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/product"
	//"github.com/mrtomyum/nopadol/auth"
)

type ProductModel struct {
	Id          int          `db:"id"`
	ItemCode    string       `db:"item_code"`
	ItemName    string       `db:"item_name"`
	BarCode     string       `db:"bar_code"`
	UnitCode    string       `db:"unit_code"`
	SalePrice1  float64      `db:"sale_price_1"`
	SalePrice2  float64      `db:"sale_price_2"`
	Rate1       float64      `db:"rate_1"`
	PicPath1    string       `db:"pic_path_1"`
	StkQty      float64      `db:"stk_qty"`
	StockType   int64        `db:"stock_type"`
	AverageCost float64      `db:"average_cost"`
	PickZoneId  string       `db:"pick_zone_id"`
	StkLocation []StockModel `db:stk_location`
	CompanyID   int          `db:"company_id"`
}

type StockModel struct {
	WHCode      string  `db:"wh_code"`
	ShelfCode   string  `db:"shelf_code"`
	Qty         float64 `db:"qty"`
	StkUnitCode string  `db:"stk_unit_code"`
}

type SearchProductStockModel struct {
	Id        int     `db:"id"`
	ItemCode  string  `db:"item_code"`
	WHCode    string  `db:"wh_code"`
	ShelfCode string  `db:"shelf_code"`
	Qty       float64 `db:"qty"`
	UnitCode  string  `db:"unit_code"`
}

type productRepository struct{ db *sqlx.DB }

func NewProductRepository(db *sqlx.DB) product.Repository {
	return &productRepository{db}
}

func (pd *productRepository) SearchByBarcode(req *product.SearchByBarcodeTemplate) (resp interface{}, err error) {
	product := ProductModel{}

	fmt.Println("barcode = ", req.BarCode)

	//sql := `set dateformat dmy     select id,item_code,item_name, bar_code, unit_code, isnull(c.saleprice1,0) as Price, isnull(d.rate,1) as Rate1, isnull(b.picfilename1,'') as PicPath from dbo.bcbarcodemaster a with (nolock) inner join dbo.bcitem b with (nolock) on a.itemcode = b.code left join dbo.bcpricelist c with (nolock) on c.saletype = 0 and c.transporttype = 0 and a.itemcode = c.itemcode and a.unitcode = c.unitcode and cast(rtrim(day(getdate()))+'/'+rtrim(month(getdate()))+'/'+rtrim(year(getdate())) as datetime) >= cast(rtrim(day(startdate))+'/'+rtrim(month(startdate))+'/'+rtrim(year(startdate)) as datetime) and cast(rtrim(day(getdate()))+'/'+rtrim(month(getdate()))+'/'+rtrim(year(getdate())) as datetime) <= cast(rtrim(day(stopdate))+'/'+rtrim(month(stopdate))+'/'+rtrim(year(stopdate)) as datetime)  left join dbo.bcstkpacking d with (nolock) on a.itemcode = d.itemcode and a.unitcode = d.unitcode where a.barcode = ?`
	sql := `select distinct a.id,a.code as item_code,a.item_name,ifnull(a.pic_path1,'') as pic_path_1,b.bar_code,b.unit_code,c.sale_price_1,c.sale_price_2, ifnull(d.rate1,1) as rate_1,ifnull(a.stock_type,0) as stock_type,ifnull((select sum(qty)  as qty from StockLocation where item_code = a.code),0) as stk_qty,ifnull(d.rate1,1)*ifnull(a.average_cost,0) as average_cost from Item a inner join Barcode b on a.code = b.item_code inner join Price c on a.code = c.item_code and b.unit_code = c.unit_code left join ItemRate d on a.code = d.item_code and c.unit_code = d.unit_code where b.bar_code = ? `
	err = pd.db.Get(&product, sql, req.BarCode)
	if err != nil {
		fmt.Println("error = ", err.Error())
		return resp, nil
	}

	pdt_resp := map_product_template(product)

	return map[string]interface{}{
		"id":           pdt_resp.Id,
		"item_code":    pdt_resp.ItemCode,
		"item_name":    pdt_resp.ItemName,
		"barcode":      pdt_resp.BarCode,
		"unit_code":    pdt_resp.UnitCode,
		"sale_price_1": pdt_resp.SalePrice1,
		"sale_price_2": pdt_resp.SalePrice2,
		"rate1":        pdt_resp.Rate1,
		"pic_path_1":   pdt_resp.PicPath1,
		"stk_qty":      pdt_resp.StkQty,
		"stock_type":   pdt_resp.StockType,
		"average_cost": pdt_resp.AverageCost,
	}, nil
	//return pdt_resp, nil
}

func (pd *productRepository) SearchByItemCode(req *product.SearchByItemCodeTemplate) (resp interface{}, err error) {
	products := []ProductModel{}
	//sql := `select id,item_code,item_name,pic_path_1,unit_code,sale_price_1,sale_price_2,rate_1,(qty/rate_1) as qty_unit from (select distinct a.id,a.code as item_code,a.item_name,ifnull(a.pic_path1,'') as pic_path_1,c.unit_code,c.sale_price_1,c.sale_price_2, ifnull(b.rate1,1) as rate_1,ifnull((select sum(qty)  as qty from StockLocation where item_code = a.code),0) as qty from Item a  inner join Price c on a.code = c.item_code left join ItemRate b on a.code = b.item_code and c.unit_code = b.unit_code where a.code = ?) as rs order by unit_code`
	sql := `select distinct a.id,a.code as item_code,a.item_name,ifnull(a.pic_path1,'') as pic_path_1,c.unit_code,c.sale_price_1,c.sale_price_2, ifnull(b.rate1,1) as rate_1,ifnull(a.stock_type,0) as stock_type,ifnull((select sum(qty)  as qty from StockLocation where item_code = a.code),0) as stk_qty,ifnull(b.rate1,1) *ifnull(a.average_cost,0) as average_cost from Item a  inner join Price c on a.code = c.item_code left join ItemRate b on a.code = b.item_code and c.unit_code = b.unit_code where a.code = ? order by unit_code`
	err = pd.db.Select(&products, sql, req.ItemCode)
	if err != nil {
		fmt.Println("error = ", err.Error())
		return resp, nil
	}

	stocks := []StockModel{}
	sql_stk := `select 	distinct ifnull(wh_code,'') as wh_code,ifnull(shelf_code,'') as shelf_code,ifnull(b.qty,0) as qty,ifnull(b.unit_code,'') as stk_unit_code  from Item a left join StockLocation b on a.code = b.item_code  where a.code = ? order by b.wh_code`
	err = pd.db.Select(&stocks, sql_stk, req.ItemCode)
	if err != nil {
		fmt.Println("error = ", err.Error())
		return resp, nil
	}

	prod := []product.ProductTemplate{}

	for _, item := range products {
		pdtline := map_product_template(item)
		for _, sub := range stocks {
			subline := map_itemstock_template(sub)
			pdtline.StkLocation = append(pdtline.StkLocation, subline)
		}
		prod = append(prod, pdtline)
	}

	return prod, nil
}

func (pd *productRepository) SearchByKeyword(req *product.SearchByKeywordTemplate) (resp interface{}, err error) {
	products := []ProductModel{}

	fmt.Println("keyword = ", req.Keyword)

	sql := `select distinct rs.id,rs.code as item_code,rs.item_name,ifnull(b.rate1,1)*ifnull(rs.average_cost,0) as average_cost,ifnull(pic_path1,'') as pic_path_1,'' as bar_code,ifnull(c.sale_price_1,0) as sale_price_1,ifnull(sale_price_2,0) as sale_price_2,ifnull(b.unit_code,'') as unit_code,ifnull(b.rate1,1) as rate_1, ifnull(stock_type,0) as stock_type, ifnull((select sum(qty)  as qty from StockLocation where item_code = rs.code),0) as stk_qty  from (select * from Item where code like concat(?,'%') or item_name like  concat(?,'%') order by code limit 20) as rs left join ItemRate b on rs.code = b.item_code left join Price c on rs.code = c.item_code and b.unit_code = c.unit_code `
	err = pd.db.Select(&products, sql, req.Keyword, req.Keyword)
	if err != nil {
		fmt.Println("error = ", err.Error())
		return resp, nil
	}

	product := []product.ProductTemplate{}
	for _, p := range products {

		pdtline := map_product_template(p)

		stocks := []StockModel{}
		sql_stk := `select 	distinct ifnull(wh_code,'') as wh_code,ifnull(shelf_code,'') as shelf_code,ifnull(b.qty,0) as qty,ifnull(b.unit_code,'') as stk_unit_code  from Item a left join StockLocation b on a.code = b.item_code  where a.code = ? order by b.wh_code`
		err = pd.db.Select(&stocks, sql_stk, p.ItemCode)
		if err != nil {
			fmt.Println("error = ", err.Error())
			return resp, nil
		}

		for _, sub := range stocks {
			subline := map_itemstock_template(sub)
			pdtline.StkLocation = append(pdtline.StkLocation, subline)
		}

		product = append(product, pdtline)
	}

	return product, nil
}

func (pd *productRepository) SearchByItemStockLocation(req *product.SearchByItemCodeTemplate) (resp interface{}, err error) {
	products := []SearchProductStockModel{}

	sql := `select 	distinct a.id,code as item_code,ifnull(wh_code,'') as wh_code,ifnull(shelf_code,'') as shelf_code,ifnull(b.qty,0) as qty,ifnull(b.unit_code,'') as unit_code,ifnull(a.average_cost,0) as average_cost  from 	Item a left join StockLocation b on a.code = b.item_code  where a.code = ? order by b.wh_code`
	err = pd.db.Select(&products, sql, req.ItemCode)
	if err != nil {
		fmt.Println("error = ", err.Error())
		return resp, nil
	}
	product := []product.SearchProductStockTemplate{}
	for _, p := range products {
		pdtline := map_stock_template(p)
		product = append(product, pdtline)

	}

	return product, nil
}

func map_itemstock_template(x StockModel) product.StockTemplate {
	return product.StockTemplate{
		WHCode:      x.WHCode,
		ShelfCode:   x.ShelfCode,
		Qty:         x.Qty,
		StkUnitCode: x.StkUnitCode,
	}
}

func map_stock_template(x SearchProductStockModel) product.SearchProductStockTemplate {
	return product.SearchProductStockTemplate{
		Id:        x.Id,
		ItemCode:  x.ItemCode,
		WHCode:    x.WHCode,
		ShelfCode: x.ShelfCode,
		Qty:       x.Qty,
		UnitCode:  x.UnitCode,
	}
}

func map_product_template(x ProductModel) product.ProductTemplate {
	var stock []product.StockTemplate
	return product.ProductTemplate{
		Id:          x.Id,
		BarCode:     x.BarCode,
		ItemCode:    x.ItemCode,
		ItemName:    x.ItemName,
		SalePrice1:  x.SalePrice1,
		SalePrice2:  x.SalePrice2,
		UnitCode:    x.UnitCode,
		Rate1:       x.Rate1,
		PicPath1:    x.PicPath1,
		StkQty:      x.StkQty,
		StockType:   x.StockType,
		AverageCost: x.AverageCost,
		StkLocation: stock,
	}
}

func (p *ProductModel) SearchByBarcode(db *sqlx.DB, access_token string, bar_code string) {
	u := UserAccess{}
	u.GetProfileByToken(db, access_token)

	fmt.Println("user = ", u.CompanyID, u.BranchID)

	m := Machine{}
	m.SearchMachineNo(db, u.CompanyID, u.BranchID, access_token)
	fmt.Println("machine = ", m.WHCode, bar_code)

	lccommand := `select distinct a.id,a.code as item_code,a.item_name,ifnull(a.pic_path1,'') as pic_path_1,b.bar_code,b.unit_code,c.sale_price_1,c.sale_price_2, ifnull(d.rate1,1) as rate_1,ifnull(a.stock_type,0) as stock_type,ifnull((select sum(qty)  as stk_qty from StockLocation where item_code = a.code),0) as stk_qty,ifnull(d.rate1,1)*ifnull(a.average_cost,0) as average_cost, ifnull(e.zone_id,'A') as pick_zone_id from Item a inner join Barcode b on a.code = b.item_code inner join Price c on a.code = c.item_code and b.unit_code = c.unit_code left join ItemRate d on a.code = d.item_code and c.unit_code = d.unit_code left join item_pick_zone e on a.code = e.item_code and b.unit_code = e.unit_code and e.wh_code = ? where b.bar_code = ?`
	fmt.Println(lccommand, m.WHCode,bar_code)
	rs := db.QueryRow(lccommand, m.WHCode, bar_code)
	rs.Scan(&p.Id, &p.ItemCode, &p.ItemName, &p.PicPath1, &p.BarCode, &p.UnitCode, &p.SalePrice1, &p.SalePrice2, &p.Rate1, &p.StockType, &p.StkQty, &p.AverageCost, &p.PickZoneId)

	fmt.Println("Zone Id =", p.PickZoneId)

	return
}

func (p *productRepository) StoreItem(req *product.ProductNewRequest) (resp interface{}, err error) {

	item := itemModel{}
	err = item.map2itemModel(p.db, req)
	if err != nil {
		fmt.Println("error p.StoreItem map2itemModel ->", err.Error())
		return nil, err
	}

	newItemID, err := item.save(p.db)
	if err != nil {
		fmt.Println("error item.save ", err.Error())
		return nil, err
	}

	//pk := packingRate{}
	//for _, value := range req.PackingRate {
	//
	//	u := itemUnitModel{}
	//	u.id = req.UnitID
	//
	//	u.getByID(p.db)
	//
	//	pk.RatePerBaseUnit = value.RatePerBaseUnit
	//	pk.ItemID = newItemID
	//	pk.ItemCode = req.ItemCode
	//	pk.UnitCode = u.unitCode
	//
	//	_, err = pk.save(p.db)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	//
	//// price insert

	//pr := priceModel{}
	//for _, value := range req.Price {
	//
	//	u := itemUnitModel{}
	//	u.id = value.UnitID
	//	u.getByID(p.db) // bind จาก id
	//
	//	pr.UnitID = req.UnitID

	//	pr.ItemId = value.ItemID

	//	pr.SalePrice1 = value.SalePrice1
	//	pr.SalePrice2 = value.SalePrice2
	//	pr.CompanyID = value.CompanyID
	//	_, err := pr.save(p.db)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//}
	//
	//// insert barcode
	//bar := barcodeModel{}
	//for _, value := range req.Barcode {
	//
	//	u := itemUnitModel{}
	//	u.id = value.UnitID
	//	u.getByID(p.db) // bind จาก id
	//
	//	bar.UnitID = req.UnitID
	//	bar.ItemCode = req.ItemCode
	//	bar.CompanyID = req.CompanyID
	//	bar.BarCode = value.Barcode
	//	_, err := bar.save(p.db)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	return map[string]interface{}{
		"result": "success",
		"new_id": newItemID,
	}, nil

	// todo : insert to Barcode table
}

func (p *productRepository) StoreBarcode(req *product.BarcodeNewRequest) (res interface{}, err error) {

	b := barcodeModel{BarCode: req.Barcode,
		ItemID: req.ItemID,
		ItemCode: req.ItemCode,
		UnitCode: req.UnitCode,
		UnitID: req.UnitID,
		ActiveStatus: req.ActiveStatus,
		CompanyID: req.CompanyID,
	}

	if b.UnitID == 0 {
		u := itemUnitModel{}
		u.unitCode = b.UnitCode
		u.getByCode(p.db)
		b.UnitID = u.id
	}
	newID, err := b.save(p.db)
	if err != nil {
		return nil, err
	}
	return newID, err
}

func (p *productRepository) StorePrice(req *product.PriceTemplate) (interface{}, error) {

	fmt.Println("start store price in mysql package req: ", req)
	item := itemModel{Code: req.ItemCode}

	itemID, err := item.getItemIDbyCode(p.db, req.ItemCode)
	if err != nil {
		return nil, err
	}
	unit := itemUnitModel{unitCode: req.UnitCode}
	//unit.getByID(p.db)
	unit.getByCode(p.db)

	if req.UnitCode == "" {
		return nil, fmt.Errorf("unitcode is empty")
	}

	pr := priceModel{
		ItemId:   itemID,
		ItemCode: req.ItemCode,
		UnitID:   unit.id,
		UnitCode: req.UnitCode,

		SalePrice1: req.SalePrice1,
		SalePrice2: req.SalePrice2,
		SaleType:   req.SaleType,
		CompanyID:  req.CompanyID,
	}

	return pr.save(p.db)
}

func (p *productRepository) StorePackingRate(req *product.PackingRate) (interface{}, error) {
	unit := itemUnitModel{id: req.UnitID}
	unit.getByID(p.db)

	rate := packingRate{
		ItemID:          req.ItemID,
		ItemCode:        req.ItemCode,
		UnitID:          req.UnitID,
		RatePerBaseUnit: req.RatePerBaseUnit,
	}

	return rate.save(p.db)
	//return nil,nil
}
