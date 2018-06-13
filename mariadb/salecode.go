package SQLserverdb

import (
	"github.com/mrtomyum/nopadol/incentive"
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/denisenkom/go-mssqldb"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// NewDomain1Repository creates domain1 repository implements by domain4
func NewSaleCodeRepository(db *sqlx.DB) incentive.Repository {
	return &incentiveRepository{db}
}

type incentiveRepository struct {db *sqlx.DB}

func (icr *incentiveRepository) SearchSaleCode(ctx context.Context, kw *incentive.EntitySearch) (ic incentive.SaleCode, err error) {
	
	fmt.Println("From API Keyword =",kw.Keyword)
	sql := `select SaleCode,SaleName from npdb.dbo.tb_inc_saleteam where SaleCode = ?`
	err = icr.db.Get(&ic, sql, kw.Keyword)
	if err != nil {
		return ic,err
	}

	fmt.Println("From API SaleCode =",ic.SaleCode)

	sqlsub := `select EnYear,MonthOfYear,ProfitCenter,TeamStatus from npdb.dbo.tb_inc_saleteam where SaleCode = ?`
	err = icr.db.Select(&ic.Subs,sqlsub,ic.SaleCode)
	if err != nil {
		return ic, err
	}



	fmt.Println("SaleCode =",ic.SaleCode)
	return so,nil
}