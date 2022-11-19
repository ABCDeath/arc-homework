package game_runner

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"arc-homework/space-game/command"
	"arc-homework/space-game/command/mocks"
	errhandler "arc-homework/space-game/error_handler"
)

func Test_configureErrorHandler(t *testing.T) {
	t.Run("put retry command into queue if Execute returns error", func(t *testing.T) {
		errHandler := errhandler.New()
		cmdQueue := NewQueue(errHandler)

		configureErrorHandler(cmdQueue, errHandler)

		cmd := mocks.Command{}
		cmd.On("Execute", context.Background()).Return(errors.New("")).Once()

		cmdQueue.Enqueue(&cmd)
		cmdQueue.dequeAndRun()

		require.False(t, cmdQueue.IsEmpty())
		actualCmdIface := cmdQueue.Dequeue()
		_, ok := actualCmdIface.(*command.RetryCommand)
		require.True(t, ok)
	})

	t.Run("retry command if Execute fails", func(t *testing.T) {
		errHandler := errhandler.New()
		cmdQueue := NewQueue(errHandler)

		configureErrorHandler(cmdQueue, errHandler)

		cmd := mocks.Command{}
		cmd.On("Execute", context.Background()).Return(errors.New("")).Once()
		cmd.On("Execute", context.Background()).Return(nil).Once()

		cmdQueue.Enqueue(&cmd)
		cmdQueue.Run()

		require.True(t, cmdQueue.IsEmpty())
	})

	t.Run("put retry2 command into queue if Execute returns error twice", func(t *testing.T) {
		errHandler := errhandler.New()
		cmdQueue := NewQueue(errHandler)

		configureErrorHandler(cmdQueue, errHandler)

		cmd := mocks.Command{}
		cmd.On("Execute", context.Background()).Return(errors.New("")).Twice()

		cmdQueue.Enqueue(&cmd)
		cmdQueue.dequeAndRun()
		cmdQueue.dequeAndRun()

		require.False(t, cmdQueue.IsEmpty())
		actualCmdIface := cmdQueue.Dequeue()
		_, ok := actualCmdIface.(*command.Retry2Command)
		require.True(t, ok)
	})

	t.Run("retry twice command if Execute fails", func(t *testing.T) {
		errHandler := errhandler.New()
		cmdQueue := NewQueue(errHandler)

		configureErrorHandler(cmdQueue, errHandler)

		cmd := mocks.Command{}
		cmd.On("Execute", context.Background()).Return(errors.New("")).Twice()
		cmd.On("Execute", context.Background()).Return(nil).Once()

		cmdQueue.Enqueue(&cmd)
		cmdQueue.Run()

		require.True(t, cmdQueue.IsEmpty())
	})

	t.Run("put log command into queue if Execute fails 3 times", func(t *testing.T) {
		errHandler := errhandler.New()
		cmdQueue := NewQueue(errHandler)

		configureErrorHandler(cmdQueue, errHandler)

		cmd := mocks.Command{}
		cmd.On("Execute", context.Background()).Return(errors.New("")).Times(3)

		cmdQueue.Enqueue(&cmd)
		cmdQueue.dequeAndRun()
		cmdQueue.dequeAndRun()
		cmdQueue.dequeAndRun()

		require.False(t, cmdQueue.IsEmpty())
		actualCmdIface := cmdQueue.Dequeue()
		_, ok := actualCmdIface.(*command.LogErrorCommand)
		require.True(t, ok)
	})
}
