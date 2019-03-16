package mysqldb

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/configuration"
)

type RequestSettingModel struct {
	Id             int64  `db:"id"`
	CompanyId      int64  `db:"company_id"`
	BranchId       int64  `db:"branch_id"`
	TaxType        string `db:"tax_type"`
	TaxRate        string `db:"tax_rate"`
	LogoPath       string `db:"logo_path"`
	DepartId       string `db:"depart_id"`
	DefSaleWhId    string `db:"def_sale_wh_id"`
	DefSaleShelfId string `db:"def_sale_shelf_id"`
	DefBuyWhId     string `db:"def_buy_wh_id"`
	DefBuyShelfId  string `db:"def_buy_shelf_id"`
	SrockStatus    int    `db:"stock_status"`
	SaleTaxType    string `db:"sale_tax_type"`
	BuyTaxType     string `db:"buy_tax_type"`
	SaleBillType   string `db:"sale_bill_type"`
	BuyBillType    string `db:"buy_bill_type"`
	UseAddress     int    `db:"use_address"`
	PosDefCustId   int    `db:"pos_def_cust_id"`
	PosDefStock    int    `db:"pos_def_stock"`
	DefCustId      string `db:"def_cust_id"`
	CreateBy       string `db:"create_by"`
	CreateTime     string `db:"create_time"`
	EditBy         string `db:"edit_by"`
	EditTime       string `db:"edit_time"`
	BranchName     string `db:"branch_name"`
	Address        string `db:"address"`
	Telephone      string `db:"telephone"`
	Fax            string `db:"fax"`
	CompanyName    string `db:"company_name"`
}

type RequestNoteModel struct {
	Id         int64  `db:"id"`
	TextNote   string `db:"text_note"`
	TypeStatus int64  `db:"type_status"`
}

type SettingRepository struct{ db *sqlx.DB }

func NewSettingRepository(db *sqlx.DB) configuration.Repository {
	return &configRepository{db}
}

