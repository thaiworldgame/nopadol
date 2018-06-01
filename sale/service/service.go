package service

import (
	"github.com/mrtomyum/nopadol/sale"
	"context"
)

func New(repo sale.Repository) sale.Service {
	return &service{repo}
}

type service struct {
	repo sale.Repository
}

func (s *service) Create(ctx context.Context, entity *sale.Entity1) (entityID string, err error) {
	return s.repo.Register(ctx, entity)
}

func (s *service) Update(ctx context.Context, entity *sale.Entity1) error {
	return nil
}