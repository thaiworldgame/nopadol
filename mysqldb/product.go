package mysqldb

import (
	"github.com/mrtomyum/nopadol/product"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ProductModel struct {
	Id         int     `db:"id"`
	ItemCode   string  `db:"item_code"`
	ItemName   string  `db:"item_name"`
	BarCode    string  `db:"bar_code"`
	UnitCode   string  `db:"unit_code"`
	SalePrice1 float64 `db:"sale_price_1"`
	SalePrice2 float64 `db:"sale_price_2"`
	Rate1      float64 `db:"rate_1"`
	PicPath1   string  `db:"pic_path_1"`
}

type productRepository struct{ db *sqlx.DB }

func NewProductRepository(db *sqlx.DB) product.Repository {
	return &productRepository{db}
}

func (pd *productRepository) SearchByBarcode(req *product.SearchByBarcodeTemplate) (resp interface{}, err error) {
	product := ProductModel{}

	//sql := `set dateformat dmy     select id,item_code,item_name, bar_code, unit_code, isnull(c.saleprice1,0) as Price, isnull(d.rate,1) as Rate1, isnull(b.picfilename1,'') as PicPath from dbo.bcbarcodemaster a with (nolock) inner join dbo.bcitem b with (nolock) on a.itemcode = b.code left join dbo.bcpricelist c with (nolock) on c.saletype = 0 and c.transporttype = 0 and a.itemcode = c.itemcode and a.unitcode = c.unitcode and cast(rtrim(day(getdate()))+'/'+rtrim(month(getdate()))+'/'+rtrim(year(getdate())) as datetime) >= cast(rtrim(day(startdate))+'/'+rtrim(month(startdate))+'/'+rtrim(year(startdate)) as datetime) and cast(rtrim(day(getdate()))+'/'+rtrim(month(getdate()))+'/'+rtrim(year(getdate())) as datetime) <= cast(rtrim(day(stopdate))+'/'+rtrim(month(stopdate))+'/'+rtrim(year(stopdate)) as datetime)  left join dbo.bcstkpacking d with (nolock) on a.itemcode = d.itemcode and a.unitcode = d.unitcode where a.barcode = ?`
	sql := `select distinct a.id,a.code as item_code,a.item_name,ifnull(a.pic_path1,'') as pic_path_1,b.bar_code,b.unit_code,c.sale_price_1,c.sale_price_2, 1 as rate_1 from Item a inner join Barcode b on a.code = b.item_code inner join Price c on a.code = c.item_code and b.unit_code = c.unit_code where b.bar_code = ?`
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
	}, nil
	//return pdt_resp, nil
}

func (pd *productRepository) SearchByKeyword(req *product.SearchByKeywordTemplate) (resp interface{}, err error) {
	products := []ProductModel{}
	fmt.Println("Product")
	//sql := `set dateformat dmy     select b.roworder as Id,b.code as ItemCode,b.name1 as ItemName, a.barcode as BarCode, a.unitcode as UnitCode, isnull(c.saleprice1,0) as Price, isnull(d.rate,1) as Rate1, isnull(b.picfilename1,'') as PicPath from dbo.bcbarcodemaster a with (nolock) inner join dbo.bcitem b with (nolock) on a.itemcode = b.code left join dbo.bcpricelist c with (nolock) on c.saletype = 0 and c.transporttype = 0 and a.itemcode = c.itemcode and a.unitcode = c.unitcode and cast(rtrim(day(getdate()))+'/'+rtrim(month(getdate()))+'/'+rtrim(year(getdate())) as datetime) >= cast(rtrim(day(startdate))+'/'+rtrim(month(startdate))+'/'+rtrim(year(startdate)) as datetime) and cast(rtrim(day(getdate()))+'/'+rtrim(month(getdate()))+'/'+rtrim(year(getdate())) as datetime) <= cast(rtrim(day(stopdate))+'/'+rtrim(month(stopdate))+'/'+rtrim(year(stopdate)) as datetime)  left join dbo.bcstkpacking d with (nolock) on a.itemcode = d.itemcode and a.unitcode = d.unitcode where (a.barcode like '%'+?+'%' or b.code like '%'+?+'%' or b.name1 like '%'+?+'%')`
	sql := `select distinct a.id,a.code as item_code,a.item_name,ifnull(a.pic_path1,'') as pic_path_1,b.bar_code,b.unit_code,c.sale_price_1,c.sale_price_2, 1 as rate_1 from Item a inner join Barcode b on a.code = b.item_code inner join Price c on a.code = c.item_code and b.unit_code = c.unit_code where (b.bar_code like concat('%',?,'%') or a.code like concat('%',?,'%') or a.item_name like concat('%',?,'%')) order by b.bar_code`
	err = pd.db.Select(&products, sql, req.Keyword, req.Keyword, req.Keyword)
	if err != nil {
		fmt.Println("error = ", err.Error())
		return resp, nil
	}
	fmt.Println("Product")
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
		SalePrice1:    x.SalePrice1,
		SalePrice2:x.SalePrice2,
		UnitCode: x.UnitCode,
		Rate1:    x.Rate1,
		PicPath1:  x.PicPath1,
	}
}
