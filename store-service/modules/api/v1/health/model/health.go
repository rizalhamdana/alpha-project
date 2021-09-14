package model

type Health struct {
	Status      string `json:"status"`
	Version     string `json:"version"`
	ServiceName string `json:"service_name"`
}
