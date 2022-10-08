package game_runner

import (
	"log"
	"reflect"

	"arc-homework/space-game/command"
	errhandler "arc-homework/space-game/error_handler"
)

func InitGame() *Queue {
	errHandler := errhandler.New()
	cmdQueue := NewQueue(errHandler)

	configureErrorHandler(cmdQueue, errHandler)

	return cmdQueue
}

func configureErrorHandler(cmdQueue *Queue, handler errhandler.Handler) {
	logger := log.Logger{}

	// any command should be retried once
	handler.RegisterDefaultHandler(func(cmd command.Command, err error) error {
		retry := command.NewRetryCommand(cmd)
		cmdQueue.Enqueue(retry)

		return nil
	})

	handler.RegisterDefaultCmdHandler(reflect.TypeOf(command.RetryCommand{}),
		func(cmd command.Command, err error) error {
			retry := command.NewRetry2Command(cmd)
			cmdQueue.Enqueue(retry)

			return nil
		},
	)

	handler.RegisterDefaultCmdHandler(reflect.TypeOf(command.Retry2Command{}),
		func(cmd command.Command, err error) error {
			l := command.NewLogErrorCommand(&logger, cmd, err)
			cmdQueue.Enqueue(l)

			return nil
		},
	)
}
