package helper

type Meta struct {
	Limit              int
	Offset             int
	CountAllData       int
	CountRetrievedData int
}

type Param struct {
	Limit  int
	Offset int
	Filter interface{}
}
