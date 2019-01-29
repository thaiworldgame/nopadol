package mysqldb

import (
	"fmt"
	"github.com/mrtomyum/nopadol/product"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_InsUpdItemRate(t *testing.T) {
	testDB, err := ConnectDB("demo")
	if err != nil {
		t.Fatal(err.Error())
	}

	pk := packingRate{
		ItemCode:        "2120250",
		UnitCode:        "พวง",
		RatePerBaseUnit: 5,
	}
	_, err = pk.save(testDB)
	assert.Nil(t, err, nil)
	return
}

func Test_InsUpd_Exists_Price(t *testing.T) {
	testDB, err := ConnectDB("demo")
	if err != nil {
		t.Fatal(err.Error())
	}

	// build request exits price in demo database
	req := product.PriceTemplate{
		ItemID:     5244,
		UnitID:     20,
		SalePrice1: 110,
		SalePrice2: 100,
		SaleType:   cashSaleType,
		CompanyID:  1,
	}

	// get ItemCode by ID
	it := itemModel{Id: req.ItemID}
	it.getItemCodeById(testDB)

	//get UnitCode if request object having UnitID only
	u := itemUnitModel{}
	u.getByID(testDB)

	fmt.Println("itemcode is ", it.Code)
	prc := priceModel{
		ItemId:     req.ItemID,
		ItemCode:   it.Code,
		UnitCode:   u.unitCode,
		UnitID:     req.UnitID,
		SalePrice1: req.SalePrice1,
		SalePrice2: req.SalePrice2,
		CompanyID:  req.CompanyID,
	}

	fmt.Println("price data test ->", prc)
	_, err = prc.save(testDB)
	assert.Nil(t, err, nil)
}

func Test_InsUpd_Barcode(t *testing.T) {
	testDB, err := ConnectDB("demo")
	if err != nil {
		t.Fatal(err.Error())
	}

	// build request exits price in demo database
	req := product.BarcodeTemplate{
		ItemID:    5244,
		UnitID:    20,
		Barcode:   "xxxxx",
		CompanyID: 1,
	}
	fmt.Println("req data ", req)
	bar := barcodeModel{
		ItemID: req.ItemID,
		//ItemID:int64(0),
		//BarCode: req.Barcode,
		BarCode:      "",
		UnitID:       req.UnitID,
		CreateBy:     "",
		EditBy:       "",
		ActiveStatus: 1,
		CompanyID:    req.CompanyID,
		ItemCode:     "2120250",
		UnitCode:     "กระป๋อง",
	}

	bar.CreateTime.Time = time.Now()
	bar.EditTime.Time = time.Now()
	fmt.Println("bar object before save ", bar)
	_, err = bar.save(testDB)
	assert.Nil(t, err, nil)

}
