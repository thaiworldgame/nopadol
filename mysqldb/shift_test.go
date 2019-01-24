package mysqldb

import (
	"testing"
	"time"
	"github.com/mrtomyum/nopadol/drivethru"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"fmt"
)


var _shift ShiftModel
func Test_getShiftProfileByUUID(t *testing.T){
	testDB, err := ConnectDB("demo")
	if err != nil {
		return
	}
	_shift.Get(testDB,"xxxxxx")
}



func Test_shiftOpen(t *testing.T){
	req := drivethru.ShiftOpenRequest{
		Token: "bdebe48c-44e3-44f8-a2ad-5722a905f84b",
		MachineID:3,
		CashierID:1059,
		WhID:1,
		ChangeAmount: 3000.00,

	}
	testDB, err := ConnectDB("demo")


	uac := UserAccess{}
	uac.GetProfileByToken(testDB, req.Token)

	// init shift objects
	sh := ShiftModel{}
	sh.docDate.Time = time.Now()
	sh.companyID = uac.CompanyID
	sh.branchID = uac.BranchID
	sh.cashierID = req.CashierID
	sh.changeAmount.Float64 = req.ChangeAmount
	sh.openBy = uac.UserCode
	sh.openTime.Time = time.Now()
	sh.machineID = req.MachineID
	sh.shiftUUid = uuid.New().String()
	sh.whID = req.WhID
	fmt.Println(sh)
	shiftUUID,err := sh.Open(testDB)
	fmt.Println(shiftUUID)
	assert.Nil(t,err)

}


func Test_shiftClose(t *testing.T){
	req := drivethru.ShiftCloseRequest{
		Token: "bdebe48c-44e3-44f8-a2ad-5722a905f84b",
		ShiftUUID: "c745673a-d282-4c62-9ae8-dedfd9643754",
		SumCashAmount: 100,
		SumCreditAmount: 2000,
		SumBankAmount:0,
		SumCouponAmount:20,
		SumDepositAmount:0,

	}
	testDB, err := ConnectDB("demo")
	uac := UserAccess{}
	uac.GetProfileByToken(testDB, req.Token)

	// init shift objects



	sh := ShiftModel{}
	sh.closeBy = uac.UserCode
	sh.closeTime.Time = time.Now()
	sh.shiftUUid = req.ShiftUUID
	sh.sumOfCashAmount = req.SumCashAmount
	sh.sumOfCreditAmount = req.SumCreditAmount
	sh.sumOfCouponAmount = req.SumCouponAmount
	sh.sumOfBankAmount = req.SumBankAmount
	sh.sumOfDepositAmount = req.SumDepositAmount

	fmt.Println(sh)
	err = sh.Close(testDB)
	assert.Nil(t,err)

}