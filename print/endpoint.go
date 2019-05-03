package print

import (
	"fmt"
	"context"
)

type (
	PosSlipRequest struct {
		DocNo string `json:"doc_no"`
	}

	PosDriveThruSlipRequest struct {
		//DbHost string `json:"db_host"`
		//DbName string `json:"db_name"`
		//DbUser string `json:"db_user"`
		//DbPass string `json:"db_pass"`
		//HostIP string `json:"host_ip"`
		AccessToken string `json:"access_token"`
		FormId      string `json:"form_id"`
		PosId       string `json:"pos_id"`
		DocNo       string `json:"doc_no"`
	}

	PosSlipResponse struct {
		DocNo string `json:"doc_no"`
	}

	AutoPrintRequest struct {
		FormType    string `json:"form_type"`
		PrinterIp   string `json:"printer_ip"`
		PrinterPort string `json:"printer_port"`
		DocType     string `json:"doc_type"`
		DocNo       string `json:"doc_no"`
		Subs        []Sub  `json:"subs"`
	}

	Sub struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}
)

func PosSlip(s Service) interface{} {
	fmt.Println("EndPoint")
	return func(ctx context.Context, req *PosSlipRequest) (interface{}, error) {
		resp, err := s.PosSlip(&PosSlipRequestTemplate{DocNo: req.DocNo})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func PosDriveThruSlip(s Service) interface{} {
	fmt.Println("EndPoint")
	return func(ctx context.Context, req *PosDriveThruSlipRequest) (interface{}, error) {
		resp, err := s.PosDriveThruSlip(&PosDriveThruSlipRequestTemplate{DocNo: req.DocNo, AccessToken: req.AccessToken, FormId: req.FormId, PosId: req.PosId})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			//return nil, fmt.Errorf(err.Error())
			return map[string]interface{}{
				"success": false,
				"error":   true,
				"message": err.Error(),
				"data":    "",
			}, nil
		}

		return map[string]interface{}{
			"success": true,
			"error":   false,
			"message": "",
			"data":    resp,
		}, nil
	}
}

//func AutoPrint(s Service) interface{} {
//	fmt.Println("EndPoint")
//	return func(ctx context.Context, req *AutoPrintRequest) (interface{}, error) {
//		//p := map_print_request(req)
//		//
//		//fmt.Println("p =", p)
//		//
//		//for _, subs := range req.Subs {
//		//	fmt.Println(subs)
//		//	itemline := map_print_sub_request(subs)
//		//	p.Subs = append(p.Subs, itemline)
//		//}
//
//		resp, err := s.AutoPrint(&AutoPrintRequest{PrinterIp: req.PrinterIp, PrinterPort: req.PrinterPort, Subs: req.Subs,})
//		if err != nil {
//			fmt.Println("endpoint error =", err.Error())
//			return nil, fmt.Errorf(err.Error())
//		}
//		return map[string]interface{}{
//			"data": resp,
//		}, nil
//	}
//}

func map_print_request(x *AutoPrintRequest) AutoPrintRequest {
	var subs []Sub

	return AutoPrintRequest{
		DocType: x.DocType,
		DocNo:   x.DocNo,
		Subs:    subs,
	}
}

func map_print_sub_request(x Sub) Sub {
	return Sub{
		Code: x.Code,
		Name: x.Name,
	}
}

//func Slip(s Service) interface{} {
//	fmt.Println("EndPoint")
//	return func(ctx context.Context, req *InvoiceRequest) (interface{}, error) {
//		resp, err := s.Slip(&InvoiceTemplate{})
//		if err != nil {
//			fmt.Println("endpoint error =", err.Error())
//			//return nil, fmt.Errorf(err.Error())
//			return map[string]interface{}{
//				"success": false,
//				"error":   true,
//				"message": err.Error(),
//				"data":    "",
//			}, nil
//		}
//
//		return map[string]interface{}{
//			"success": true,
//			"error":   false,
//			"message": "",
//			"data":    resp,
//		}, nil
//	}
//}
