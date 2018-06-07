package service

import (
	"github.com/mrtomyum/nopadol/sale"
	"context"
	"fmt"
)

func New(repo sale.Repository) sale.Service {
	return &service{repo}
}

type service struct {
	repo sale.Repository
}

func (s *service) NewSO(ctx context.Context, so *sale.SaleOrder) (Id int64, err error){
	return s.repo.NewSO(ctx, so)
}

func (s *service) Create(ctx context.Context, entity *sale.Entity1) (entityID string, err error) {
	fmt.Println("Entity1= ",entity.Field1)
	return s.repo.Register(ctx, entity)
}

func (s *service) Search(ctx context.Context, kw *sale.EntitySearch) (so sale.SaleOrder, err error) {
	fmt.Println("keyword service = ",kw.Keyword)
	return s.repo.Search(ctx, kw)
}

func (s *service) Update(ctx context.Context, entity *sale.Entity1) error {
	return nil
}