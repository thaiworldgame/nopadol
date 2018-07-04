package delivery

type Repository interface {
	ReportDaily() (interface{}, error)
}