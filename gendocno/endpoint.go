package gendocno

import (
	"context"
	"fmt"
)

type (
	DocNoRequest struct {
		BranchId     int    `json:"branch_id" db:"BranchId"`
		TableCode    string `json:"table_code" db:"TableCode"`
		BillType     int64  `json:"bill_type" db:"BillType"`
		Header       string `json:"header" db:"Header"`
		UseYear      int    `json:"use_year" db:"UseYear"`
		UseMonth     int    `json:"use_month" db:"UseMonth"`
		UseDay       int    `json:"use_day" db:"UseDay"`
		UseDash      int    `json:"use_dash" db:"UseDash"`
		FormatNumber int    `json:"format_number" db:"FormatNumber"`
		ActiveStatus int    `json:"active_status" db:"ActiveStatus"`
	}
)

func Gen(s Service) interface{} {
	return func(ctx context.Context, req *DocNoRequest) (string, error) {
		resp, err := s.Gen(&DocNoTemplate{
			TableCode: req.TableCode,
			BillType:  req.BillType,
		})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return "", fmt.Errorf(err.Error())
		}
		return resp, nil
	}
}
