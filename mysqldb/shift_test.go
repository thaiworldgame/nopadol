package mysqldb

import (
	"testing"
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

}