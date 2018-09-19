package sale

type Repository interface {
	Create(req *NewQTTemplate) (interface{}, error)
}