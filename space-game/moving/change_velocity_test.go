package moving

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"arc-homework/space-game/moving/direction"
	"arc-homework/space-game/moving/direction/mocks"
)

func TestChangeVelocity_Execute(t *testing.T) {
	t.Run("error if GetDirection() returns error", func(t *testing.T) {
		directionObj := mocks.Direction{}
		directionObj.On("GetDirection").Return(0, direction.ErrNoDirection).Once()

		cmd := NewChangeVelocity(&directionObj)
		err := cmd.Execute()
		assert.ErrorIs(t, err, direction.ErrNoDirection)
		directionObj.AssertExpectations(t)
	})

	t.Run("error if GetDirectionsNum() returns error", func(t *testing.T) {
		directionObj := mocks.Direction{}
		directionObj.On("GetDirection").Return(0, nil).Once()
		directionObj.On("GetDirectionsNum").Return(0, direction.ErrNoDirection).Once()

		cmd := NewChangeVelocity(&directionObj)
		err := cmd.Execute()
		assert.ErrorIs(t, err, direction.ErrNoDirection)
		directionObj.AssertExpectations(t)
	})

	t.Run("error if GetAngularVelocity() returns error", func(t *testing.T) {
		directionObj := mocks.Direction{}
		directionObj.On("GetDirection").Return(0, nil).Once()
		directionObj.On("GetDirectionsNum").Return(0, nil).Once()
		directionObj.On("GetAngularVelocity").Return(0, direction.ErrNoDirection).Once()

		cmd := NewChangeVelocity(&directionObj)
		err := cmd.Execute()
		assert.ErrorIs(t, err, direction.ErrNoDirection)
		directionObj.AssertExpectations(t)
	})

	t.Run("error if SetDirection() returns error", func(t *testing.T) {
		directionObj := mocks.Direction{}
		directionObj.On("GetDirection").Return(0, nil).Once()
		directionObj.On("GetDirectionsNum").Return(1, nil).Once()
		directionObj.On("GetAngularVelocity").Return(0, nil).Once()
		directionObj.On("SetDirection", 0).Return(direction.ErrNoDirection).Once()

		cmd := NewChangeVelocity(&directionObj)
		err := cmd.Execute()
		assert.ErrorIs(t, err, direction.ErrNoDirection)
		directionObj.AssertExpectations(t)
	})

	t.Run("change direction corresponding to angular velocity", func(t *testing.T) {
		directionObj := mocks.Direction{}
		directionObj.On("GetDirection").Return(2, nil).Once()
		directionObj.On("GetDirectionsNum").Return(4, nil).Once()
		directionObj.On("GetAngularVelocity").Return(1, nil).Once()
		directionObj.On("SetDirection", 3).Return(nil).Once()

		cmd := NewChangeVelocity(&directionObj)
		err := cmd.Execute()
		assert.NoError(t, err)
		directionObj.AssertExpectations(t)
	})
}
