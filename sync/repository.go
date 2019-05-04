package sync

//import "github.com/mrtomyum/nopadol/sales"

type Repository interface {
	GetNewQuotaion() (interface{}, error)
	GetNewSaleOrder() (interface{}, error)
	ConfirmTransfer(req *Logs) (interface{}, error)
}
