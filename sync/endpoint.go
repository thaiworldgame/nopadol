package sync

import (
	"context"
	"fmt"
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

func NewSaleOrder(s Service) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		resp, err := s.GetNewSaleOrder()
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

func ConfirmTransfer(s Service) interface{} {
	return func(ctx context.Context, req *Logs) (interface{}, error) {
		fmt.Println("req = ", &req)

		for _, l := range req.LogsUUID {
			fmt.Println("l=", &l.LogUUID,l.LogUUID)
		}

		fmt.Println("byteData = ", req)
		//return nil, nil

		resp, err := s.ConfirmTransfer(req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type Before struct {
	m map[string]string
}

func contrivedAfter(b interface{}) interface{} {
	return struct {
		Before
		s []string
	}{b.(Before), []string{"new value"}}
}

func MapDataLog(req Logs) Logs {
	return Logs{
		LogsUUID: req.LogsUUID,
	}

}
