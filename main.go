package main

import (
<<<<<<< HEAD
	"github.com/mrtomyum/nopadol/incentive"
	"github.com/mrtomyum/nopadol/sqldb"
	incentiveService "github.com/mrtomyum/nopadol/incentive/service"
	incentiveEndpoint "github.com/mrtomyum/nopadol/incentive/endpoint"
	incentivehandler "github.com/mrtomyum/nopadol/incentive/handler"
	"net/http"
	"github.com/jmoiron/sqlx"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

var dbc *sqlx.DB


func init() {
	dbc = ConnectSql()

=======
	"net/http"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/denisenkom/go-mssqldb"
	"fmt"

	"github.com/mrtomyum/nopadol/sale"
	saleservice "github.com/mrtomyum/nopadol/sale/service"
	saleendpoint "github.com/mrtomyum/nopadol/sale/endpoint"
	salehandler "github.com/mrtomyum/nopadol/sale/handler"

	"github.com/mrtomyum/nopadol/customer"
	customerservice "github.com/mrtomyum/nopadol/customer/service"
	customerenpoint "github.com/mrtomyum/nopadol/customer/endpoint"


	"github.com/mrtomyum/nopadol/mysqldb"
	"github.com/mrtomyum/nopadol/sqldb"
)

var mysql_dbc *sqlx.DB
var sql_dbc *sqlx.DB

func init() {
	//db, err := ConnectDB("npdl")
	mysql_db, err := ConnectMySqlDB("DriveThru_Test")
	if err != nil {
		fmt.Println(err.Error())
	}
	mysql_dbc = mysql_db

	sql_db, err := ConnectSqlDB()
	if err != nil {
		fmt.Println(err.Error())
	}
	sql_dbc = sql_db


}

func ConnectMySqlDB(dbName string) (db *sqlx.DB, err error) {
	fmt.Println("Connect MySql")
	//dsn := "root:[ibdkifu88@tcp(nopadol.net:3306)/" + dbName + "?parseTime=true&charset=utf8&loc=Local"
	dsn := "it:[ibdkifu@tcp(192.168.0.89:3306)/" + dbName + "?parseTime=true&charset=utf8&loc=Local"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("sql error =", err)
		return nil, err
	}
	return db, err
}

func ConnectSqlDB() (msdb *sqlx.DB, err error) {
	db_host := "192.168.0.7"
	db_name := "bcnp"
	db_user := "sa"
	db_pass := "[ibdkifu"
	port := "1433"
	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", db_host, db_user, db_pass, port, db_name)
	msdb = sqlx.MustConnect("mssql", dsn)
	if msdb.Ping() != nil {
		fmt.Println("Error ")
	}

	return msdb, nil
>>>>>>> 9ac04656f87c91029ac46d7f1357a258638af627
}

func main() {

	// init repos
<<<<<<< HEAD
	incentiveRepo := sqldb.NewSaleCodeRepository(dbc)

	// init services
	incentiveService := incentiveservice.New(incentiveRepo)

	// init endpoints
	incentiveEndpoint := incentiveendpoint.New(incentiveService)

	mux := http.NewServeMux()
	mux.Handle("/", incentivehandler.New(incentiveService))
	mux.Handle("/incentive/", http.StripPrefix("/incentive", incentive.NewHTTPTransport(incentiveEndpoint)))

	http.ListenAndServe(":9020", mux)
}


func ConnectSql()(msdb *sqlx.DB) {
	db_host := "192.168.0.7"
	db_name := "npdb"
	db_user := "sa"
	db_pass := "[ibdkifu"
	port := "1433"
	//dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", db_host, db_user, db_pass, port, db_name)
	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", db_host, db_user, db_pass, port, db_name)
	msdb = sqlx.MustConnect("mssql",dsn)

	if (msdb.Ping() != nil) {
		fmt.Println("Error")
	}
	return msdb
=======
	saleRepo := mysqldb.NewSaleRepository(mysql_dbc)
	//saleRepo := mock.NewSaleRepository(dbc)

	// init services
	saleService := saleservice.New(saleRepo)

	// init endpoints
	saleEndpoint := saleendpoint.New(saleService)

	// init customer
	customerRepo := sqldb.NewCustomerRepository(sql_dbc)
	customerService := customerservice.New(customerRepo)
	customerEndpoint := customerenpoint.New(customerService)

	mux := http.NewServeMux()
	mux.Handle("/", salehandler.New(saleService))
	mux.Handle("/sale/", http.StripPrefix("/sale", sale.NewHTTPTransport(saleEndpoint)))
	mux.Handle("/customer/", http.StripPrefix("/customer", customer.NewHttpTransport(customerEndpoint)))

	http.ListenAndServe(":8081", mux)
>>>>>>> 9ac04656f87c91029ac46d7f1357a258638af627
}
