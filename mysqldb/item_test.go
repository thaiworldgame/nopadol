package mysqldb

import (
	"fmt"
	"github.com/mrtomyum/nopadol/product"
	"github.com/stretchr/testify/assert"
	"testing"
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

func Test_InsUpdPrice(t *testing.T) {
	testDB, err := ConnectDB("demo")
	if err != nil {
		t.Fatal(err.Error())
	}

	req := product.PriceTemplate{
		ItemID:     5244,
		UnitID:     20,

		SalePrice1: 115,
		SalePrice2: 110,
		SaleType:   cashSaleType,
		CompanyID:  1,
	}
	it := itemModel{Id:req.ItemID}
	it.getItemCodeById(testDB)
	fmt.Println("itemcode is ",it.Code)
	prc := priceModel{
		ItemId:     req.ItemID,
		ItemCode:   it.Code,
		UnitCode:   getUnitCodeByUnitID(testDB, req.UnitID),
		UnitID:     req.UnitID,
		SalePrice1: req.SalePrice1,
		SalePrice2: req.SalePrice2,
		CompanyID:  req.CompanyID,
	}

	fmt.Println("price data test ->", prc)
	_, err = prc.save(testDB)
	assert.Nil(t, err, nil)
}
