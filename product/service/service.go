package service

//import (
//	"github.com/mrtomyum/nopadol/product"
//	"context"
//)
//
//func New(repo product.Repository) product.Service{
//	return &service{repo}
//}
//
//type service struct {
//	repo product.Repository
//}
//
//func (s *service) SearchProductByBarcode(ctx context.Context, req *product.SearchByBarcodeTemplate) (resp product.ProductTemplate, err error){
//	return s.repo.SearchProductByBarcode(ctx, req)
//}