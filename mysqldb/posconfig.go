package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/posconfig"
	"fmt"
)

type PosConfigModel struct {
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

type posconfigRepository struct{ db *sqlx.DB }

func NewPosConfigRepository(db *sqlx.DB) posconfig.Repository {
	return &posconfigRepository{db}
}

func (repo *posconfigRepository) Create(req *posconfig.PosConfigTemplate) (resp interface{}, err error) {
	return map[string]interface{}{
		"company_name":     req.CompanyName,
		"company_address": req.CompanyAddress,
	}, nil
}

func (repo *posconfigRepository)SearchById() (resp interface{}, err error){
	p := PosConfigModel{}

	sql := `select isnull(CompanyName,'') as CompanyName,isnull(CompanyAddress,'') as CompanyAddress,isnull(Telephone,'') as Telephone,isnull(TaxId,'') as TaxId,isnull(ArCode,'') as ArCode,isnull(PosId,'') as PosId,isnull(WhCode,'') as WhCode,isnull(ShelfCode,'') as ShelfCode,isnull(PrinterPosIp,'') as PrinterPosIp,isnull(PrinterCopyIp,'') as PrinterCopyIp,isnull(MachineNo,'') as MachineNo,isnull(MachineCode,'') as MachineCode from posconfig`
	err = repo.db.Get(&p, sql)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	config_resp := map_posconfig_template(p)

	return config_resp, nil
}


func map_posconfig_template(x PosConfigModel) posconfig.PosConfigTemplate {
	return posconfig.PosConfigTemplate{
		ArCode:x.ArCode,
		CompanyName:x.CompanyName,
		CompanyAddress:x.CompanyAddress,
		Telephone:x.Telephone,
		TaxId:x.TaxId,
		PosId:x.PosId,
		WhCode:x.WhCode,
		ShelfCode:x.ShelfCode,
		PrinterCopyIp:x.PrinterCopyIp,
		PrinterPosIp:x.PrinterPosIp,
		MachineNo:x.MachineNo,
		MachineCode:x.MachineCode,
		TaxRate:x.TaxRate,

	}
}