package auth

import "errors"

// Errors
var (
	ErrTokenNotFound = errors.New("auth: token not found")
	ErrTokenExpired  = errors.New("auth: token expired")
)

// Token domain-model
type Token struct {
	ID        string
	CompanyID int
	UserID    int64
	BranchID  int64
	ZoneID    string
	TokenID   string
	UserName  string
	UserCode  string
}

// Account domain-model
type Account struct {
	ID           int64
	CompanyID    int64
	FullName     string
	Email        string
	PasswordHash string
	Status       bool
}

// Access domain-model
type Access struct {
	TokenID   string
	AccountID int64
	ClientID  int64
	// VendingID   int64
	// Meta        string
	// VendingUUID string
}

// AccountToken domain-model
type AccountToken struct {
	Token        string
	AccountID    int64
	FullName     string
	Email        string
	IsRegistered bool
	MQTTKey      string
	MQTTChannel  string
}

// Profile domain model
type Profile struct {
	AccountID  int64
	FullName   string
	Email      string
	PictureURL *string
	Status     int
}

// Repository is the auth storage
type Repository interface {
	GetToken(tokenID string) (*Token, error)
	//GetAccount(email string) (*Account, error)
	//SaveAccess(access *Access) error
	//SaveDeviceToken(accountID int64, token string) error
	//DeleteAccess(token string) error
	//GetClientMQTTKey(clientID int64) (*string, error)
	//GetProfile(token string) (*Profile, error)
}
