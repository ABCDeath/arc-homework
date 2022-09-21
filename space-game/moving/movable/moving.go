package movable

import (
	"errors"
	"math"

	"arc-homework/space-game/moving/object"
	"arc-homework/space-game/moving/vector"
)

const (
	PositionPropName      = "position"
	DirectionPropName     = "direction"
	DirectionsNumPropName = "directions_num"
	VelocityPropName      = "velocity"
)

var (
	ErrNotMovable = errors.New("object is not movable")
)

type Movable interface {
	GetPosition() (vector.Vector, error)
	SetPosition(v vector.Vector) error
	GetVelocity() (vector.Vector, error)
	SetVelocity(v vector.Vector) error
}

type adapter struct {
	obj object.Object
}

func (a *adapter) GetPosition() (vector.Vector, error) {
	position, err := getProperty[vector.Vector](a.obj, PositionPropName)
	if err != nil {
		return vector.Vector{}, err
	}

	return *position, nil
}

func (a *adapter) SetPosition(v vector.Vector) error {
	err := a.obj.SetProperty(PositionPropName, &v)
	if err != nil {
		if errors.Is(err, object.ErrNoProperty) {
			return ErrNotMovable
		}

		return err
	}

	return nil
}

func (a *adapter) GetVelocity() (vector.Vector, error) {
	direction, err := getProperty[int](a.obj, DirectionPropName)
	if err != nil {
		return vector.Vector{}, err
	}

	directionsNum, err := getProperty[int](a.obj, DirectionsNumPropName)
	if err != nil {
		return vector.Vector{}, err
	}

	velocity, err := getProperty[int](a.obj, VelocityPropName)
	if err != nil {
		return vector.Vector{}, err
	}

	angle := a.calcDirectionRad(*direction, *directionsNum)
	x := float64(*velocity) * math.Cos(angle)
	y := float64(*velocity) * math.Sin(angle)
	v := vector.New(int(x), int(y))

	return v, nil
}

func (a *adapter) calcDirectionRad(direction, directionsNum int) float64 {
	return float64(direction) * math.Pi / float64(directionsNum)
}

func (a *adapter) SetVelocity(v vector.Vector) error {
	err := a.obj.SetProperty(VelocityPropName, &v)
	if err != nil {
		if errors.Is(err, object.ErrNoProperty) {
			return ErrNotMovable
		}

		return err
	}

	return nil
}

func getProperty[T any](obj object.Object, name string) (*T, error) {
	propPtr, err := object.GetObjectProperty[T](obj, name)
	if err != nil {
		if errors.Is(err, object.ErrNoProperty) {
			return nil, ErrNotMovable
		}

		return nil, err
	}

	return propPtr, nil
}

func New(obj object.Object) Movable {
	return &adapter{obj: obj}
}