func (repo *configRepository) ConfigSetting(req *configuration.RequestSettingTemplate) (resp interface{}, err error) {
	var check_id_exist int

	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))
	req.CreateTime = now.String()
	req.EditTime = now.String()

	sqlexist := `select count(id) as check_exist from configuration where id = ?`
	fmt.Println("Id = ", req.Id)
	err = repo.db.Get(&check_id_exist, sqlexist, req.Id)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	fmt.Println("check_doc_exist", check_id_exist)

	if check_id_exist == 0 {

		sql := `INSERT INTO configuration(company_id, branch_id, tax_rate, logo_path, 
			depart_id, def_sale_wh_id, def_sale_shelf_id, def_buy_wh_id, def_buy_shelf_id, 
			def_cust_id, create_by, create_time, edit_by, edit_time)
			values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		fmt.Println("sql update = ", sql) //INSERT INTO branch(branch_name, address, telephone, fax) values(?,?,?,?)
		//stock_status, sale_tax_type, buy_tax_type, sale_bill_type, buy_bill_type, pos_def_cust_id,tax_type, pos_def_stock, ?,?,?,?,?,?,?
		resp, err := repo.db.Exec(sql,
			req.CompanyId,
			req.BranchId,
			//req.TaxType,
			req.TaxRate,
			req.LogoPath,
			req.DepartId,
			req.DefSaleWhId,
			req.DefSaleShelfId,
			req.DefBuyWhId,
			req.DefBuyShelfId,
			//req.SrockStatus,
			//req.SaleTaxType,
			//req.BuyTaxType,
			//req.SaleBillType,
			//req.BuyBillType,
			//req.PosDefCustId,
			//req.PosDefStock,
			req.DefCustId,
			req.CreateBy,
			req.CreateTime,
			req.EditBy,
			req.EditTime,
		)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}
		id, _ := resp.LastInsertId()

		req.Id = id
	} else {

		sql := `Update configuration a,branch b set a.company_id=?, a.branch_id=?, a.tax_type=?, a.tax_rate=?, a.logo_path=?, 
	a.depart_id=?, a.def_sale_wh_id=?, a.def_sale_shelf_id=?, a.def_buy_wh_id=?,a.def_buy_shelf_id=?, a.stock_status=?, a.sale_tax_type=?, 
	a.buy_tax_type=?,a.sale_bill_type=?, a.buy_bill_type=?,a.use_address=?, a.pos_def_cust_id=?,a.pos_def_stock=?,a.def_cust_id=?, 
	a.create_by=?, a.create_time=?, a.edit_by=?,a.edit_time =?
	where a.id=?` //	b.branch_name=?, b.address=?, b.telephone=?, b.fax=? where a.id=? and b.id = a.branch_id
		fmt.Println("sql update = ", sql)
		id, err := repo.db.Exec(sql,
			req.CompanyId,
			req.BranchId,
			req.TaxType,
			req.TaxRate,
			req.LogoPath,
			req.DepartId,
			req.DefSaleWhId,
			req.DefSaleShelfId,
			req.DefBuyWhId,
			req.DefBuyShelfId,
			req.SrockStatus,
			req.SaleTaxType,
			req.BuyTaxType,
			req.SaleBillType,
			req.BuyBillType,
			req.UseAddress,
			req.PosDefCustId,
			req.PosDefStock,
			req.DefCustId,
			req.CreateBy,
			req.CreateTime,
			req.EditBy,
			req.EditTime,
			/*req.BranchName,
			req.Address,
			req.Telephone,
			req.Fax,*/
			req.Id,
		)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}

		rowAffect, err := id.RowsAffected()
		fmt.Println("Row Affect = ", rowAffect)
	}
	return map[string]interface{}{
		"id":                req.Id,
		"company_id":        req.CompanyId,
		"branch_id":         req.BranchId,
		"tax_type":          req.TaxType,
		"tax_rate":          req.TaxRate,
		"logo_path":         req.LogoPath,
		"depart_id":         req.DepartId,
		"def_sale_wh_id":    req.DefSaleWhId,
		"def_sale_shelf_id": req.DefSaleShelfId,
		"def_buy_wh_id":     req.DefBuyWhId,
		"def_buy_shelf_id":  req.DefBuyShelfId,
		"stock_status":      req.SrockStatus,
		"sale_tax_type":     req.SaleTaxType,
		"buy_tax_type":      req.BuyTaxType,
		"sale_bill_type":    req.SaleBillType,
		"buy_bill_type":     req.BuyBillType,
		"use_address":       req.UseAddress,
		"pos_def_cust_id":   req.PosDefCustId,
		"pos_def_stock":     req.PosDefStock,
		"def_cust_id":       req.DefCustId,
		"create_by":         req.CreateBy,
		"create_time":       req.CreateTime,
		"edit_by":           req.EditBy,
		"edit_time":         req.EditTime,
		"branch_name":       req.BranchName,
		"address":           req.Address,
		"telephone":         req.Telephone,
		"fax":               req.Fax,
	}, nil
}

func map_setting_template(x RequestSettingModel) configuration.RequestSettingTemplate {
	return configuration.RequestSettingTemplate{
		Id:             x.Id,
		CompanyId:      x.CompanyId,
		BranchId:       x.BranchId,
		TaxType:        x.TaxType,
		TaxRate:        x.TaxRate,
		LogoPath:       x.LogoPath,
		DepartId:       x.DepartId,
		DefSaleWhId:    x.DefSaleWhId,
		DefSaleShelfId: x.DefSaleShelfId,
		DefBuyWhId:     x.DefBuyWhId,
		DefBuyShelfId:  x.DefBuyShelfId,
		SrockStatus:    x.SrockStatus,
		SaleTaxType:    x.SaleTaxType,
		BuyTaxType:     x.BuyTaxType,
		SaleBillType:   x.SaleBillType,
		BuyBillType:    x.BuyBillType,
		UseAddress:     x.UseAddress,
		PosDefCustId:   x.PosDefCustId,
		PosDefStock:    x.PosDefStock,
		DefCustId:      x.DefCustId,
		CreateBy:       x.CreateBy,
		CreateTime:     x.CreateTime,
		EditBy:         x.EditBy,
		EditTime:       x.EditTime,
		BranchName:     x.BranchName,
		Address:        x.Address,
		Telephone:      x.Telephone,
		Fax:            x.Fax,
		CompanyName:    x.CompanyName,
	}
}

func (repo *configRepository) SearchSettingById(req *configuration.SearchByIdRequestTemplate) (resp interface{}, err error) {

	a := RequestSettingModel{}
	sql := `select  a.id, a.company_id, a.branch_id, a.tax_type, a.tax_rate, ifnull(a.logo_path,'') as logo_path, 
	a.depart_id, a.def_sale_wh_id, a.def_sale_shelf_id, a.def_buy_wh_id,
	a.def_buy_shelf_id, a.stock_status, a.sale_tax_type, a.buy_tax_type, a.sale_bill_type, a.buy_bill_type, 
	a.use_address, a.pos_def_cust_id, a.pos_def_stock,
	a.def_cust_id, ifnull(a.create_by,'') as create_by, ifnull(a.create_time,'') as create_time, 
	ifnull(a.edit_by,'') as edit_by,ifnull(a.edit_time,'') as edit_time, 
	b.company_id,b.branch_name, b.address, b.telephone,b.fax
	from configuration a left join branch b on a.company_id = b.company_id 
	where a.id = ? `
	err = repo.db.Get(&a, sql, req.Id)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	config_resp := map_setting_template(a)
	return map[string]interface{}{
		"id":                config_resp.Id,
		"company_id":        config_resp.CompanyId,
		"branch_id":         config_resp.BranchId,
		"tax_type":          config_resp.TaxType,
		"tax_rate":          config_resp.TaxRate,
		"logo_path":         config_resp.LogoPath,
		"depart_id":         config_resp.DepartId,
		"def_sale_wh_id":    config_resp.DefSaleWhId,
		"def_sale_shelf_id": config_resp.DefSaleShelfId,
		"def_buy_wh_id":     config_resp.DefBuyWhId,
		"def_buy_shelf_id":  config_resp.DefBuyShelfId,
		"stock_status":      config_resp.SrockStatus,
		"sale_tax_type":     config_resp.SaleTaxType,
		"buy_tax_type":      config_resp.BuyTaxType,
		"sale_bill_type":    config_resp.SaleBillType,
		"buy_bill_type":     config_resp.BuyBillType,
		"use_address":       config_resp.UseAddress,
		"pos_def_cust_id":   config_resp.PosDefCustId,
		"pos_def_stock":     config_resp.PosDefStock,
		"def_cust_id":       config_resp.DefCustId,
		"create_by":         config_resp.CreateBy,
		"create_time":       config_resp.CreateTime,
		"edit_by":           config_resp.EditBy,
		"edit_time":         config_resp.EditTime,
		"branch_name":       config_resp.BranchName,
		"address":           config_resp.Address,
		"telephone":         config_resp.Telephone,
		"fax":               config_resp.Fax,
	}, nil

}

func (repo *configRepository) SearchSettingByKeyword(req *configuration.SearchByKeywordRequestTemplate) (resp interface{}, err error) {
	var sql string
	d := []RequestSettingModel{}
	if req.Keyword == "" {
		sql = `select  a.id, a.company_id, a.branch_id, a.tax_rate, 
		b.company_id,b.branch_name, b.address, b.telephone,b.fax, c.company_name
		from configuration a left join branch b on a.branch_id = b.id left join company c on a.company_id = c.id `
		err = repo.db.Select(&d, sql)
	} else {
		sql = `select  a.id, a.company_id, a.branch_id, a.tax_rate, 
		b.company_id,b.branch_name, b.address, b.telephone,b.fax, c.company_name
		from configuration a left join branch b on a.branch_id = b.id left join company c on a.company_id = c.id 
		where a.id = ? or b.company_id = ? or b.branch_name`
		err = repo.db.Select(&d, sql, req.Keyword, req.Keyword, req.Keyword)
	}
	fmt.Println("sql = ", sql, req.Keyword)
	if err != nil {
		fmt.Println("errsss = ", err.Error())
		return resp, err
	}

	dp := []configuration.RequestSettingTemplate{}

	for _, dep := range d {
		dpline := map_setting_template(dep)
		dp = append(dp, dpline)
	}
	return dp, nil
}

func map_note_template(x RequestNoteModel) configuration.RequestNoteTemplate {
	return configuration.RequestNoteTemplate{
		Id:         x.Id,
		TextNote:   x.TextNote,
		TypeStatus: x.TypeStatus,
	}
}

func (repo *configRepository) SearchNote(req *configuration.SearchByIdRequestTemplate) (resp interface{}, err error) {
	var sql string
	d := []RequestNoteModel{}

	sql = `select  a.id, a.text_note
		from note a 
		where type_status = ?`

	err = repo.db.Select(&d, sql, req.TypeStatus)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	dp := []configuration.RequestNoteTemplate{}

	for _, dep := range d {
		dpline := map_note_template(dep)
		dp = append(dp, dpline)
	}
	return dp, nil
}
