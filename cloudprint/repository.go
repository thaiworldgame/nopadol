package cloudprint

type Repository interface {
	CloudPrint(req *CloudPrintRequest) (interface{}, error)
}
