package sqldb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/product"
)

type ProductModel struct {
	Id       int     `db:"Id"`
	ItemCode string  `db:"ItemCode"`
	ItemName string  `db:"ItemName"`
	BarCode  string  `db:"BarCode"`
	UnitCode string  `db:"UnitCode"`
	Price    float64 `db:"Price"`
	Rate1    float64 `db:"Rate1"`
	PicPath  string  `db:"PicPath"`
}

type productRepository struct{ db *sqlx.DB }

func NewProductRepository(db *sqlx.DB) product.Repository {
	return &productRepository{db}
}

func (pd *productRepository) SearchByBarcode(req *product.SearchByBarcodeTemplate) (resp interface{}, err error) {
	product := ProductModel{}

	sql := `set dateformat dmy     select b.roworder as Id,b.code as ItemCode,b.name1 as ItemName, a.barcode as BarCode, a.unitcode as UnitCode, isnull(c.saleprice1,0) as Price, isnull(d.rate,1) as Rate1, isnull(b.picfilename1,'') as PicPath from dbo.bcbarcodemaster a with (nolock) inner join dbo.bcitem b with (nolock) on a.itemcode = b.code left join dbo.bcpricelist c with (nolock) on c.saletype = 0 and c.transporttype = 0 and a.itemcode = c.itemcode and a.unitcode = c.unitcode and cast(rtrim(day(getdate()))+'/'+rtrim(month(getdate()))+'/'+rtrim(year(getdate())) as datetime) >= cast(rtrim(day(startdate))+'/'+rtrim(month(startdate))+'/'+rtrim(year(startdate)) as datetime) and cast(rtrim(day(getdate()))+'/'+rtrim(month(getdate()))+'/'+rtrim(year(getdate())) as datetime) <= cast(rtrim(day(stopdate))+'/'+rtrim(month(stopdate))+'/'+rtrim(year(stopdate)) as datetime)  left join dbo.bcstkpacking d with (nolock) on a.itemcode = d.itemcode and a.unitcode = d.unitcode where a.barcode = ?`
	err = pd.db.Get(&product, sql, req.BarCode)
	if err != nil {
		fmt.Println("error = ", err.Error())
		return resp, nil
	}

	pdt_resp := map_product_template(product)

	return map[string]interface{}{
		"id":        pdt_resp.Id,
		"item_code": pdt_resp.ItemCode,
		"item_name": pdt_resp.ItemName,
		"barcode":   pdt_resp.BarCode,
		"unit_code": pdt_resp.UnitCode,
		"price":     pdt_resp.Price,
		"rate1":     pdt_resp.Rate1,
		"pic_path":  pdt_resp.PicPath,
	}, nil
	//return pdt_resp, nil
}

func (pd *productRepository) SearchByKeyword(req *product.SearchByKeywordTemplate) (resp interface{}, err error) {
	products := []ProductModel{}

	sql := `set dateformat dmy     select b.roworder as Id,b.code as ItemCode,b.name1 as ItemName, a.barcode as BarCode, a.unitcode as UnitCode, isnull(c.saleprice1,0) as Price, isnull(d.rate,1) as Rate1, isnull(b.picfilename1,'') as PicPath from dbo.bcbarcodemaster a with (nolock) inner join dbo.bcitem b with (nolock) on a.itemcode = b.code left join dbo.bcpricelist c with (nolock) on c.saletype = 0 and c.transporttype = 0 and a.itemcode = c.itemcode and a.unitcode = c.unitcode and cast(rtrim(day(getdate()))+'/'+rtrim(month(getdate()))+'/'+rtrim(year(getdate())) as datetime) >= cast(rtrim(day(startdate))+'/'+rtrim(month(startdate))+'/'+rtrim(year(startdate)) as datetime) and cast(rtrim(day(getdate()))+'/'+rtrim(month(getdate()))+'/'+rtrim(year(getdate())) as datetime) <= cast(rtrim(day(stopdate))+'/'+rtrim(month(stopdate))+'/'+rtrim(year(stopdate)) as datetime)  left join dbo.bcstkpacking d with (nolock) on a.itemcode = d.itemcode and a.unitcode = d.unitcode where (a.barcode like '%'+?+'%' or b.code like '%'+?+'%' or b.name1 like '%'+?+'%')`
	err = pd.db.Select(&products, sql, req.Keyword, req.Keyword, req.Keyword)
	if err != nil {
		fmt.Println("error = ", err.Error())
		return resp, nil
	}

	product := []product.ProductTemplate{}
	for _, p := range products {
		pdtline := map_product_template(p)
		product = append(product, pdtline)

	}

	return product, nil
}

func map_product_template(x ProductModel) product.ProductTemplate {
	return product.ProductTemplate{
		Id:       x.Id,
		BarCode:  x.BarCode,
		ItemCode: x.ItemCode,
		ItemName: x.ItemName,
		Price:    x.Price,
		UnitCode: x.UnitCode,
		Rate1:    x.Rate1,
		PicPath:  x.PicPath,
	}
}
