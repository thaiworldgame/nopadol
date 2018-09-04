package delivery

import (
	"context"
	"fmt"
	//"github.com/gin-gonic/gin"
)

type reportDORequest struct {
	Date string `json:"date"`
}

func makeReportDoData(s Service) interface{} {
	return func(ctx context.Context, req *reportDORequest) (interface{}, error) {

		fmt.Println("begin endpoint.makeListUpdateByVending")
		//fmt.Println("vending id : ", req.Vending_id)
		resp, err := s.ReportDaily(req.Date) // string pass through Service
		if err != nil {
			fmt.Println("endpoint error ", err.Error())
			return nil, err
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
		//return resp, nil

	}
}

func makeGetSalesData(s Service) interface{} {
	return func(ctx context.Context, req *reportDORequest) (interface{}, error) {

		fmt.Println("begin endpoint.makeGetSalesData")
		//fmt.Println("vending id : ", req.Vending_id)
		resp, err := s.GetTeam() // string pass through Service
		if err != nil {
			fmt.Println("endpoint error ", err.Error())
			return nil, err
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
		//return resp, nil
	}
}

func makeGetTeamData(s Service) interface{} {

	return func(ctx context.Context, req *reportDORequest) (interface{}, error) {

		fmt.Println("begin endpoint.makeGetSalesData")
		//fmt.Println("vending id : ", req.Vending_id)
		resp, err := s.GetSales() // string pass through Service
		if err != nil {
			fmt.Println("endpoint error ", err.Error())
			return nil, err
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
		//return resp, nil
	}
}
