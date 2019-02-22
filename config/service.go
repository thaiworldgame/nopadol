package config

type SearchByKeywordTemplate struct {
	Keyword string `json:"keyword"`
}

type SearchByIdTemplate struct {
	Id int64 `json:"id"`
}

type SettingTemplate struct {
	Id             int64  `json:"company_id"`
	BranchId       int64  `json:"branch_id"`
	TaxType        int64  `json:"tax_type"`
	TaxRate        int64  `json:"tax_rate"`
	LogoPath       string `json:"logo_path"`
	DepartId       int64  `json:"depart_id"`
	DefSaleWhId    int64  `json:"def_sale_wh_id"`
	DefSaleShelfId int64  `json:"def_sale_shelf_id"`
	DefBuyWhId     int64  `json:"def_buy_wh_id"`
	DefBuyShelfId  int64  `json:"def_buy_shelf_id"`
	StockStatus    int64  `json:"stock_status"`
	SaleTaxType    int64  `json:"sale_tax_type"`
	BuyTaxType     int64  `json:"buy_tax_type"`
	SaleBillType   int64  `json:"sale_bill_type"`
	BuyBillType    int64  `json:"buy_bill_type"`
	UseAddress     string `json:"use_address"`
	PosDefCustId   int64  `json:"pos_def_cust_id"`
	PosDefStock    int64  `json:"pos_def_stock"`
	DefCustId      int64  `json:"def_cust_id"`
	CreateBy       string `json:"create_by"`
	CreateTime     string `json:"create_time"`
	EditBy         string `json:"edit_by"`
	EditTime       string `json:"edit_time"`
}

func New(repo Repository) Service {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	SettingSys(req *SearchByIdTemplate) (interface{}, error)
}

func (s *service) SettingSys(req *SearchByIdTemplate) (interface{}, error) {
	resp, err := s.repo.SettingSys(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
