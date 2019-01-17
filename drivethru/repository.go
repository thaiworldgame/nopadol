package drivethru

type Repository interface {
	SearchListCompany() (interface{}, error)
	SearchListMachine() (interface{}, error)
	SearchCarBrand(string) (interface{}, error)
	SearchCustomer(string) (interface{},error)
	SearchItem(string) (interface{},error)
}
