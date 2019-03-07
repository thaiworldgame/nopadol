package sync

import (
	"context"
)

type Log struct {
	LogUUID string `json:"log_uuid"`
}

type Logs struct {
	LogsUUID []*Log `json:"data"`
}

func NewQuotation(s Service) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		resp, err := s.GetNewQuotaion()
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

func confirmTransfer(s Service) interface{} {
	type request struct {
		log_uuid string
	}
	return func(ctx context.Context,req *request) (interface{}, error) {
		qt := Log{LogUUID:req.log_uuid}

		resp, err := s.ConfirmTransfer(qt)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}



