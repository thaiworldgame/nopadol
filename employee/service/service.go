package service

import (
	"github.com/mrtomyum/nopadol/employee"
	"context"
)

func New(repo employee.Repository) employee.Service{
	return &service{repo}
}

type service struct {
	repo employee.Repository
}

func(s *service) SearchEmployeeById(ctx context.Context, req *employee.SearchById) (emp employee.EmployeeTemplate, err error){
	return s.repo.SearchEmployeeById(ctx, req)
}