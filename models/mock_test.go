package models

func setup() ([]*Period,[]*Customer, []*CusRank) {
	p := make([]*Period, 12)
	p[0] = &Period{Year:2015, Month:1}
	p[1] = &Period{Year:2015, Month:2}
	p[2] = &Period{Year:2015, Month:3}
	p[3] = &Period{Year:2015, Month:4}
	p[4] = &Period{Year:2015, Month:5}
	p[5] = &Period{Year:2015, Month:6}
	p[6] = &Period{Year:2015, Month:7}
	p[7] = &Period{Year:2015, Month:8}
	p[8] = &Period{Year:2015, Month:9}
	p[9] = &Period{Year:2015, Month:10}
	p[10] = &Period{Year:2015, Month:11}
	p[11] = &Period{Year:2015, Month:12}

	c := make([]*Customer, 5)
	c[0] = &Customer{Name: "Dummy Customer", CusType: COMPANY}
	c[1] = &Customer{Name: "บจก.เชียงใหม่คอนสตรัคชั่น", CusType: COMPANY}
	c[2] = &Customer{Name: "บจก.ปาล์ม การ์เด้น เชียงใหม่", CusType: COMPANY}
	c[3] = &Customer{Name: "บริษัท ส.เต็งไตรรัตน์(น่าน) จำกัด", CusType: COMPANY}

	cr := make([]*CusRank, 5)
	cr[0] = &CusRank{Period: p[0], Customer: c[0], Rank: A, KI1Continous: 4, KI2PaymentDue: 4, KI3Responsibility: 4, KI4Charactor: 4}
	cr[1] = &CusRank{Period: p[0], Customer: c[1], Rank: A, KI1Continous: 4, KI2PaymentDue: 4, KI3Responsibility: 4, KI4Charactor: 4}
	cr[2] = &CusRank{Period: p[0], Customer: c[2], Rank: A, KI1Continous: 4, KI2PaymentDue: 4, KI3Responsibility: 4, KI4Charactor: 4}
	cr[3] = &CusRank{Period: p[0], Customer: c[3], Rank: A, KI1Continous: 4, KI2PaymentDue: 4, KI3Responsibility: 4, KI4Charactor: 4}
	cr[4] = &CusRank{Period: p[0], Customer: c[4], Rank: A, KI1Continous: 4, KI2PaymentDue: 4, KI3Responsibility: 4, KI4Charactor: 4}

	return p, c, cr
}