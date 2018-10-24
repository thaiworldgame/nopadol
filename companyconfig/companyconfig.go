package companyconfig

type RequestConfigTemplate struct {
	Id              int    `json:"id"`
	ComSysId        int64  `json:"com_sys_id"`
	CompanyName     string `json:"company_name"`
	NameEng         string `json:"name_eng"`
	Address         string `json:"address"`
	Telephone       string `json:"telephone"`
	Fax             string `json:"fax"`
	TaxNumber       string `json:"tax_number"`
	Email           string `json:"email"`
	WebSite         string `json:"web_site"`
	TaxRate         int    `json:"tax_rate"`
	BranchName      string `json:"branch_name"`
	BranchAddress   string `json:"branch_address"`
	BranchTelephone string `json:"branch_telephone"`
	BranchFax       string `json:"branch_fax"`
	StockStatus     int    `json:"stock_status"`
	SaleTaxType     int    `json:"sale_tax_type"`
	BuyTaxType      int    `json:"buy_tax_type"`
	DefSaleWh       string `json:"def_sale_wh"`
	DefSaleShelf    string `json:"def_sale_shelf"`
	DefBuyWh        string `json:"def_buy_wh"`
	DefBuyShelf     string `json:"def_buy_shelf"`
	LogoPath        string `json:"logo_path"`
	ActiveStatus    int    `json:"active_status"`
	CreateBy        string `json:"create_by"`
	CreateTime      string `json:"create_time"`
	EditBy          string `json:"edit_by"`
	EditTime        string `json:"edit_time"`
}

type SearchByIdRequestTemplate struct {
	Id int64 `json:"id"`
}

type SearchByKeywordRequestTemplate struct {
	Keyword string `json:"keyword"`
}
