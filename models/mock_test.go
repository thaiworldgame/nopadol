package models

import (
	"github.com/shopspring/decimal"
	"time"
)

func setup() ([]*Period, []*Person, []*Customer, []*CusRank, []*Employee) {
	pr := make([]*Period, 12)
	pr[0] = &Period{Year: 2015, Month: 1}
	pr[1] = &Period{Year: 2015, Month: 2}
	pr[2] = &Period{Year: 2015, Month: 3}
	pr[3] = &Period{Year: 2015, Month: 4}
	pr[4] = &Period{Year: 2015, Month: 5}
	pr[5] = &Period{Year: 2015, Month: 6}
	pr[6] = &Period{Year: 2015, Month: 7}
	pr[7] = &Period{Year: 2015, Month: 8}
	pr[8] = &Period{Year: 2015, Month: 9}
	pr[9] = &Period{Year: 2015, Month: 10}
	pr[10] = &Period{Year: 2015, Month: 11}
	pr[11] = &Period{Year: 2015, Month: 12}

	ps := make([]*Person, 5)
	ps[0] = &Person{Fname: "เกษม", Lname: "อานนทวิลาศ", Nname: "Tom", Bdate: time.Date(1974, time.October, 4, 0, 0, 0, 0, time.UTC)}
	ps[1] = &Person{Fname: "จิราภรณ์", Lname: "อานนทวิลาศ", Nname: "Jip", Bdate: time.Date(1976, time.August, 12, 0, 0, 0, 0, time.UTC)}
	ps[2] = &Person{Fname: "ทนัฐพร", Lname: "อานนทวิลาศ", Nname: "Tim", Bdate: time.Date(1976, time.August, 31, 0, 0, 0, 0, time.UTC)}
	ps[3] = &Person{Fname: "ธนันท์", Lname: "อานนทวิลาศ", Nname: "Tam", Bdate: time.Date(1974, time.June, 7, 0, 0, 0, 0, time.UTC)}

	c := make([]*Customer, 5)
	c[0] = &Customer{CusType: COMPANY, Contact: ps[0], Name: "Dummy Customer", debit: decimal.Zero, credit: decimal.Zero}
	c[1] = &Customer{CusType: COMPANY, Contact: nil, Name: "บจก.เชียงใหม่คอนสตรัคชั่น"}
	c[2] = &Customer{CusType: COMPANY, Contact: nil, Name: "บจก.ปาล์ม การ์เด้น เชียงใหม่"}
	c[3] = &Customer{CusType: COMPANY, Contact: nil, Name: "บริษัท ส.เต็งไตรรัตน์(น่าน) จำกัด"}

	cr := make([]*CusRank, 5)
	cr[0] = &CusRank{Period: pr[0], Customer: c[0], Rank: A, KI1Continous: 4, KI2PaymentDue: 4, KI3Responsibility: 4, KI4Charactor: 4}
	cr[1] = &CusRank{Period: pr[0], Customer: c[1], Rank: A, KI1Continous: 4, KI2PaymentDue: 4, KI3Responsibility: 4, KI4Charactor: 4}
	cr[2] = &CusRank{Period: pr[0], Customer: c[2], Rank: A, KI1Continous: 4, KI2PaymentDue: 4, KI3Responsibility: 4, KI4Charactor: 4}
	cr[3] = &CusRank{Period: pr[0], Customer: c[3], Rank: A, KI1Continous: 4, KI2PaymentDue: 4, KI3Responsibility: 4, KI4Charactor: 4}
	cr[4] = &CusRank{Period: pr[0], Customer: c[4], Rank: A, KI1Continous: 4, KI2PaymentDue: 4, KI3Responsibility: 4, KI4Charactor: 4}

	emp := make([]*Employee, 1)
	emp[0] = &Employee{Person: ps[0], Code: "39001", Title: "MD", salary: decimal.New(40000, 0)}

	return pr, ps, c, cr, emp
}
