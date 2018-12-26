package sync



import (

	"github.com/mrtomyum/nopadol/product"
)

func New(repo Repository) (Service) {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	ProductUpdate(product.ProductTemplate)(string error)
}


func (s *Service)ProductUpdate(product.ProductTemplate)(string,error){
	return "success",nil
}

