package mysqldb

import (
	"fmt"
	"testing"
	"time"

	"github.com/mrtomyum/nopadol/product"
	"github.com/stretchr/testify/assert"
)

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
		BarCode: req.Barcode,
		//BarCode:      "",
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
	assert.Equal(t, err, "X")

}
