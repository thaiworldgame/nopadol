package configuration

type Repository interface {
	ConfigSetting(req *RequestSettingTemplate) (interface{}, error)
	SearchSettingById(req *SearchByIdRequestTemplate) (interface{}, error)
}
