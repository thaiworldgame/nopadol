package sync

import "context"

func NewQuotation(s Service) interface{}{
	return func(ctx context.Context)(interface{}, error){
		resp,err := s.GetNewQoutaion()
		if err != nil {
			return nil,err
		}
		return resp,nil
	}
}



func makeDone(s Service) interface{}{
	return func(ctx context.Context)(interface{}, error){
		type request struct {
			req_uuid string
		}
		return map[string]interface{}{
			"result " : "done",
		},nil
	}
}