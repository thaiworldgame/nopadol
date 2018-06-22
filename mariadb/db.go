package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

func NewDB(dbType, conn string) gorm.DB {
	db, err := gorm.Open(dbType, conn)
	if err != nil {
		panic(err.Error())
	}
//	defer db.Close()
	// db, err := gorm.Open("foundation", "dbname=gorm") // FoundationDB.
	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	// db, err := gorm.Open("sqlite3", "/tmp/gorm.db")

	// You can also use an existing database connection handle
	// dbSql, _ := sql.Open("postgres", "user=gorm dbname=gorm sslmode=disable")
	// db, _ := gorm.Open("postgres", dbSql)

	// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	// db.DB()
	// Then you could invoke `*sql.DB`'s functions with it
	err = db.DB().Ping()
	if err != nil {
		fmt.Println("DB Connection Error!!!")
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(true)
	fmt.Println("DB Connected")
	return db
}
