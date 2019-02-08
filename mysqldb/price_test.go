package mysqldb

import (
	"github.com/mrtomyum/nopadol/product"
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

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
