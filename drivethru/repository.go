package drivethru

type Repository interface {
	UserLogIn(req *UserLogInRequest) (interface{}, error)
	LogIn(req *LoginRequest) (interface{}, error)
	SearchListCompany() (interface{}, error)
	SearchListMachine() (interface{}, error)
	SearchCarBrand(string) (interface{}, error)

	SearchCustomer(string) (interface{}, error)
	SearchItem(string) (interface{}, error)
	PickupNew(req *NewPickupRequest) (interface{}, error)
	CancelQueue(req *QueueStatusRequest) (interface{}, error)
	ManagePickup(req *ManagePickupRequest) (interface{}, error)
	ManageCheckout(req *ManageCheckoutRequest) (interface{}, error)
	ListQueue(req *ListQueueRequest) (interface{}, error)
	QueueEdit(req *QueueEditRequest) (interface{}, error)
	QueueStatus(req *QueueStatusRequest) (interface{}, error)
	QueueProduct(req *QueueProductRequest) (interface{}, error)
	BillingDone(req *BillingDoneRequest) (interface{}, error)
	//ShiftOpen(string,float64,string)(interface{},error)

	ShiftOpen(*ShiftOpenRequest) (interface{}, error)
	ShiftClose(*ShiftCloseRequest) (interface{}, error)
}
