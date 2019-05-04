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
		PrinterIp   string `json:"printer_ip"`
		PrinterPort string `json:"printer_port"`
		DocType     string `json:"doc_type"`
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
