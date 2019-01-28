package mysqldb

import ("testing"
	"github.com/stretchr/testify/assert"
)


func Test_InsUpdItemRate(t *testing.T){
	testDB, err := ConnectDB("demo")
	if err != nil {
		t.Fatal(err.Error())
	}

	pk := packingRate{
			ItemCode: "2120250",
			UnitCode:"พวง",
			RatePerBaseUnit:5,
		}
	_,err = pk.save(testDB)
	assert.Nil(t,err,nil)
	return
}
