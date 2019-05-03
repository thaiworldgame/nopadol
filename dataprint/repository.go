package dataprint

type Repository interface {
	DataPrint() (interface{}, error)
}

