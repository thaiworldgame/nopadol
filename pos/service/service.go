package service

import (
	"github.com/mrtomyum/nopadol/pos"
	"context"
)

func New(repo pos.Repository) pos.Service {
	return &service{repo}
}

type service struct {
	repo pos.Repository
}

func (s service) NewPos(ctx context.Context, req *pos.NewPosTemplate) (resp pos.NewPosResponseTemplate, err error) {
	return s.repo.NewPos(ctx, req)
}

