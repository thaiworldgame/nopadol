package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/mrtomyum/nopadol/mysqldb"
	"github.com/mrtomyum/nopadol/postgres"
	"github.com/mrtomyum/nopadol/sqldb"
	"database/sql"
	"github.com/jmoiron/sqlx"

	"github.com/mrtomyum/nopadol/delivery"

	customerservice "github.com/mrtomyum/nopadol/customer"
	employeeservice "github.com/mrtomyum/nopadol/employee"
	productservice "github.com/mrtomyum/nopadol/product"
	posservice "github.com/mrtomyum/nopadol/pos"
	posconfigservice "github.com/mrtomyum/nopadol/posconfig"
	printservice "github.com/mrtomyum/nopadol/print"
	salesservice "github.com/mrtomyum/nopadol/sales"
	gendocnoservice "github.com/mrtomyum/nopadol/gendocno"
	envservice "github.com/mrtomyum/nopadol/environment"
	configservice "github.com/mrtomyum/nopadol/companyconfig"
	p9service "github.com/mrtomyum/nopadol/p9"
)

var mysql_np *sqlx.DB
var mysql_dbc *sqlx.DB
var sql_dbc *sqlx.DB
var nebula_dbc *sqlx.DB
var (
	pgEnv     = "development" //default
	pgSSLMode = "disable"
	pgDbHost  = "192.168.0.163"
	pgDbUser  = "postgres"
	pgDbPass  = "postgres"
	pgDbName  = "backup"
	pgDbPort  = "5432"
)

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

func ConnectMysqlNP(dbName string) (db *sqlx.DB, err error) {
	fmt.Println("Connect MySql")
	dsn := "root:[ibdkifu88@tcp(nopadol.net:3306)/" + dbName + "?parseTime=true&charset=utf8&loc=Local"
	fmt.Println(dsn,"DBName =", dbName)
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

func ConnectNebula() (msdb *sqlx.DB, err error) {
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
}

func init() {
	//db, err := ConnectDB("npdl")
	mysql_db, err := ConnectMySqlDB("DriveThru_Test")
	if err != nil {
		fmt.Println(err.Error())
	}
	mysql_dbc = mysql_db

	mysql_nopadol, err := ConnectMysqlNP("npdl")
	if err != nil {
		fmt.Println(err.Error())
	}
	mysql_np = mysql_nopadol

	sql_db, err := ConnectSqlDB()
	if err != nil {
		fmt.Println(err.Error())
	}
	sql_dbc = sql_db

	nebula, err := ConnectNebula()
	if err != nil {
		fmt.Println(err.Error())
	}
	nebula_dbc = nebula

}

func main() {

	pgConn := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		pgDbName, pgDbUser, pgDbPass, pgDbHost, pgDbPort, pgSSLMode)

	fmt.Println(pgConn)

	pgDb, err := sql.Open("postgres", pgConn)
	must(err)
	defer pgDb.Close()

	// doRepo
	doRepo := postgres.NewDeliveryRepository(pgDb)
	doService := delivery.NewService(doRepo)

	// init customer
	customerRepo := mysqldb.NewCustomerRepository(mysql_np)
	customerService := customerservice.New(customerRepo)
	//customerEndpoint := customerenpoint.New(customerService)

	// init employee
	employeeRepo := mysqldb.NewEmployeeRepository(mysql_np)
	employeeService := employeeservice.New(employeeRepo)
	//employeeEndpoint := employeeendpoint.New(employeeService)

	//init product
	productRepo := mysqldb.NewProductRepository(mysql_np)
	productService := productservice.New(productRepo)
	//productEndpoint := productendpoint.New(productService)

	//init posconfig
	posconfigRepo := mysqldb.NewPosConfigRepository(mysql_dbc)
	posconfigService := posconfigservice.New(posconfigRepo)
	//posEndpoint := posendpoint.New(posService)

	//init pos
	posRepo := sqldb.NewPosRepository(sql_dbc)
	posService := posservice.New(posRepo)
	//posEndpoint := posendpoint.New(posService)

	printRepo := sqldb.NewPrintRepository(sql_dbc)
	printService := printservice.New(printRepo)

	salesRepo := mysqldb.NewSalesRepository(mysql_np)
	salesService := salesservice.New(salesRepo)

	gendocnoRepo := mysqldb.NewGenDocNoRepository(mysql_np)
	gendocnoService := gendocnoservice.New(gendocnoRepo)

	envRepo := mysqldb.NewEnvironmentRepository(mysql_np)
	envService := envservice.New(envRepo)

	configRepo := mysqldb.NewConfigRepository(mysql_np)
	configService := configservice.New(configRepo)

	p9Repo := mysqldb.NewP9Repository(mysql_np)
	p9Service := p9service.New(p9Repo)

	mux := http.NewServeMux()
	mux.Handle("/delivery/", http.StripPrefix("/delivery", delivery.MakeHandler(doService)))
	mux.Handle("/customer/", http.StripPrefix("/customer/v1", customerservice.MakeHandler(customerService)))
	mux.Handle("/employee/", http.StripPrefix("/employee/v1", employeeservice.MakeHandler(employeeService)))
	mux.Handle("/product/", http.StripPrefix("/product/v1", productservice.MakeHandler(productService)))
	mux.Handle("/posconfig/", http.StripPrefix("/posconfig/v1", posconfigservice.MakeHandler(posconfigService)))
	mux.Handle("/pos/", http.StripPrefix("/pos/v1", posservice.MakeHandler(posService)))
	mux.Handle("/print/", http.StripPrefix("/print/v1", printservice.MakeHandler(printService)))
	mux.Handle("/sales/", http.StripPrefix("/sales/v1", salesservice.MakeHandler(salesService)))
	mux.Handle("/gendocno/", http.StripPrefix("/gendocno/v1", gendocnoservice.MakeHandler(gendocnoService)))
	mux.Handle("/env/",http.StripPrefix("/env/v1",envservice.MakeHandler(envService)))
	mux.Handle("/config/",http.StripPrefix("/config/v1", configservice.MakeHandler(configService)))

	mux.Handle("/p9/",http.StripPrefix("/p9/v1", p9service.MakeHandler(p9Service)))

	http.ListenAndServe(":8081", mux)
}

func must(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
	}
}
