package direction

import "errors"

var ErrNoDirection = errors.New("object does not have direction")

type Direction interface {
	GetDirection() (int, error)
	GetDirectionsNum() (int, error)
	SetDirection(int) error
	GetAngularVelocity() (int, error)
}
