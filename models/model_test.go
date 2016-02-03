package models

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	p, c, cr := setup()
	for x := range p {
		fmt.Println("Period:", p[x].Year, "/", p[x].Month)
	}
	for y := range c {
		fmt.Println("Customer:", c[y])
	}
//	for z := range cr {
//		fmt.Println("Customr Rank:", cr[z].Period, cr[z].Customer.Name, cr[z].Rank, cr[z].KI1Continous, cr[z].KI2PaymentDue, cr[z].KI3Responsibility, cr[z].KI4Charactor )
//	}
}
