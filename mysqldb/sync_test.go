package mysqldb

import ("testing"
	//"github.com/jmoiron/sqlx"
	"fmt"
	"github.com/stretchr/testify/assert"
)


func Test_sync_get_qoutation(t *testing.T){
	db,err := connectDemo("npdl")
	if err != nil {
		t.Fatalf("error connect db ",err.Error())
	}
	sync := syncLogs{}
	synclist,err := sync.getWaitQuotation(db)
	fmt.Println(synclist)
	assert.Equal(t,err,nil)
}