package pos

import "context"

type Service interface {
	NewPos(context.Context, *NewPosTemplate) (NewPosResponseTemplate, error)
}
