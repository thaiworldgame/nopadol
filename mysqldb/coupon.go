package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type Coupon struct {
	CouponCode string  `json:"coupon_code"`
	Amount     float64 `json:"amount"`
}

func (c *Coupon) CheckCouponUsed(db *sqlx.DB, coupon_code string, amount float64) (bool, string) {
	var exist int
	var used int

	lccommand := `select count(*) as vCount from coupon where coupon_code = ? and coupon_value = ?`
	err := db.Get(&exist, lccommand, coupon_code, amount)
	if err != nil {
		fmt.Println("error check coupon exist = ", err.Error())
		return false, err.Error()
	}

	lccommand1 := `select count(*) as vCount from coupon_receive where coupon_code = ? `
	err = db.Get(&used, lccommand1, coupon_code)
	if err != nil {
		fmt.Println("error check coupon used = ", err.Error())
		return false, err.Error()
	}

	if exist != 0 && used == 0 {
		return true, ""
	} else if (exist == 0) {
		return false, "master not have coupon"
	} else if used != 0 {
		return false, "coupon is used"
	} else {
		return true, ""
	}

}

