package pos

type Repository interface {
	Create(req *NewPosTemplate) (interface{}, error)
	SearchById(req *SearchPosByIdRequestTemplate)(interface{}, error)
	//New(context.Context, *NewPosTemplate) (NewPosResponseTemplate, error)
	//SearchById(context.Context, *SearchPosByIdRequestTemplate) (SearchPosByIdResponseTemplate, error)
}
