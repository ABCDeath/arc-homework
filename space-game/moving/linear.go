package moving

import (
	"arc-homework/space-game/command"
	"arc-homework/space-game/moving/movable"
)

type Move struct {
	obj movable.Movable
}

func (m *Move) Execute() error {
	position, err := m.obj.GetPosition()
	if err != nil {
		return err
	}

	velocity, err := m.obj.GetVelocity()
	if err != nil {
		return err
	}

	return m.obj.SetPosition(position.Add(velocity))
}

func NewMove(obj movable.Movable) *Move {
	return &Move{
		obj: obj,
	}
}

type MoveAndBurnFuel struct {
	cmd command.Command
}

func (m *MoveAndBurnFuel) Execute() error {
	return m.cmd.Execute()
}

func NewMoveAndBurnFuel(check *CheckFuel, move *Move, burn *BurnFuel) *MoveAndBurnFuel {
	macro := command.NewMacroCommand(check, move, burn)

	return &MoveAndBurnFuel{
		cmd: macro,
	}
}
