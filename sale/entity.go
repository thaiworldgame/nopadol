package sale

import "github.com/itnopadol/go-ddd-service-boilerplate/domain3"

// Entity1 type
type Entity1 struct {
	ID     string
	Field1 string
	Field2 Entity2
	Field3 int
	Field4 domain3.Entity1
}

// Entity2 type
type Entity2 struct {
	Field1 string
	Field2 bool
}

type SaleOrder struct {
	DocNo string
	DocDate string
	ArCode string
	Subs []*SaleOrderSub
}

type SaleOrderSub struct {
	ItemCode string
	ItemName string
}

