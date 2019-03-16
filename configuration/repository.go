package configuration

type Repository interface {
	ConfigSetting(req *RequestSettingTemplate) (interface{}, error)
	SearchSettingById(req *SearchByIdRequestTemplate) (interface{}, error)
	SearchSettingByKeyword(req *SearchByKeywordRequestTemplate) (interface{}, error)
	SearchNote(req *SearchByIdRequestTemplate) (interface{}, error)
}
