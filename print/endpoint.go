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
		FormId string `json:"form_id"`
		PosId string `json:"pos_id"`
		DocNo string `json:"doc_no"`
	}

	PosSlipResponse struct {
		DocNo string `json:"doc_no"`
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
		resp, err := s.PosDriveThruSlip(&PosDriveThruSlipRequestTemplate{DocNo: req.DocNo, AccessToken:req.AccessToken, FormId:req.FormId, PosId:req.PosId})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			//return nil, fmt.Errorf(err.Error())
			return map[string]interface{}{
				"success": false,
				"error": true,
				"message": err.Error(),
				"data": "",
			}, nil
		}

		return map[string]interface{}{
			"success": true,
			"error": false,
			"message": "",
			"data": resp,
		}, nil
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

