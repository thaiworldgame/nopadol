package sync

import "context"

func NewQuotation(s Service) interface{}{
	return func(ctx context.Context)(interface{}, error){

		return nil,nil
	}
}