package mysqldb

import "github.com/jmoiron/sqlx"

type Coupon struct {
	CouponCode string  `json:"coupon_code"`
	Amount     float64 `json:"amount"`
}


func (c *Coupon)CheckCouponUsed(db *sqlx.DB, coupon_code string ,amount float64) bool {
	return true
}