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
	"github.com/mrtomyum/nopadol/postgres"
	"github.com/mrtomyum/nopadol/delivery"
	"fmt"
	"database/sql"
)

const (
	dbUser    = "root"
	dbPass    = "[ibdkifu88"
	dbAddress = "http://nopadol.net:3306"
	dbName    = "npdl"
)

var (
	pgEnv = "development" //default
	pgSSLMode = "disable"
	pgDbHost = "192.168.0.163"
	pgDbUser = "postgres"
	pgDbPass = "postgres"
	pgDbName = "backup"
	pgDbPort = "5432"
)
func main() {

	// Postgresql  Connect
	pgConn := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		pgDbName, pgDbUser, pgDbPass, pgDbHost, pgDbPort, pgSSLMode)

	fmt.Println(pgConn)
	//init db
	pgDb, err := sql.Open("postgres", pgConn)
	must(err)
	defer pgDb.Close()
	log.Println("connected postgres do database")

	// mySql Connect

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


	// doRepo
	doRepo := postgres.NewDeliveryRepository(pgDb)
	doService := delivery.NewService(doRepo)


	mux := http.NewServeMux()
	mux.Handle("/", salehandler.New(saleService))
	mux.Handle("/sale/", http.StripPrefix("/sale", sale.NewHTTPTransport(saleEndpoint)))
	mux.Handle("/delivery/", http.StripPrefix("/delivery", delivery.MakeHandler(doService)))

	http.ListenAndServe(":8081", mux)
}

func must(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
	}
}