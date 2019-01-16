package drivethru

type Repository interface {
	SearchListCompany() (interface{}, error)
	SearchListMachine() (interface{}, error)
}
