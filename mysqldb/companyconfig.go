package mysqldb

import (
	"github.com/mrtomyum/nopadol/companyconfig"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

type RequestConfigModel struct {
	Id              int    `db:"id"`
	ComSysId        int64  `db:"com_sys_id"`
	CompanyName     string `db:"company_name"`
	NameEng         string `db:"name_eng"`
	Address         string `db:"address"`
	Telephone       string `db:"telephone"`
	Fax             string `db:"fax"`
	TaxNumber       string `db:"tax_number"`
	Email           string `db:"email"`
	WebSite         string `db:"web_site"`
	TaxRate         int    `db:"tax_rate"`
	BranchName      string `db:"branch_name"`
	BranchAddress   string `db:"branch_address"`
	BranchTelephone string `db:"branch_telephone"`
	BranchFax       string `db:"branch_fax"`
	StockStatus     int    `db:"stock_status"`
	SaleTaxType     int    `db:"sale_tax_type"`
	BuyTaxType      int    `db:"buy_tax_type"`
	DefSaleWh       string `db:"def_sale_wh"`
	DefSaleShelf    string `db:"def_sale_shelf"`
	DefBuyWh        string `db:"def_buy_wh"`
	DefBuyShelf     string `db:"def_buy_shelf"`
	SaleBillType    int64  `db:"sale_bill_type"`
	BuyBillType     int64  `db:"buy_bill_type"`
	LogoPath        string `db:"logo_path"`
	DefCustId       int    `db:"def_cust_id"`
	DefCustCode     string `db:"def_cust_code"`
	ActiveStatus    int    `db:"active_status"`
	CreateBy        string `db:"create_by"`
	CreateTime      string `db:"create_time"`
	EditBy          string `db:"edit_by"`
	EditTime        string `db:"edit_time"`
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

	sql := `select a.id,com_sys_id,a.company_name,ifnull(a.name_eng,'') as name_eng,ifnull(a.address,'') as address,ifnull(a.telephone,'') as telephone,ifnull(a.fax,'') as fax,ifnull(a.tax_number,'') as tax_number,ifnull(email,'') as email,ifnull(a.web_site,'') as web_site,a.tax_rate,ifnull(b.branch_name,'') as branch_name,ifnull(b.address,'') as branch_address,ifnull(b.telephone,'') as branch_telephone,ifnull(b.fax,'') as branch_fax,b.stock_status,b.sale_tax_type,b.buy_tax_type,ifnull(b.def_sale_wh,'') as def_sale_wh,ifnull(b.def_sale_shelf,'') as def_sale_shelf,ifnull(b.def_buy_wh,'') as def_buy_wh,ifnull(b.def_buy_shelf,'') as def_buy_shelf,ifnull(sale_bill_type,0) as sale_bill_type,ifnull(buy_bill_type,0) as buy_bill_type,ifnull(b.logo_path,'') as logo_path,ifnull(c.def_cust_id,0) as def_cust_id,a.active_status from Company a inner join Branch b on a.id = b.company_id where b.id = ? and active_status = 1`
	err = repo.db.Get(&a, sql, req.Id)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	config_resp := map_config_template(a)

	return map[string]interface{}{
		"id":               config_resp.Id,
		"com_sys_id":       config_resp.ComSysId,
		"company_name":     config_resp.CompanyName,
		"eng_name":         config_resp.NameEng,
		"address":          config_resp.Address,
		"telephone":        config_resp.Telephone,
		"fax":              config_resp.Fax,
		"tax_number":       config_resp.TaxNumber,
		"email":            config_resp.Email,
		"web_site":         config_resp.WebSite,
		"tax_rate":         config_resp.TaxRate,
		"branch_name":      config_resp.BranchName,
		"branch_address":   config_resp.BranchAddress,
		"branch_telephone": config_resp.BranchTelephone,
		"branch_fax":       config_resp.BranchFax,
		"stock_status":     config_resp.StockStatus,
		"sale_tax_type":    config_resp.SaleTaxType,
		"buy_tax_type":     config_resp.BuyTaxType,
		"def_sale_wh":      config_resp.DefSaleWh,
		"def_sale_shelf":   config_resp.DefSaleShelf,
		"def_buy_wh":       config_resp.DefBuyWh,
		"def_buy_shelf":    config_resp.DefBuyShelf,
		"sale_bill_type":   config_resp.SaleBillType,
		"buy_bill_type":    config_resp.BuyBillType,
		"logo_path":        config_resp.LogoPath,
	}, nil
}

func map_config_template(x RequestConfigModel) companyconfig.RequestConfigTemplate {
	fmt.Println("CompanyName =", x.CompanyName)
	return companyconfig.RequestConfigTemplate{
		Id:              x.Id,
		ComSysId:        x.ComSysId,
		CompanyName:     x.CompanyName,
		NameEng:         x.NameEng,
		Address:         x.Address,
		Telephone:       x.Telephone,
		Fax:             x.Fax,
		TaxNumber:       x.TaxNumber,
		Email:           x.Email,
		WebSite:         x.WebSite,
		TaxRate:         x.TaxRate,
		BranchName:      x.BranchName,
		BranchAddress:   x.BranchAddress,
		BranchTelephone: x.BranchTelephone,
		BranchFax:       x.BranchFax,
		StockStatus:     x.StockStatus,
		SaleTaxType:     x.SaleTaxType,
		BuyTaxType:      x.BuyTaxType,
		DefSaleWh:       x.DefSaleWh,
		DefSaleShelf:    x.DefSaleShelf,
		DefBuyWh:        x.DefBuyWh,
		DefBuyShelf:     x.DefBuyShelf,
		SaleBillType:    x.SaleBillType,
		BuyBillType:     x.BuyBillType,
	}
}

func GenUUID() string {
	//u1 := uuid.Must(uuid.NewV4())
	//fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling
	uuid, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return err.Error()
	}
	fmt.Printf("UUIDv4: %s\n", uuid)

	// Parsing UUID from string input
	//u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	//if err != nil {
	//	fmt.Printf("Something went wrong: %s", err)
	//	return err.Error()
	//}
	//fmt.Printf("Successfully parsed: %s", u2)

	return uuid.String()
}

