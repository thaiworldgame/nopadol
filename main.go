package main

import (
	"github.com/jmoiron/sqlx"
	"net/http"
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/mrtomyum/nopadol/mysqldb"
	"github.com/mrtomyum/nopadol/sale"
	log "github.com/Sirupsen/logrus"
	saleendpoint "github.com/mrtomyum/nopadol/sale/endpoint"
	salehandler "github.com/mrtomyum/nopadol/sale/handler"
	saleservice "github.com/mrtomyum/nopadol/sale/service"
)

const (
	dbUser    = "root"
	dbPass    = "[ibdkifu88"
	dbAddress = "http://nopadol.net:3306"
	dbName    = "npdl"
)

func main() {
	conn := dbUser + ":" + dbPass + "@tcp(" + dbAddress + ")/" + dbName + "?parseTime=true&charset=utf8&loc=Local"
	db, err := sqlx.Open("mysql", conn)
	if err != nil {
		log.Error("sql error =", err)
	}
	defer db.Close()
	log.Println("Connect MySql")

	// init repos
	saleRepo := mysqldb.NewSaleRepository(db)

	// init services
	saleService := saleservice.New(saleRepo)

	// init endpoints
	saleEndpoint := saleendpoint.New(saleService)

	mux := http.NewServeMux()
	mux.Handle("/", salehandler.New(saleService))
	mux.Handle("/sale/", http.StripPrefix("/sale", sale.NewHTTPTransport(saleEndpoint)))

	http.ListenAndServe(":8081", mux)
}
