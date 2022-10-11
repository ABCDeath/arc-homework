package moving

import (
	"arc-homework/space-game/command"
	"arc-homework/space-game/moving/rotatable"
)

const (
	Degrees = 360
)

type Rotate struct {
	obj rotatable.Rotatable
}

func (r *Rotate) Execute() error {
	directionAngle, err := r.obj.GetAngle()
	if err != nil {
		return err
	}

	angularVelocity, err := r.obj.GetAngularVelocity()
	if err != nil {
		return err
	}

	newAngle := (directionAngle + angularVelocity) % Degrees
	if newAngle < 0 {
		newAngle = Degrees + newAngle
	}

	return r.obj.SetAngle(newAngle)
}

func NewRotate(obj rotatable.Rotatable) *Rotate {
	return &Rotate{
		obj: obj,
	}
}

type RotateAndChangeVelocity struct {
	cmd command.Command
}

func (r *RotateAndChangeVelocity) Execute() error {
	return r.cmd.Execute()
}

func NewRotateAndChangeVelocity(rotate *Rotate, changeVelocity *ChangeVelocity) *RotateAndChangeVelocity {
	macro := command.NewMacroCommand(rotate, changeVelocity)

	return &RotateAndChangeVelocity{
		cmd: macro,
	}
}
