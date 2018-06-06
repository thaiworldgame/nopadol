package main

import (
	"github.com/mrtomyum/nopadol/mariadb"
	"github.com/mrtomyum/nopadol/incentive"
)

const (
	dbUser = "root"
	dbPass = "[ibdkifu88"
	dbAddress = "http://nopadol.net:3306"
	dbName = "nopadol"
)

func main() {
	conn := dbUser + ":" + dbPass + "@tcp(" + dbAddress + ")/" + dbName
	db, err := sql.Open("mysql", conn)
	must(err)
	defer db.Close()

	// init repos
	saleRepo := mariadb.NewIncentiveRepository()

	// init services
	saleService := sale.Service.New(saleRepo)

	// init
}

func must(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
	}
}