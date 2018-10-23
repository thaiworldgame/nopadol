package mysqldb

import (
	"github.com/mrtomyum/nopadol/companyconfig"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type RequestConfigModel struct {
	Id           int    `db:"id"`
	ComSysId     int    `db:"com_sys_id"`
	CompanyName  string `db:"company_name"`
	NameEng      string `db:"name_eng"`
	Address      string `db:"address"`
	Telephone    string `db:"telephone"`
	Fax          string `db:"fax"`
	TaxNumber    string `db:"tax_number"`
	Email        string `db:"email"`
	WebSite      string `db:"web_site"`
	TaxRate      int    `db:"tax_rate"`
	StockStatus  int    `db:"stock_status"`
	SaleTaxType  int    `db:"sale_tax_type"`
	BuyTaxType   int    `db:"buy_tax_type"`
	LogoPath     string `db:"logo_path"`
	ActiveStatus int    `db:"active_status"`
	CreateBy     string `db:"create_by"`
	CreateTime   string `db:"create_time"`
	EditBy       string `db:"edit_by"`
	EditTime     string `db:"edit_time"`
}

type configRepository struct{ db *sqlx.DB }

func NewConfigRepository(db *sqlx.DB) companyconfig.Repository {
	return &configRepository{db}
}

func (repo *configRepository) Create(req *companyconfig.RequestConfigTemplate) (resp interface{}, err error) {
	return map[string]interface{}{
		"id":           req.Id,
		"com_sys_id":   req.ComSysId,
		"company_name": req.CompanyName,
		"name_eng":     req.NameEng,
	}, nil
}

func (repo *configRepository) SearchById(req *companyconfig.SearchByIdRequestTemplate) (resp interface{}, err error) {

	a := RequestConfigModel{}

	sql := `select id,com_sys_id,company_name,ifnull(name_eng,'') as name_eng,ifnull(address,'') as address,ifnull(telephone,'') as telephone,ifnull(fax,'') as fax,ifnull(tax_number,'') as tax_number,ifnull(email,'') as email,ifnull(web_site,'') as web_site,tax_rate,stock_status,sale_tax_type,buy_tax_type,ifnull(logo_path,'') as logo_path from Company where id = ? and active_status =1`
	err = repo.db.Get(&a, sql, req.Id)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	config_resp := map_config_template(a)

	return map[string]interface{}{
		"id":            config_resp.Id,
		"com_sys_id":    config_resp.ComSysId,
		"company_name":  config_resp.CompanyName,
		"eng_name":      config_resp.NameEng,
		"address":       config_resp.Address,
		"telephone":     config_resp.Telephone,
		"fax":           config_resp.Fax,
		"tax_number":    config_resp.TaxNumber,
		"email":         config_resp.Email,
		"web_site":      config_resp.WebSite,
		"tax_rate":      config_resp.TaxRate,
		"stock_status":  config_resp.StockStatus,
		"sale_tax_type": config_resp.SaleTaxType,
		"buy_tax_type":  config_resp.BuyTaxType,
		"logo_path":     config_resp.LogoPath,
	}, nil
}

func map_config_template(x RequestConfigModel) companyconfig.RequestConfigTemplate {
	fmt.Println("CompanyName =", x.CompanyName)
	return companyconfig.RequestConfigTemplate{
		Id:          x.Id,
		ComSysId:    x.ComSysId,
		CompanyName: x.CompanyName,
		NameEng:     x.NameEng,
		Address:     x.Address,
		Telephone:   x.Telephone,
		Fax:         x.Fax,
		TaxNumber:   x.TaxNumber,
		Email:       x.Email,
		WebSite:     x.WebSite,
		TaxRate:     x.TaxRate,
		StockStatus: x.StockStatus,
		SaleTaxType: x.SaleTaxType,
		BuyTaxType:  x.BuyTaxType,
	}
}
