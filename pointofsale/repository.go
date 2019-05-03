package pointofsale

type Repository interface {
	Create(req *BasketTemplate) (interface{}, error)
	ManageBasket(req *BasketTemplate)(interface{}, error)
}
