package configuration

type RequestSettingTemplate struct {
	Id             int64  `json:"id"`
	CompanyId      int64  `json:"company_id"`
	BranchId       int64  `json:"branch_id"`
	TaxType        string `json:"tax_type"`
	TaxRate        string `json:"tax_rate"`
	LogoPath       string `json:"logo_path"`
	DepartId       string `json:"depart_id"`
	DefSaleWhId    string `json:"def_sale_wh_id"`
	DefSaleShelfId string `json:"def_sale_shelf_id"`
	DefBuyWhId     string `json:"def_buy_wh_id"`
	DefBuyShelfId  string `json:"def_buy_shelf_id"`
	SrockStatus    int    `json:"stock_status"`
	SaleTaxType    string `json:"sale_tax_type"`
	BuyTaxType     string `json:"buy_tax_type"`
	SaleBillType   string `json:"sale_bill_type"`
	BuyBillType    string `json:"buy_bill_type"`
	UseAddress     int    `json:"use_address"`
	PosDefCustId   int    `json:"pos_def_cust_id"`
	PosDefStock    int    `json:"pos_def_stock"`
	DefCustId      string `json:"def_cust_id"`
	CreateBy       string `json:"create_by"`
	CreateTime     string `json:"create_time"`
	EditBy         string `json:"edit_by"`
	EditTime       string `json:"edit_time"`
	BranchName     string `json:"branch_name"`
	Address        string `json:"address"`
	Telephone      string `json:"telephone"`
	Fax            string `json:"fax"`
	CompanyName    string `json:"company_name"`
}

type SearchByIdRequestTemplate struct {
	Id int64 `json:"id"`
}

type SearchByKeywordRequestTemplate struct {
	Keyword string `json:"keyword"`
}
