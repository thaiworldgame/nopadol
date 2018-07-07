package employee

import "errors"

var (
	StatusNotFound = errors.New("sql: no rows in result set")
	ErrMethodNotAllowed = errors.New("customer: data not found")
	ErrForbidden = errors.New("mssql: Error converting data type varchar to bigint.")
)
