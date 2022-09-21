package object

import (
	"errors"
	"fmt"
)

var (
	ErrNoProperty   = errors.New("object does not have property")
	ErrPropertyType = errors.New("property type cast error")
)

type Object interface {
	GetProperty(name string) (interface{}, error)
	SetProperty(name string, value interface{}) error
}

func GetObjectProperty[T any](obj Object, name string) (*T, error) {
	propPtr, err := obj.GetProperty(name)
	if err != nil {
		return nil, err
	}

	propValue, castOk := propPtr.(*T)
	if !castOk {
		return nil, fmt.Errorf("%w: %s", ErrPropertyType, name)
	}

	return propValue, nil
}
