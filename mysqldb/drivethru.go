package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/drivethru"
)

type drivethruRepository struct{ db *sqlx.DB }


func NewDrivethruRepository(db *sqlx.DB) drivethru.Repository {
	return &drivethruRepository{db}
}


func (repo *drivethruRepository) SearchListCompany() (interface{}, error) {
	sql := `select `
	return nil,nil
}