package config

import (
	"os"
	"fmt"
	"encoding/json"
)

type Default struct {
	TaxRateDefault      float64 `json:"tax_rate_default"`
	ExchangeRateDefault float64 `json:"exchange_rate_default"`
	PosTaxType          int     `json:"pos_tax_type"`
	PosSaveForm         int     `json:"pos_save_form"`
	PosGLFormat         string  `json:"pos_gl_format"`
	PosSource           int     `json:"pos_source"`
	PosMyType           int     `json:"pos_my_type"`
	PosMachineNo        string  `json:"pos_machine_no"`
}

func LoadDefaultData(fileName string) (df Default) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Err Open file %v: Error is:", file, err)
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&df)
	if err != nil {
		fmt.Println("error Decode Json:", err)
	}

	return df
}
