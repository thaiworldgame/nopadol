package endpoint

import (
	"github.com/mrtomyum/nopadol/customer"
	"context"
	"fmt"
)

func New(s customer.Service) customer.Endpoint{
	return &endpoint{s}
}

type endpoint struct {
	s customer.Service
}

func (ep *endpoint)SearchCustomerById(ctx context.Context, req *customer.SearchCustomerByIdRequest) (*customer.SearchCustomerResponse, error){
	cust, err := ep.s.SearchCustomerById(ctx, &customer.SearchByIdTemplate{
		Id:req.Id,
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}

	fmt.Println("Customer DB = ",cust)

	resp := map_cutomer_response(cust)

	return &resp, nil

	//return &customer.SearchCustomerResponse{
	//	CustomerId: Resp.CustomerId, CustomerCode: Resp.CustomerCode, CustomerName: Resp.CustomerName,
	//},nil

}

func map_cutomer_response(x customer.CustomerTemplate) customer.SearchCustomerResponse{
	return customer.SearchCustomerResponse{
		CustomerId:x.CustomerId,
		CustomerCode:x.CustomerCode,
		CustomerName:x.CustomerName,
	}
}