package sales

type Repository interface {
	Create(req *NewQuoTemplate) (interface{}, error)
}