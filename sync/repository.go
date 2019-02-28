package sync

//import "github.com/mrtomyum/nopadol/sales"

type Repository interface{
	GetNewQuotaion() (interface{},error)

}