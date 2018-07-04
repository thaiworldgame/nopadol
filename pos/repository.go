package pos

import "context"

type Repository interface {
	New(context.Context, *NewPosTemplate) (NewPosResponseTemplate, error)
	SearchById(context.Context, *SearchPosByIdRequestTemplate) (SearchPosByIdResponseTemplate, error)
}
