package main

import (
	"github.com/mrtomyum/nopadol/sale"
	saleservice "github.com/mrtomyum/nopadol/sale/service"
	saleendpoint "github.com/mrtomyum/nopadol/sale/endpoint"
	salehandler "github.com/mrtomyum/nopadol/sale/handler"
	"github.com/mrtomyum/nopadol/mysqldb"
	"net/http"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var dbc *sqlx.DB

func init() {
	//db, err := ConnectDB("npdl")
	db, err := ConnectDB("DriveThru_Test")
	if err != nil {
		fmt.Println(err.Error())
	}
	dbc = db
}

func main() {

	// init repos
	saleRepo := mysqldb.NewSaleRepository(dbc)

	// init services
	saleService := saleservice.New(saleRepo)

	// init endpoints
	saleEndpoint := saleendpoint.New(saleService)

	mux := http.NewServeMux()
	mux.Handle("/", salehandler.New(saleService))
	mux.Handle("/sale/", http.StripPrefix("/sale", sale.NewHTTPTransport(saleEndpoint)))

	http.ListenAndServe(":8081", mux)
}

func ConnectDB(dbName string) (db *sqlx.DB, err error) {
	fmt.Println("Connect MySql")
	//dsn := "root:[ibdkifu88@tcp(nopadol.net:3306)/" + dbName + "?parseTime=true&charset=utf8&loc=Local"
	dsn := "it:[ibdkifu@tcp(192.168.0.89:3306)/" + dbName + "?parseTime=true&charset=utf8&loc=Local"
	db, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		fmt.Println("sql error =",err)
		return nil, err
	}
	return db, err
}


