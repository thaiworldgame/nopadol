package pos

import "context"

type Repository interface {
	NewPos(context.Context, *NewPosTemplate)(NewPosResponseTemplate, error)
}