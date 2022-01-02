package share

// Parameters data structure
type Parameters struct {
	Query   string
	Limit   int
	Offset  int
	Sort    string
	OrderBy string
	All     bool
	Search  string
	GroupBy string
	Filter  interface{}
}

// ResultRepository struct
type ResultRepository struct {
	Result interface{}
	Count  int
	Error  error
}

// List Result
type ResponseList struct {
	IsSuccess bool        `json:"is_success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Meta      *Meta       `json:"meta"`
}

type ResponseDetail struct {
	IsSuccess bool        `json:"is_success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
}

// Meta data structure
type Meta struct {
	Offset int64 `json:"offest"`
	Limit  int64 `json:"limit"`
	Count  int64 `json:"count"`
}
