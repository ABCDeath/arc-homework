package rotatable

import (
	"errors"

	"arc-homework/space-game/moving/object"
)

const (
	DirectionPropName       = "direction"
	DirectionsNumPropName   = "directions_num"
	AngularVelocityPropName = "angular_velocity"
)

var (
	ErrNotRotatable = errors.New("object can not be rotated")
)

type Rotatable interface {
	GetAngle() (int, error)
	SetAngle(v int) error
	GetAngularVelocity() (int, error)
	SetAngularVelocity(v int) error
}

type adapter struct {
	obj object.Object
}
