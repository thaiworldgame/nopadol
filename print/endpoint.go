package print

import (
	"fmt"
	"context"
)

type (
	PosSlipRequest struct {
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

