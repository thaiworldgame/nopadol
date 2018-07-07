package delivery

type Repository interface {
	ReportDaily(req string) (interface{}, error)
}