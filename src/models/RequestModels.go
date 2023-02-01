package models

type IRequestData interface {
	Read([]byte) (int, error)
}
