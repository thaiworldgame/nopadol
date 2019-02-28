package mysqldb

import ("testing"
	"github.com/jmoiron/sqlx"
	"fmt"
	"github.com/stretchr/testify/assert"
)


func connectDemo(dbName string) (db *sqlx.DB, err error) {
	fmt.Println("Connect MySql")
	dsn := "root:[ibdkifu88@tcp(nopadol.net:3306)/" + dbName + "?parseTime=true&charset=utf8&loc=Local"
	//fmt.Println(dsn,"DBName =", dbName)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("sql error =", err)
		return nil, err
	}
	return db, err
}
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