package worker

import (
	"context"
	"errors"
	"log"
	"sync"

	"arc-homework/space-game/command"
	"arc-homework/space-game/queue"
)

type startCommand struct {
	wg     *sync.WaitGroup
	worker Worker
}

func (c *startCommand) Execute(ctx context.Context) error {
	c.wg.Add(1)

	go func() {
		err := c.worker.Run(ctx)
		if err != nil && !errors.Is(err, context.Canceled) {
			log.Fatalln(err)
		}
	}()

	return nil
}

func NewStartCommand(
	wg *sync.WaitGroup,
	commandQueue queue.Queue[command.Command],
	cmdErrorHandler func(err error),
	stopIfQueueEmpty SoftStopSignal,
) command.Command {
	w := New(commandQueue, cmdErrorHandler, wg, stopIfQueueEmpty)

	return &startCommand{
		wg:     wg,
		worker: w,
	}
}
