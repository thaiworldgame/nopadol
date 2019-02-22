package config

type Repository interface {
	SettingSys(req *SearchByIdTemplate) (interface{}, error)
}
