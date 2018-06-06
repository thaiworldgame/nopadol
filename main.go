package main

import (
	"database/sql"
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
	if err != nil {
		return
	}
	defer db.Close()

	// init repos

	// init services

	// init
}