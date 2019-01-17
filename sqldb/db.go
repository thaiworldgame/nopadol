package sqldb

import (
	//"github.com/mrtomyum/nopadol/incentive"
	"context"
)

// NewDomain1Repository creates domain1 repository implements by domain4
//func NewIncentiveRepository() incentive.Repository {
//	return &incentiveRepository{}
//}

type incentiveRepository struct {}

//func (incentiveRepository) Register(ctx context.Context, entity *incentive.Entity1) (string, error) {
//	return "", nil
//}

// SetField3 sets field3 for Entity1
func (incentiveRepository) SetField3(ctx context.Context, entityID string, field3 int) error {
	return nil
}