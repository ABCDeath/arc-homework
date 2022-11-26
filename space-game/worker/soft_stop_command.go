package worker

import (
	"context"

	"arc-homework/space-game/command"
)

type softStopCommand struct {
	softStopSignal SoftStopSignal
}

func (c *softStopCommand) Execute(_ context.Context) error {
	close(c.softStopSignal)

	return nil
}

func NewSoftStopCommand(softStopSignal SoftStopSignal) command.Command {
	return &softStopCommand{
		softStopSignal: softStopSignal,
	}
}
