package print

type PosSlipRequestTemplate struct {
	DocNo string `json:"doc_no"`
}

type PosDriveThruSlipRequestTemplate struct {
	//DbHost string `json:"db_host"`
	//DbName string `json:"db_name"`
	//DbUser string `json:"db_user"`
	//DbPass string `json:"db_pass"`
	//HostIP string `json:"host_ip"`

	AccessToken string `json:"access_token"`
	FormId string `json:"form_id"`
	PosId string `json:"pos_id"`
	DocNo string `json:"doc_no"`
}

type PosSlipResponseTemplate struct {
	DocNo string `json:"doc_no"`
}
