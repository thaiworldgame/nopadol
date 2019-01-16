package drivethru

type Repository interface {
	SearchListCompany() (interface{}, error)
	SearchListMachine() (interface{}, error)
	SearchCarBrand(string) (interface{}, error)
}
