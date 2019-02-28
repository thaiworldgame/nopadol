package mysqldb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type priceModel struct {
	Id         int64   `db:"id"`
	ItemId     int64   `db:"item_id"`
	ItemCode   string  `db:"item_code"`
	UnitID     int64   `db:"unit_id"`
	UnitCode   string  `db:"unit_code"`
	SalePrice1 float64 `db:"sale_price_1"`
	SalePrice2 float64 `db:"sale_price_2"`
	SaleType   int     `db:"salt_type"`
	CompanyID  int     `db:"company_id"`
	// todo : add start , end date
	// todo : add volumn qty
}

func (pr *priceModel) checkExitsByItemcodeUnitcodeSaletype(db *sqlx.DB) (id int64, result bool) {
	id = -1
	u := itemUnitModel{}
	u.id = pr.UnitID
	//get unitcode by id
	u.getByID(db)
	pr.UnitCode = u.unitCode
	rs := db.QueryRow(`select id
		from Price
		where item_code=? and unit_code=? and sale_type = ? limit 1 `,
		pr.ItemCode, pr.UnitCode, pr.SaleType)
	rs.Scan(&id)

	fmt.Printf("check price_id = %v  from item_id %v , unit_code %v, sale_type %v \n",
		id, pr.ItemId, pr.UnitCode, pr.SaleType)

	if id == -1 {
		return -1, false
	}
	pr.Id = id
	fmt.Printf(" price is id : %v \n\n", pr.Id)
	return id, true
}
func (pr *priceModel) verifyRequestData(db *sqlx.DB) (bool, error) {
	if pr.ItemId == 0 {
		return false, fmt.Errorf("cannot save price : item id not found ")

	}

	if pr.UnitID == 0 {
		return false, fmt.Errorf("cannot save price : unit id not found ")

	}

	if pr.CompanyID == 0 {
		return false, fmt.Errorf("cannot save price : company id not found ")
	}
	return true,nil
}

func (pr *priceModel) save(db *sqlx.DB) (result interface{}, err error) {
	//check req data
	fmt.Println("start price save ", pr)
	ok,err := pr.verifyRequestData(db)
	if err != nil {
		log.Printf(" error verify data is not ready: data -> %v", pr)
		return nil, fmt.Errorf(err.Error())
	}

	// todo : check exists item_code+unit_code+sale_type
	curID, ok := pr.checkExitsByItemcodeUnitcodeSaletype(db)
	fmt.Printf("check exists result is : %v \n ", ok)

	// case update
	if ok {
		// update and replace data with id
		fmt.Printf("case update  \n")
		_, err := db.Exec(`update Price set sale_price_1 = ? , sale_price_2 = ? where id = ?`,
			pr.SalePrice1, pr.SalePrice2, curID)

		if err != nil {
			log.Printf("update state sql command %v", err.Error())
			return nil, err
		}

		fmt.Printf("update price id %v ",curID)
		pr.Id = curID

	} else {
		//	case insert
		fmt.Printf("case insert \n")

		lcCommand := `insert into Price (
		item_code,
		sale_type,
		unit_code,
		sale_price_1,
		sale_price_2,
		unit_id,
		company_id
		) VALUES(?,?,?,?,?,?,?)`

		rs, err := db.Exec(lcCommand,
			pr.ItemCode,
			pr.SaleType,
			pr.UnitCode,
			pr.SalePrice1,
			pr.SalePrice2,
			pr.UnitID,
			pr.CompanyID,
		)

		if err != nil {
			log.Printf("sql command %v", err.Error())
			return nil, err
		}
		newID, err := rs.LastInsertId()
		if err != nil {
			log.Printf("error insert new price ", err.Error())
			return nil, err
		}

		pr.Id = newID
	}

	return map[string]interface{}{
		"result":"sucess",
	}, nil
}
