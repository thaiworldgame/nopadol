package mysqldb

import (
	"github.com/mrtomyum/nopadol/drivethru"
	"testing"
	"time"

	"fmt"
	//"github.com/google/uuid"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

var _shift ShiftModel

func Test_getShiftProfileByUUID(t *testing.T) {
	testDB, err := ConnectDB("demo")
	if err != nil {
		return
	}
	_shift.Get(testDB, "xxxxxx")
}

func Test_shiftOpen(t *testing.T) {
	req := drivethru.ShiftOpenRequest{
		AccessToken:        "bdebe48c-44e3-44f8-a2ad-5722a905f84b",
		MachineID:    "3",
		ChangeAmount: "3000.00",
		Remark:"Test",
	}
	testDB, err := ConnectDB("demo")

	uac := UserAccess{}
	uac.GetProfileByToken(testDB, req.AccessToken)

	// init shift objects
	sh := ShiftModel{}
	sh.DocDate = time.Now().String()
	sh.CompanyID = uac.CompanyID
	sh.BranchID = uac.BranchID
	sh.CashierID = uac.Id
	sh.ChangeAmount = req.ChangeAmount
	sh.OpenBy = uac.UserCode
	sh.OpenTime = time.Now().String()
	sh.MachineID = req.MachineID
	shiftUUid , err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Get UUID Error = ", err)
	}
	sh.ShiftUUid = shiftUUid.String()
	shiftUUID, err := sh.Open(testDB)
	fmt.Println(shiftUUID)
	assert.Nil(t, err)
	return
}

func Test_shiftClose(t *testing.T) {

	req := drivethru.ShiftCloseRequest{
		AccessToken:            "bdebe48c-44e3-44f8-a2ad-5722a905f84b",
		ShiftUUID:        "c745673a-d282-4c62-9ae8-dedfd9643754",
		SumCashAmount:    "100",
		SumCreditAmount:  "2000",
		SumBankAmount:    "0",
		SumCouponAmount:  "20",
		SumDepositAmount: "0",
	}
	testDB, err := ConnectDB("demo")
	uac := UserAccess{}
	uac.GetProfileByToken(testDB, req.AccessToken)

	// init shift objects

	sh := ShiftModel{}
	sh.CloseBy = uac.UserCode
	sh.CloseTime = time.Now().String()
	sh.ShiftUUid = req.ShiftUUID
	sh.SumOfCashAmount = req.SumCashAmount
	sh.SumOfCreditAmount = req.SumCreditAmount
	sh.SumOfCouponAmount = req.SumCouponAmount
	sh.SumOfBankAmount = req.SumBankAmount
	sh.SumOfDepositAmount = req.SumDepositAmount

	fmt.Println(sh)
	err = sh.Close(testDB)
	assert.Nil(t, err)

}