func (cf *RequestConfigModel) Search(db *sqlx.DB, company_id int, branch_id int) {
	lccommand := `select 	a.id,com_sys_id,a.company_name,ifnull(a.name_eng,'') as name_eng,ifnull(a.address,'') as address,ifnull(b.telephone,'') as telephone,ifnull(b.fax,'') as fax,
							ifnull(a.tax_number,'') as tax_number,ifnull(email,'') as email,ifnull(a.web_site,'') as web_site,c.tax_rate,ifnull(b.branch_name,'') as branch_name,
							ifnull(b.address,'') as branch_address,ifnull(b.telephone,'') as branch_telephone,ifnull(b.fax,'') as branch_fax,c.stock_status,c.sale_tax_type,
							c.buy_tax_type,ifnull(d.wh_code,'') as def_sale_wh,'-' as def_sale_shelf,ifnull(e.wh_code,'') as def_buy_wh,'-' as def_buy_shelf,ifnull(c.sale_bill_type,0) as sale_bill_type,
							ifnull(buy_bill_type,0) as buy_bill_type,ifnull(c.logo_path,'') as logo_path, ifnull(c.def_cust_id,0)as def_cust_id, ifnull(f.code,'') as  def_cust_code, a.active_status,a.create_by,a.create_time,a.edit_by,a.edit_time from company a inner join branch b on a.id = b.company_id left join configuration c on a.id = c.company_id and b.id = c.branch_id left join warehouse d on c.def_sale_wh_id = d.id and a.id = d.company_id and b.id = d.branch_id left join warehouse e on c.def_buy_wh_id = e.id  and a.id = e.company_id and b.id = e.branch_id left join Customer f on c.def_cust_id = f.id where a.id = ? and b.id = ? and a.active_status = 1 limit 1`
	fmt.Println("config lccommand =", lccommand, company_id, branch_id)

	err := db.Get(cf, lccommand, company_id, branch_id)
	if err != nil {
		fmt.Println("err = ", err.Error())
	}
	return

}
