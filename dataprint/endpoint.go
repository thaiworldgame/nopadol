package dataprint

import (
	"fmt"
	"context"
)

func DataPrint(s Service) interface{} {
	fmt.Println("EndPoint")
	return func(ctx context.Context) (interface{}, error) {
		resp, err := s.DataPrint()
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}
