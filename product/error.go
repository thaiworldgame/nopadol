package product

import "errors"

var (
	StatusNotFound = errors.New("sql: no rows in result set")
	ErrMethodNotAllowed = errors.New("Customer: Data not found")
	ErrForbidden = errors.New("Customer: Data not found")
)

