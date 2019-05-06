package mysqldb

import (
	"testing"
	//"fmt"
	"github.com/magiconair/properties/assert"
)

func Test_user_get_by_token(t *testing.T) {
	u := UserAccess{}
	testDb, _ := ConnectDB("demo")
	u.GetProfileByToken(testDb, "bdebe48c-44e3-44f8-a2ad-5722a905f84b")
	//fmt.Println(u)
	assert.Equal(t, u.UserCode, "450404")
}
