package main_test

import (
	"database/sql"
	"testing"
)

func Test_DatabaseConnection(t *testing.T) {
	const (
		dbUser    = "root"
		dbPass    = "[ibdkifu88"
		dbAddress = "http://nopadol.net:3306"
		dbName    = "nopadol"
	)
	conn := dbUser + ":" + dbPass + "@tcp(" + dbAddress + ")/" + dbName
	db, err := sql.Open("mysql", conn)
	if err != nil {
		t.Errorf("Error = %s", err)
	}
	defer db.Close()
	t.Logf("DB Connection Success: %s", db.Stats())
}
