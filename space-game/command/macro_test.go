package command

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"arc-homework/space-game/command/mocks"
)

func TestMacroCommand_Execute(t *testing.T) {
	t.Run("error if any command fails", func(t *testing.T) {
		cmd1 := mocks.Command{}
		cmd1.On("Execute").Return(ErrCommand).Once()

		cmd2 := mocks.Command{}

		macro := NewMacroCommand(&cmd1, &cmd2)

		err := macro.Execute()
		assert.ErrorIs(t, err, ErrCommand)
		cmd1.AssertExpectations(t)
	})

	t.Run("executes every passed command", func(t *testing.T) {
		cmd1 := mocks.Command{}
		cmd1.On("Execute").Return(nil).Once()

		cmd2 := mocks.Command{}
		cmd2.On("Execute").Return(nil).Once()

		cmd3 := mocks.Command{}
		cmd3.On("Execute").Return(nil).Once()

		macro := NewMacroCommand(&cmd1, &cmd2, &cmd3)

		err := macro.Execute()
		assert.NoError(t, err)
		cmd1.AssertExpectations(t)
		cmd2.AssertExpectations(t)
		cmd3.AssertExpectations(t)
	})
}
