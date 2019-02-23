package sync

import "github.com/mrtomyum/nopadol/sales"

type Repository interface{
	GetNewQoutation() (sales.NewQuoTemplate,error)

}