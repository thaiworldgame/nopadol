package drivethru

import (
	"fmt"
	"context"
)



func SearchListCompany(s Service) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		resp, err := s.SearchListCompany()
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}


		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":"Search Brand",
				"processDesc":"Success",
				"isSuccess":true,
			},
			"data": resp,
		}, nil
	}
}

func MakeListMachine(s Service) interface{}{
	return func(ctx context.Context)(interface{}, error){
		resp,err := s.SearchListMachine()
		if err != nil {
			return nil,err
		}
		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":"List Pos Machine",
				"processDesc":"Success",
				"isSuccess":true,
			},
			"data": resp,
		}, nil
	}
}


