package mysqldb

import (
	"github.com/jmoiron/sqlx"
	sync "github.com/mrtomyum/nopadol/dataimport"
	"github.com/mrtomyum/nopadol/product"
)

type BCItemModel struct {
	code           int64  `db:"code"`
	name1          string `db:"name1"`
	defstkunitcode string `db:"defstkunitcode"`
}

type syncRepository struct {
	db *sqlx.DB
}

func NewSyncRepository(db *sqlx.DB) sync.Repository {
	return &syncRepository{db}
}

func (repo *syncRepository)ProductUpdate(product.ProductTemplate)(string,error){
	return "success",nil
}
