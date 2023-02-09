package models

type IRequestData interface {
	Read([]byte) (int, error)
}

type IndexResponseModel struct {
	Auth  string `json:"auth"`
	Error string `json:"error"`
}
