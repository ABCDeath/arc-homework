package moving

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	cmdmock "arc-homework/space-game/command/mocks"
	"arc-homework/space-game/moving/movable"
	"arc-homework/space-game/moving/movable/mocks"
	"arc-homework/space-game/moving/vector"
)

func TestMove_Execute(t *testing.T) {
	t.Run("error if GetPosition() returns error", func(t *testing.T) {
		movableObj := mocks.Movable{}
		movableObj.On("GetPosition", context.Background()).Return(vector.Vector{}, movable.ErrNotMovable).Once()

		err := NewMove(&movableObj).Execute(context.Background())
		assert.ErrorIs(t, err, movable.ErrNotMovable)
		movableObj.AssertExpectations(t)
	})

	t.Run("error if GetVelocity() returns error", func(t *testing.T) {
		movableObj := mocks.Movable{}
		movableObj.On("GetPosition", context.Background()).Return(vector.New(1, 2), nil).Once()
		movableObj.On("GetVelocity", context.Background()).Return(vector.Vector{}, movable.ErrNotMovable).Once()

		err := NewMove(&movableObj).Execute(context.Background())
		assert.ErrorIs(t, err, movable.ErrNotMovable)
		movableObj.AssertExpectations(t)
	})

	t.Run("error if SetPosition() returns error", func(t *testing.T) {
		movableObj := mocks.Movable{}
		movableObj.On("GetPosition", context.Background()).Return(vector.New(0, 0), nil).Once()
		movableObj.On("GetVelocity", context.Background()).Return(vector.New(0, 0), nil).Once()
		movableObj.On("SetPosition", context.Background(), vector.New(0, 0)).Return(vector.Vector{}, movable.ErrNotMovable).Once()

		err := NewMove(&movableObj).Execute(context.Background())
		assert.ErrorIs(t, err, movable.ErrNotMovable)
		movableObj.AssertExpectations(t)
	})

	t.Run("moves obj(12, 5) v(-7, 3) to (5, 8)", func(t *testing.T) {
		movableObj := mocks.Movable{}
		movableObj.On("GetPosition", context.Background()).Return(vector.New(12, 5), nil).Once()
		movableObj.On("GetVelocity", context.Background()).Return(vector.New(-7, 3), nil).Once()
		movableObj.On("SetPosition", context.Background(), vector.New(5, 8)).Return(vector.New(5, 8), nil).Once()

		err := NewMove(&movableObj).Execute(context.Background())
		assert.NoError(t, err)
		movableObj.AssertExpectations(t)
	})
}

func TestMoveAndBurnFuel_Execute(t *testing.T) {
	t.Run("error if underlying command returns error", func(t *testing.T) {
		expectedErr := errors.New("")
		cmd := cmdmock.Command{}
		cmd.On("Execute", context.Background()).Return(expectedErr).Once()

		command := MoveAndBurnFuel{cmd: &cmd}

		err := command.Execute(context.Background())
		assert.ErrorIs(t, err, expectedErr)
		cmd.AssertExpectations(t)
	})

	t.Run("executes underlying command", func(t *testing.T) {
		cmd := cmdmock.Command{}
		cmd.On("Execute", context.Background()).Return(nil).Once()

		command := MoveAndBurnFuel{cmd: &cmd}

		err := command.Execute(context.Background())
		assert.NoError(t, err)
		cmd.AssertExpectations(t)
	})
}
