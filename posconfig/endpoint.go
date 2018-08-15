package posconfig

import (
	"fmt"
	"context"
)

type PosConfig struct {
	CompanyName    string `db:"CompanyName"`
	CompanyAddress string `db:"CompanyAddress"`
	Telephone      string `db:"Telephone"`
	TaxId          string `db:"TaxId"`
	ArCode         string `db:"ArCode"`
	PosId          string `db:"PosId"`
	WhCode         string `db:"WhCode"`
	ShelfCode      string `db:"ShelfCode"`
	PrinterPosIp   string `db:"PrinterPosIp"`
	PrinterCopyIp  string `db:"PrinterCopyIp"`
	MachineNo      string `db:"MachineNo"`
	MachineCode    string `db:"MachineCode"`
	TaxRate        int64  `db:"TaxRate"`
}

func Create(s Service) interface{} {
	return func(ctx context.Context, req *PosConfig) (interface{}, error) {

		resp, err := s.Create(&PosConfigTemplate{
			CompanyName:    req.CompanyName,
			CompanyAddress: req.CompanyAddress,
			Telephone:      req.Telephone,
			TaxId:          req.TaxId,
			ArCode:         req.ArCode,
			TaxRate:        req.TaxRate,
			PosId:          req.PosId,
			WhCode:         req.WhCode,
			ShelfCode:      req.ShelfCode,
			MachineCode:    req.MachineCode,
			MachineNo:      req.MachineNo,
			PrinterPosIp:   req.PrinterPosIp,
			PrinterCopyIp:  req.PrinterCopyIp,
		})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil

	}
}

func SearchById(s Service) interface{}{
	return func(ctx context.Context) (interface{}, error){
		resp, err := s.SearchById()
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		},nil
	}
}