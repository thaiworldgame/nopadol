package mysqldb

import (
	"fmt"
	"github.com/mrtomyum/nopadol/product"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_ItemInsUpd(t *testing.T) {
	testDB, err := ConnectDB("demo")
	if err != nil {
		t.Fatal(err.Error())
	}
	bars := []barcodeModel{}
	bar := barcodeModel{ItemID: 5244, BarCode: "xxxx", UnitID: 20}
	bars = append(bars, bar)

	fmt.Println("barcode object ", bars)

	req := product.ProductNewRequest{
		ItemCode:  "2120250",
		ItemName:  "น้ำยาเชื่อมๆ ",
		UnitID:    20,
		Picture:   "",
		StockType: itemStockTypeKeepStock,
	}
	fmt.Println("")
	u := itemUnitModel{}
	u.id = req.UnitID
	u.getByID(testDB)

	it := itemModel{
		Code:   req.ItemCode,
		Name:   req.ItemName,
		UnitID: req.UnitID,
		UnitCode: u.unitCode,
		StockType: req.StockType,
		PicPath1: req.Picture,
		CreateBy: "",
	}
	it.CreateTime.Time = time.Now()
	fmt.Println("req data is ",it)
	_, err = it.save(testDB)
	assert.Nil(t, err, nil)

}
