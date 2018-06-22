package endpoint

import (
	"github.com/mrtomyum/nopadol/employee"
	"context"
	"fmt"
)

func New(s employee.Service) employee.Endpoint {
	return &endpoint{s}
}

type endpoint struct {
	s employee.Service
}

func (ep *endpoint) SearchEmployeeById(ctx context.Context, req *employee.SearchEmployeeByIdRequest) (*employee.SearchEmployeeResponse, error) {
	emp, err := ep.s.SearchEmployeeById(ctx, &employee.SearchById{
		Id: req.Id,
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}

	resp := map_employee_response(emp)

	return &employee.SearchEmployeeResponse{EmployeeId:resp.EmployeeId, EmployeeCode:resp.EmployeeCode, EmployeeName:resp.EmployeeName, }, nil
}

func map_employee_response(x employee.EmployeeTemplate) employee.SearchEmployeeResponse{
	return employee.SearchEmployeeResponse{
		EmployeeId:x.Id,
		EmployeeCode:x.SaleCode,
		EmployeeName:x.SaleName,
	}
}
