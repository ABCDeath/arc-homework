package worker

import (
	"context"

	"arc-homework/space-game/command"
)

type hardStopCommand struct {
	cancelWorkerContext context.CancelFunc
}

func (c *hardStopCommand) Execute(_ context.Context) error {
	c.cancelWorkerContext()

	return nil
}

func NewHardStopCommand(cancelFunc context.CancelFunc) command.Command {
	return &hardStopCommand{
		cancelWorkerContext: cancelFunc,
	}
}
