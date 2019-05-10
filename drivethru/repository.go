package drivethru

type Repository interface {
	UserLogIn(req *UserLogInRequest) (interface{}, error)
	LogIn(req *LoginRequest) (interface{}, error)
	SearchListCompany() (interface{}, error)
	SearchListMachine() (interface{}, error)
	SearchListZone(string) (interface{}, error)
	SearchCarBrand(string) (interface{}, error)
	SearchListUser(req *UserRequest) (interface{}, error)

	SearchCustomer(string) (interface{}, error)
	SearchItem(string) (interface{}, error)
	PickupNew(req *NewPickupRequest) (interface{}, error)
	CancelQueue(req *PickupCancelRequest) (interface{}, error)
	ManagePickup(req *ManagePickupRequest) (interface{}, error)
	ManageCheckout(req *ManageCheckoutRequest) (interface{}, error)
	ListQueue(req *ListQueueRequest) (interface{}, error)
	PickupEdit(req *PickupEditRequest) (interface{}, error)
	QueueEdit(req *QueueEditRequest) (interface{}, error)
	EditCustomerQueue(req *QueueEditCustomer) (interface{}, error)
	QueueStatus(req *QueueStatusRequest) (interface{}, error)
	QueueProduct(req *QueueProductRequest) (interface{}, error)
	BillingDone(req *BillingDoneRequest) (interface{}, error)
	ListInvoice(req *AccessTokenRequest) (interface{}, error)
	ListPrinter(req *ListPrinterRequest) (interface{}, error)
	PrintSubmit(req *PrintSubmitRequest) (interface{}, error)
	PosList(req *AccessTokenRequest) (interface{}, error)
	PosCancel(req *QueueProductRequest) (interface{}, error)
	//ShiftOpen(string,float64,string)(interface{},error)

	ShiftOpen(*ShiftOpenRequest) (interface{}, error)
	ShiftClose(*ShiftCloseRequest) (interface{}, error)
}
