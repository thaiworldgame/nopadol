package main

import (
	"net/http"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/denisenkom/go-mssqldb"
	"fmt"
	"github.com/mrtomyum/nopadol/mysqldb"
	"github.com/mrtomyum/nopadol/sqldb"

	"github.com/mrtomyum/nopadol/sale"
	saleservice "github.com/mrtomyum/nopadol/sale/service"
	saleendpoint "github.com/mrtomyum/nopadol/sale/endpoint"
	salehandler "github.com/mrtomyum/nopadol/sale/handler"

	"github.com/mrtomyum/nopadol/customer"
	customerservice "github.com/mrtomyum/nopadol/customer/service"
	customerenpoint "github.com/mrtomyum/nopadol/customer/endpoint"

	"github.com/mrtomyum/nopadol/employee"
	employeeservice "github.com/mrtomyum/nopadol/employee/service"
	employeeendpoint "github.com/mrtomyum/nopadol/employee/endpoint"

	"github.com/mrtomyum/nopadol/product"
	productservice "github.com/mrtomyum/nopadol/product/service"
	productendpoint "github.com/mrtomyum/nopadol/product/endpoint"

	"github.com/mrtomyum/nopadol/pos"
	posservice "github.com/mrtomyum/nopadol/pos/service"
	posendpoint "github.com/mrtomyum/nopadol/pos/endpoint"
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
	db_name := "expertshop"
	db_user := "sa"
	db_pass := "[ibdkifu"
	port := "1433"
	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", db_host, db_user, db_pass, port, db_name)
	msdb = sqlx.MustConnect("mssql", dsn)
	if msdb.Ping() != nil {
		fmt.Println("Error ")
	}

	return msdb, nil
}


func main() {

	// init repos
	saleRepo := mysqldb.NewSaleRepository(mysql_dbc)

	// init services
	saleService := saleservice.New(saleRepo)

	// init endpoints
	saleEndpoint := saleendpoint.New(saleService)

	// init customer
	customerRepo := sqldb.NewCustomerRepository(sql_dbc)
	customerService := customerservice.New(customerRepo)
	customerEndpoint := customerenpoint.New(customerService)

	// init employee
	employeeRepo := sqldb.NewEmployeeRepository(sql_dbc)
	employeeService := employeeservice.New(employeeRepo)
	employeeEndpoint := employeeendpoint.New(employeeService)

	//init product
	productRepo := sqldb.NewProductRepository(sql_dbc)
	productService := productservice.New(productRepo)
	productEndpoint := productendpoint.New(productService)

	//init pos
	posRepo := sqldb.NewPosRepository(sql_dbc)
	posService := posservice.New(posRepo)
	posEndpoint := posendpoint.New(posService)

	mux := http.NewServeMux()
	mux.Handle("/", salehandler.New(saleService))
	mux.Handle("/sale/", http.StripPrefix("/sale", sale.NewHTTPTransport(saleEndpoint)))
	mux.Handle("/customer/", http.StripPrefix("/customer/v1", customer.NewHttpTransport(customerEndpoint)))
	mux.Handle("/employee/", http.StripPrefix("/employee/v1", employee.NewHttpTransport(employeeEndpoint)))
	mux.Handle("/product/", http.StripPrefix("/product/v1", product.NewHttpTransport(productEndpoint)))
	mux.Handle("/pos/", http.StripPrefix("/pos/v1", pos.NewHttpTransport(posEndpoint)))

	http.ListenAndServe(":8081", mux)
}

