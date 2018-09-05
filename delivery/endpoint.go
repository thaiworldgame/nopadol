package delivery

import (
	"context"
	"fmt"
	//"github.com/gin-gonic/gin"
)

type reportDORequest struct {
	Date string `json:"date"`
}
type ReportDORequestByTeam struct {
	Date   string `json:"date"`
	Profit string `json:"profit"`
}

func makeReportDoData(s Service) interface{} {
	return func(ctx context.Context, req *reportDORequest) (interface{}, error) {

		fmt.Println("begin endpoint.makeReportDoData")
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

func makeReportDoByTeam(s Service) interface{} {
	return func(ctx context.Context, req *ReportDORequestByTeam) (interface{}, error) {
		fmt.Println("begin endpoint.makeReportDoByTeam")
		//fmt.Println("vending id : ", req.Vending_id)
		resp, err := s.ReportDailyByTeam(req) // string pass through Service
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

type GetSaleRequest struct {
	ProfitCenter string `json:"profit_center"`
}
func makeGetSalesData(s Service) interface{} {
	return func(ctx context.Context, req *GetSaleRequest) (interface{}, error) {
		fmt.Println("begin endpoint.makeGetSalesData param ->", req.ProfitCenter)
		//fmt.Println("vending id : ", req.Vending_id)

		resp, err := s.GetSales(req.ProfitCenter) // string pass through Service
		if err != nil {
			fmt.Println("endpoint error ", err.Error())
			return nil, err
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func makeGetTeamData(s Service) interface{} {

	return func(ctx context.Context) (interface{}, error) {

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
