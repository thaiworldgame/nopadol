package incentive

import "context"

// Endpoint is the sale endpoint
type Endpoint interface {
	SearchSaleCode(context.Context, *SearchSaleCodeRequest) (*SearchSaleCodeResponse, error)
}

type (
	SearchSaleCodeRequest struct {
		Keyword string `json:"keyword"`
	}

	SearchSaleCodeResponse struct {
		Incentive SaleCode
	}
)