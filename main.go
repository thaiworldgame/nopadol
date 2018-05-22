package main

import (
	"github.com/mrtomyum/nopadol/mariadb"
	"github.com/mrtomyum/nopadol/sale/service"
)


func main() {
	// init repos
	saleRepo := mariadb.NewWebRepository()

	// init services
	saleService := saleService.New(saleRepo)

	// init
}