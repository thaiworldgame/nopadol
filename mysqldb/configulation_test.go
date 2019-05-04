package mysqldb

/*import (
	"fmt"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
	"github.com/mrtomyum/nopadol/configuration"
)

func Test_Condata(t *testing.T) {
	testDB, err := ConnectDB("demo")
	if err != nil {
		return
	}
}
func Test_InsConfig(t *testing.T) {
	t.Errorf("this test naja")
	var err error
	testDB, err := ConnectDB("demo")
	if err != nil {
		return
	}

	// build request exits price in demo database
	req := configuration.RequestSettingTemplate{
		CompanyId:      1,
		BranchId:       1,
		TaxRate:        "7",
		LogoPath:       "",
		DepartId:       "1",
		DefSaleWhId:    "1",
		DefSaleShelfId: "1",
		DefBuyWhId:     "1",
		DefBuyShelfId:  "1",
		DefCustId:      "1",
		CreateBy:       "",
		//CreateTime:     "",
		EditBy: "",
		//EditTime:       "",
	}
	fmt.Println("req data ", req)

	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))
	req.CreateTime = now.String()
	req.EditTime = now.String()
	con := RequestSettingModel{
		CompanyId:      req.CompanyId,
		BranchId:       req.BranchId,
		TaxRate:        req.TaxRate,
		LogoPath:       req.LogoPath,
		DepartId:       req.DepartId,
		DefSaleWhId:    req.DefSaleWhId,
		DefSaleShelfId: req.DefSaleShelfId,
		DefBuyWhId:     req.DefBuyWhId,
		DefBuyShelfId:  req.DefBuyShelfId,
		DefCustId:      req.DefCustId,
		CreateBy:       req.CreateBy,
		CreateTime:     req.CreateTime,
		EditBy:         req.EditBy,
		EditTime:       req.EditTime,
	}
	fmt.Println("con object before save ", con)
	_, err = con.ConfigSetting(testDB)
	assert.Equal(t, err, "X")
}

func Test_searchConfigById(t *testing.T) {
	var Id = 1
	test := configuration.SearchByIdRequestTemplate{
		Id: 1,
	}
	got := SearchSettingByKeyword(Id)
	want := RequestSettingModel{
		Id:             1,
		CompanyId:      1,
		BranchId:       1,
		TaxType:        "0",
		TaxRate:        "7",
		LogoPath:       "",
		DepartId:       "1",
		DefSaleWhId:    "1",
		DefSaleShelfId: "1",
		DefBuyWhId:     "1",
		DefBuyShelfId:  "1",
		SrockStatus:    0,
		SaleTaxType:    "1",
		BuyTaxType:     "1",
		SaleBillType:   "1",
		BuyBillType:    "1",
		UseAddress:     0,
		PosDefCustId:   0,
		PosDefStock:    0,
		DefCustId:      "1",
		CreateBy:       "",
		CreateTime:     "",
		EditBy:         "",
		EditTime:       "",
		BranchName:     "",
		Address:        "",
		Telephone:      "",
		Fax:            "",
	}
}*/
