package worker

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"arc-homework/space-game/command"
	"arc-homework/space-game/queue"
)

var (
	ErrWorker = errors.New("")
	ErrType   = fmt.Errorf("%wCommand type cast error", ErrWorker)
)

type SoftStopSignal chan struct{}

type Worker interface {
	Run(ctx context.Context) error
}

type worker struct {
	wg               *sync.WaitGroup
	commandQueue     queue.Queue[command.Command]
	errHandler       func(err error)
	stopIfQueueEmpty SoftStopSignal
	stopIfEmpty      bool
}

func (w *worker) Run(ctx context.Context) error {
	defer w.wg.Done()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-w.stopIfQueueEmpty:
			w.stopIfEmpty = true
			w.stopIfQueueEmpty = nil
		default:
			err := w.getCommandAndExecute(ctx)
			if err != nil {
				if w.stopIfEmpty && errors.Is(err, queue.ErrQueueEmpty) {
					return nil
				}

				return err
			}
		}
	}
}

func (w *worker) getCommandAndExecute(ctx context.Context) error {
	var cmdIface interface{}
	var err error

	if w.stopIfEmpty {
		cmdIface, err = w.commandQueue.Dequeue()
	} else {
		cmdIface, err = w.commandQueue.DequeueOrWait(ctx)
	}

	if err != nil {
		return err
	}

	cmd, ok := (cmdIface).(command.Command)
	if !ok {
		return ErrType
	}

	err = cmd.Execute(ctx)
	if err != nil {
		w.errHandler(err)
	}

	return nil
}

func New(
	commandQueue queue.Queue[command.Command],
	cmdErrorHandler func(err error),
	wg *sync.WaitGroup,
	stopIfQueueEmpty SoftStopSignal,
) Worker {
	return &worker{
		commandQueue:     commandQueue,
		errHandler:       cmdErrorHandler,
		wg:               wg,
		stopIfQueueEmpty: stopIfQueueEmpty,
	}
}
