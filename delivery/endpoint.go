package delivery

import (
	"fmt"
	"context"
	//"github.com/gin-gonic/gin"
)

func makeReportDoData(s Service) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		fmt.Println("begin endpoint.makeListUpdateByVending")
		//fmt.Println("vending id : ", req.Vending_id)
		req := ""
		resp, err := s.ReportDaily(req)
		if err != nil {
			fmt.Println("endpoint error ", err.Error())
			return nil, err
		}
		//return resp, nil
		return resp, nil
	}
}

