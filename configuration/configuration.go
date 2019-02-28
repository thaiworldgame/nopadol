package configuration

type RequestSettingTemplate struct {
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
}

type SearchByIdRequestTemplate struct {
	Id int64 `json:"id"`
}

type SearchByKeywordRequestTemplate struct {
	Keyword string `json:"keyword"`
}
