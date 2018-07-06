package employee

import "errors"

var (
	ErrEmployeeNotFound = errors.New("Employee: Data not found")
	ErrMethodNotAllowed = errors.New("")
	StatusNotFound = errors.New("")
	ErrForbidden = errors.New("")
)
