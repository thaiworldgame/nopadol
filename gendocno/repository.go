package gendocno

type Repository interface {
	Gen(req *DocNoTemplate) (string, error)
}
