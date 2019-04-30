package posconfig

type PosConfigTemplate struct {
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
	Telephone      string `json:"telephone"`
	TaxId          string `json:"tax_id"`
	ArCode         string `json:"ar_code"`
	PosId          string `json:"pos_id"`
	WhCode         string `json:"wh_code"`
	ShelfCode      string `json:"shelf_code"`
	PrinterPosIp   string `json:"printer_pos_ip"`
	PrinterCopyIp  string `json:"printer_copy_ip"`
	MachineNo      string `json:"machine_no"`
	MachineCode    string `json:"machine_code"`
	TaxRate        int64  `json:"tax_rate"`
}
