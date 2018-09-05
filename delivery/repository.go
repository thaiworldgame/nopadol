package delivery

type Repository interface {
	ReportDaily(req string) (interface{}, error)
	ReportDailyByTeam(req *ReportDORequestByTeam) (interface{}, error)
	GetTeam() (interface{}, error)
	GetSale(req string) (interface{}, error)
}

