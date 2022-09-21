package moving

import (
	"arc-homework/space-game/moving/rotatable"
)

const (
	Degrees = 360
)

type Rotate struct{}

func (r Rotate) Execute(obj rotatable.Rotatable) error {
	directionAngle, err := obj.GetAngle()
	if err != nil {
		return err
	}

	angularVelocity, err := obj.GetAngularVelocity()
	if err != nil {
		return err
	}

	newAngle := (directionAngle + angularVelocity) % Degrees
	if newAngle < 0 {
		newAngle = Degrees + newAngle
	}

	return obj.SetAngle(newAngle)
}
