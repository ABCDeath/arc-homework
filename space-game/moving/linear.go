package moving

import (
	"context"

	"arc-homework/space-game/command"
	"arc-homework/space-game/moving/movable"
)

type Move struct {
	obj movable.Movable
}

func (m *Move) Execute(ctx context.Context) error {
	position, err := m.obj.GetPosition(ctx)
	if err != nil {
		return err
	}

	velocity, err := m.obj.GetVelocity(ctx)
	if err != nil {
		return err
	}

	_, err = m.obj.SetPosition(ctx, position.Add(velocity))

	return err
}

func NewMove(obj movable.Movable) *Move {
	return &Move{
		obj: obj,
	}
}

type MoveAndBurnFuel struct {
	cmd command.Command
}

func (m *MoveAndBurnFuel) Execute(ctx context.Context) error {
	return m.cmd.Execute(ctx)
}

func NewMoveAndBurnFuel(check *CheckFuel, move *Move, burn *BurnFuel) *MoveAndBurnFuel {
	macro := command.NewMacroCommand(check, move, burn)

	return &MoveAndBurnFuel{
		cmd: macro,
	}
}
