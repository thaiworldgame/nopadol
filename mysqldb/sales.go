package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/sales"
)

type NewQuoModel struct {
	Id                  int64             `json:"id"`
	DocNo               string            `json:"doc_no"`
	DocDate             string            `json:"doc_date"`
	ArId                int64             `json:"ar_id"`
	ArCode              string            `json:"ar_code"`
	ArName              string            `json:"ar_name"`
	SaleId              int               `json:"sale_id"`
	SaleCode            string            `json:"sale_code"`
	SaleName            string            `json:"sale_name"`
	BillType            int64             `json:"bill_type"`
	TaxType             int64             `json:"tax_type"`
	TaxRate             int64             `json:"tax_rate"`
	DepartCode          string            `json:"depart_code"`
	RefNo               string            `json:"ref_no"`
	IsConfirm           int64             `json:"is_confirm"`
	BillStatus          int64             `json:"bill_status"`
	DueDate             string            `json:"due_date"`
	ExpireDate          string            `json:"expire_date"`
	DeliveryDate        string            `json:"delivery_date"`
	AssertStatus        int64             `json:"assert_status"`
	IsConditionSend     int64             `json:"is_condition_send"`
	MyDescription       string            `json:"my_description"`
	SumItemAmount       float64           `json:"sum_item_amount"`
	DiscountWord        string            `json:"discount_word"`
	DiscountAmount      float64           `json:"discount_amount"`
	AfterDiscountAmount float64           `json:"after_discount"`
	BeforeTaxAmount     float64           `json:"before_tax_amount"`
	TaxAmount           float64           `json:"tax_amount"`
	TotalAmount         float64           `json:"total_amount"`
	NetAmount           float64           `json:"net_debt_amount"`
	ProjectId           int64             `json:"project_id"`
	ProjectCode         string            `json:"project_code"`
	IsCancel            int64             `json:"is_cancel"`
	CreateBy            string            `json:"creator_by"`
	CreateTime          string            `json:"create_time"`
	EditBy              string            `json:"edit_by"`
	EditTime            string            `json:"edit_time"`
	CancelBy            string            `json:"cancel_by"`
	CancelTime          string            `json:"cancel_time"`
	Subs                []NewQuoItemModel `json:"subs"`
}

type NewQuoItemModel struct {
	IdSub           int64   `json:"id_sub"`
	QTId            int64   `json:"qt_id"`
	ItemId          int64   `json:"item_id"`
	ItemCode        string  `json:"item_code"`
	BarCode         string  `json:"bar_code"`
	ItemName        string  `json:"item_name"`
	Qty             float64 `json:"qty"`
	RemainQty       float64 `json:"remain_qty"`
	Price           float64 `json:"price"`
	DiscountWord    string  `json:"discount_word"`
	DiscountAmount  float64 `json:"discount_amount"`
	UnitCode        string  `json:"unit_code"`
	ItemAmount      float64 `json:"item_amount"`
	ItemDescription string  `json:"item_description"`
	PackingRate1    float64 `json:"packing_rate_1"`
	LineNumber      int     `json:"line_number"`
}

type salesRepository struct{ db *sqlx.DB }

func NewSalesRepository(db *sqlx.DB) sales.Repository {
	return &salesRepository{db}
}

func (repo *salesRepository) Create(req *sales.NewQuoTemplate) (resp interface{}, err error) {
	return nil, nil
}
