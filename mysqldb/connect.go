package mysqldb

import (
	"github.com/jmoiron/sqlx"
)

func ConnectDB(dbName string)(*sqlx.DB,error){

	//fmt.Println("Connect MySql")
	dsn := "root:[ibdkifu88@tcp(nopadol.net:3306)/" + dbName + "?parseTime=true&charset=utf8&loc=Local"
	//fmt.Println(dsn,"DBName =", dbName)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		//fmt.Println("sql error =", err)
		return nil, err
	}
	return db, err
}