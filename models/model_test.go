package models

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	pr, ps, c, _, emp := setup()
	for x := range pr {
		fmt.Println("Period:", pr[x].Year, "/", pr[x].Month)
	}
	for y := range c {
		fmt.Println("Customer:", c[y])
	}

	for p := range ps {
		fmt.Println("Person:", ps[p])
	}

	for e := range emp {
		fmt.Println("Employee:", emp[e].Fname, emp[e].Title, emp[e].Bdate, emp[e].salary)
	}
	//	for z := range cr {
	//		fmt.Println("Customr Rank:", cr[z].Period, cr[z].Customer.Name, cr[z].Rank, cr[z].KI1Continous, cr[z].KI2PaymentDue, cr[z].KI3Responsibility, cr[z].KI4Charactor )
	//	}
}
