package game_runner

import (
	"errors"
	"testing"

	cmdmock "arc-homework/space-game/command/mocks"
	"arc-homework/space-game/error_handler/mocks"
)

func TestQueue_Run(t *testing.T) {
	t.Run("deque command and execute it", func(t *testing.T) {
		errHandler := mocks.Handler{}

		cmd := cmdmock.Command{}
		cmd.On("Execute").Return(nil).Once()

		q := NewQueue(&errHandler)
		q.Enqueue(&cmd)

		q.Run()

		cmd.AssertExpectations(t)
	})

	t.Run("call error handler if command.Execute() returns error", func(t *testing.T) {
		err := errors.New("")
		cmd := cmdmock.Command{}
		cmd.On("Execute").Return(err).Once()

		errHandler := mocks.Handler{}
		errHandler.On("Handle", &cmd, err).Return(nil).Once()

		q := NewQueue(&errHandler)
		q.Enqueue(&cmd)

		q.Run()

		cmd.AssertExpectations(t)
		errHandler.AssertExpectations(t)
	})
}
