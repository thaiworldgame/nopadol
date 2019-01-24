package drivethru

type Repository interface {
	UserLogIn(req *UserLogInRequest) (interface{}, error)
	SearchListCompany() (interface{}, error)
	SearchListMachine() (interface{}, error)
	SearchCarBrand(string) (interface{}, error)

	SearchCustomer(string) (interface{}, error)
	SearchItem(string) (interface{}, error)
	PickupNew(req *NewPickupRequest) (interface{}, error)
	ManagePickup(req *ManagePickupRequest) (interface{}, error)
	ListQueue(req *ListQueueRequest) (interface{}, error)
	//ShiftOpen(string,float64,string)(interface{},error)

	ShiftOpen(*ShiftOpenRequest) (interface{}, error)
	ShiftClose(*ShiftCloseRequest) (interface{}, error)


}
