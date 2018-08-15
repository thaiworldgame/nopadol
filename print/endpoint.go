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
		DbHost string `json:"db_host"`
		DbName string `json:"db_name"`
		DbUser string `json:"db_user"`
		DbPass string `json:"db_pass"`
		HostIP string `json:"host_ip"`
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
		resp, err := s.PosDriveThruSlip(&PosDriveThruSlipRequestTemplate{DocNo: req.DocNo, DbHost:req.DbHost, DbName:req.DbName, DbUser:req.DbUser, DbPass:req.DbPass, HostIP:req.HostIP})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

