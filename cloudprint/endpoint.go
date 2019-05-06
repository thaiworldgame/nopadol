package cloudprint

import (
	"fmt"
	"context"
)

type (
	CloudPrintResponse struct {
		DocNo string `json:"doc_no"`
	}

	CloudPrintRequest struct {
		FormType    string `json:"form_type"`
		PrinterIp   string `json:"printer_ip"`//แสดงชื่อ เครื่องพิมพ์
		PrinterPort string `json:"printer_port"`
		DocType     string `json:"doc_type"`
		//"is_print_short_form": 0,
		//“is_print_full_form”: 0,
		//"is_print_cash_form": 0,
		//"is_print_credit_form": 0

		Data       string `json:"data"`
	}

)

func CloudPrint(s Service) interface{} {
	fmt.Println("EndPoint")
	return func(ctx context.Context, req *CloudPrintRequest) (interface{}, error) {
		resp, err := s.CloudPrint(&CloudPrintRequest{FormType: req.FormType, PrinterIp: req.PrinterIp, PrinterPort: req.PrinterPort})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}
