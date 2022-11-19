package moving

import (
	"context"

	"arc-homework/space-game/moving/direction"
)

type ChangeVelocity struct {
	obj direction.Direction
}

func (c *ChangeVelocity) Execute(_ context.Context) error {
	dir, err := c.obj.GetDirection()
	if err != nil {
		return err
	}

	directionsNum, err := c.obj.GetDirectionsNum()
	if err != nil {
		return err
	}

	angularVelocity, err := c.obj.GetAngularVelocity()
	if err != nil {
		return err
	}

	newDirection := (dir + angularVelocity) % directionsNum

	return c.obj.SetDirection(newDirection)
}

func NewChangeVelocity(obj direction.Direction) *ChangeVelocity {
	return &ChangeVelocity{
		obj: obj,
	}
}
