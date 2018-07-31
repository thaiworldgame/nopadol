package print

type PosSlipRequestTemplate struct {
	DocNo string `json:"doc_no"`
}

type PosSlipResponseTemplate struct {
	DocNo string `json:"doc_no"`
}
