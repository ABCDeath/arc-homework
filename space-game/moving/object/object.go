package object

import "errors"

var (
	ErrNoProperty = errors.New("object does not have property")
)

type Object interface {
	GetProperty(name string) (interface{}, error)
	SetProperty(name string, value interface{}) error
}
