package service

import (
	"github.com/mrtomyum/nopadol/incentive"
	"context"
	"fmt"
)

func New(repo incentive.Repository) incentive.Service {
	return &service{repo}
}

type service struct {
	repo incentive.Repository
}

func (s *service) SearchSaleCode(ctx context.Context, kw *incentive.EntitySearch) (sic incentive.SaleCode, err error) {
	fmt.Println("keyword service = ",kw.Keyword)
	return s.repo.SearchSaleCode(ctx, kw)
}
