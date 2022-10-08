package command

import (
	"testing"

	"github.com/stretchr/testify/require"

	"arc-homework/space-game/command/mocks"
)

func TestRetryCommand_Execute(t *testing.T) {
	t.Run("execute executes underlying command", func(t *testing.T) {
		cmd := mocks.Command{}
		cmd.On("Execute").Return(nil).Once()

		retry := NewRetryCommand(&cmd)

		err := retry.Execute()

		require.NoError(t, err)
		cmd.AssertExpectations(t)
	})
}

func TestRetry2Command_Execute(t *testing.T) {
	t.Run("execute executes underlying command", func(t *testing.T) {
		cmd := mocks.Command{}
		cmd.On("Execute").Return(nil).Once()

		retry := NewRetry2Command(&cmd)

		err := retry.Execute()

		require.NoError(t, err)
		cmd.AssertExpectations(t)
	})
}
