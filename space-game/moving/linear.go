package moving

import "arc-homework/space-game/moving/movable"

type Move struct{}

func (m Move) Execute(obj movable.Movable) error {
	position, err := obj.GetPosition()
	if err != nil {
		return err
	}

	velocity, err := obj.GetVelocity()
	if err != nil {
		return err
	}

	return obj.SetPosition(position.Add(velocity))
}
