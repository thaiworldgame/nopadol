package delivery

type Repository interface {
	ReportDaily(req string) (interface{}, error)
	GetTeam() (interface{}, error)
	GetSale() (interface{}, error)
}

