package main

import (
	"github.com/mrtomyum/nopadol/sale"
	saleservice "github.com/mrtomyum/nopadol/sale/service"
	saleendpoint "github.com/mrtomyum/nopadol/sale/endpoint"
	salehandler "github.com/mrtomyum/nopadol/sale/handler"
	"github.com/mrtomyum/nopadol/mysqldb"
	"net/http"
)


func main() {

	// init repos
	saleRepo := mysqldb.NewSaleRepository()

	// init services
	saleService := saleservice.New(saleRepo)

	// init endpoints
	saleEndpoint := saleendpoint.New(saleService)

	mux := http.NewServeMux()
	mux.Handle("/", salehandler.New(saleService))
	mux.Handle("/sale/", http.StripPrefix("/sale", sale.NewHTTPTransport(saleEndpoint)))

	http.ListenAndServe(":8081", mux)
}