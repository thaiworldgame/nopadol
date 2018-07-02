package service

import (
	"github.com/mrtomyum/nopadol/customer"
	"context"
)

func New(repo customer.Repository) customer.Service{
	return &service{repo}
}

type service struct {
	repo customer.Repository
}

func(s *service) SearchCustomerById(ctx context.Context, req *customer.SearchByIdTemplate) (cust customer.CustomerTemplate, err error){
	return s.repo.SearchCustomerById(ctx, req)
}
