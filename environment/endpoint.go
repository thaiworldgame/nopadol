package environment

import (
	"context"
	"fmt"
)

type (
	DepartmentResponse struct {
		Id           int64  `json:"id"`
		Code         string `json:"code"`
		Name         string `json:"name"`
		ActiveStatus int64  `json:"active_status"`
		CreateBy     string `json:"create_by"`
		CreateTime   string `json:"create_time"`
		EditBy       string `json:"edit_by"`
		EditTime     string `json:"edit_time"`
	}

	SearchByIdRequest struct {
		Id int64 `json:"id"`
	}

	SearchByKeywordRequest struct {
		Keyword string `json:"keyword"`
	}
)

//Department///////////////////////////////////////////////////////////////////////////////////

func SearchDepartmentById(s Service) interface{} {
	return func(ctx context.Context, req *SearchByIdRequest) (interface{}, error) {
		resp, err := s.SearchDepartmentById(&SearchByIdTemplate{Id: req.Id})
		if err != nil {
			fmt.Println("endpoint error = ", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func SearchDepartmentByKeyword(s Service) interface{} {
	return func(ctx context.Context, req *SearchByKeywordRequest) (interface{}, error) {
		resp, err := s.SearchDepartmentByKeyword(&SearchByKeywordTemplate{Keyword: req.Keyword})
		if err != nil {
			fmt.Println("endpoint error = ", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

//Project///////////////////////////////////////////////////////////////////////////////////

func SearchProjectById(s Service) interface{} {
	return func(ctx context.Context, req *SearchByIdRequest) (interface{}, error) {
		resp, err := s.SearchProjectById(&SearchByIdTemplate{Id: req.Id})
		if err != nil {
			fmt.Println("endpoint error = ", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func SearchProjectByKeyword(s Service) interface{} {
	return func(ctx context.Context, req *SearchByKeywordRequest) (interface{}, error) {
		resp, err := s.SearchProjectByKeyword(&SearchByKeywordTemplate{Keyword: req.Keyword})
		if err != nil {
			fmt.Println("endpoint error = ", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

//Allocate///////////////////////////////////////////////////////////////////////////////////

func SearchAllocateById(s Service) interface{} {
	return func(ctx context.Context, req *SearchByIdRequest) (interface{}, error) {
		resp, err := s.SearchAllocateById(&SearchByIdTemplate{Id: req.Id})
		if err != nil {
			fmt.Println("endpoint error = ", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func SearchAllocateByKeyword(s Service) interface{} {
	return func(ctx context.Context, req *SearchByKeywordRequest) (interface{}, error) {
		resp, err := s.SearchAllocateByKeyword(&SearchByKeywordTemplate{Keyword: req.Keyword})
		if err != nil {
			fmt.Println("endpoint error = ", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func SearchCustContactById(s Service) interface{} {
	type response struct {
		Response string                 `json:"response"`
		Message  string                 `json:"message"`
		Data     []FindCustContactModel `json:"data"`
	}
	return func(ctx context.Context, req *SearchByIdRequest) (*response, error) {
		resp, err := s.SearchCustContactByIdService(&SearchByIdTemplate{Id: req.Id})
		if err != nil {
			return &response{Response: "fail", Message: err.Error()}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func SearchCustContactByKeyword(s Service) interface{} {
	type response struct {
		Response string                 `json:"response"`
		Message  string                 `json:"message"`
		Data     []FindCustContactModel `json:"data"`
	}
	return func(ctx context.Context, req *SearchByKeywordRequest) (*response, error) {
		resp, err := s.SearchCustContactByKeywordService(&SearchByKeywordTemplate{Keyword: req.Keyword})
		if err != nil {
			return &response{Response: "fail", Message: err.Error()}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}
