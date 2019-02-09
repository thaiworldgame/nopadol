package auth

import (
	"context"
	//"strconv"
	//"strings"
	//"github.com/gofrs/uuid"
	//"golang.org/x/crypto/bcrypt"
	"fmt"
)

type (
	keyToken struct{}
)

// Service is the auth service
type Service interface {
	GetToken(tokenID string) (*Token, error)
	//Signin(email, password string, fcmtoken *string) (*AccountToken, error)
	//Signout(xat string) error
	//Profile(xat string) (*Profile, error)
}

// NewService creates new auth service
func NewService(auths Repository) (Service, error) {
	s := service{auths}
	return &s, nil
}

// GetClientID returns client id from context
func GetCompanyID(ctx context.Context) int {
	x, ok := ctx.Value(keyToken{}).(*Token)
	if !ok {
		return -1
	}
	return x.CompanyID
}

// GetAccountID returns account id from context
func GetUserID(ctx context.Context) int64 {
	x, ok := ctx.Value(keyToken{}).(*Token)
	if !ok {
		return -1
	}
	return x.UserID
}

// GetAccountID returns account id from context
func GetUserCode(ctx context.Context) string {
	x, ok := ctx.Value(keyToken{}).(*Token)
	if !ok {
		return ""
	}
	return x.UserCode
}

// GetVendingID returns vending id from context
func GetBranchID(ctx context.Context) int64 {
	x, ok := ctx.Value(keyToken{}).(*Token)
	if !ok {
		return -1
	}
	return x.BranchID
}

// GetTokenID return access token
func GetTokenID(ctx context.Context) string {
	x, ok := ctx.Value(keyToken{}).(*Token)
	if !ok {
		return ""
	}
	return x.TokenID
}


type service struct {
	auths Repository
}

func (s *service) GetToken(tokenID string) (*Token, error) {
	tk, err := s.auths.GetToken(tokenID)
	fmt.Println(tk)
	if err != nil {
		return nil, err
	}
	return tk, nil
}
