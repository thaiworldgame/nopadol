package incentive

// Entity1 type
type Entity1 struct {
	ID     string
	Field1 string
	Field2 Entity2
	Field3 int
}

// Entity2 type
type Entity2 struct {
	Field1 string
	Field2 bool
}

type EntitySearch struct {
	Keyword string `json:"keyword"`
}