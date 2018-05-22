package service

import (
	"github.com/mrtomyum/nopadol/sale"
)

func New(repo sale.Repository) sale.Service {
	return &service{repo}
}