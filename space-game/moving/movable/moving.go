package movable

import (
	"context"
	"errors"

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

//go:generate python $PROJECT_ROOT/codegen/main.py --object-type=adapter --dependency-module=operations.moving.movable
type Movable interface {
	GetPosition(ctx context.Context) (vector.Vector, error)
	SetPosition(ctx context.Context, v vector.Vector) (vector.Vector, error)
	GetVelocity(ctx context.Context) (vector.Vector, error)
	SetVelocity(ctx context.Context, v vector.Vector) (vector.Vector, error)
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
