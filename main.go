package main

import (
	"github.com/mrtomyum/nopadol/incentive"
	incentiveService "github.com/mrtomyum/nopadol/incentive/service"
	incentiveEndpoint "github.com/mrtomyum/nopadol/incentive/endpoint"
	incentivehandler "github.com/mrtomyum/nopadol/incentive/handler"
	"net/http"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

var dbc *sqlx.DB


func init() {
	dbc = ConnectSql()

}

func main() {

	// init repos
	incentiveRepo := mariadb.NewSaleCodeRepository(dbc)

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
}
