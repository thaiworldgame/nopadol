package pos

import "context"

type Service interface {
	New(context.Context, *NewPosTemplate) (NewPosResponseTemplate, error)
	SearchById(ctx context.Context, request *SearchPosByIdRequestTemplate) (SearchPosByIdResponseTemplate, error)
}
