package moving

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"arc-homework/space-game/moving/rotatable"
	"arc-homework/space-game/moving/rotatable/mocks"
)

func TestRotate_Execute(t *testing.T) {
	t.Run("error if GetAngle() returns error", func(t *testing.T) {
		rotatableObj := mocks.Rotatable{}
		rotatableObj.On("GetAngle").Return(0, rotatable.ErrNotRotatable).Once()

		err := Rotate{}.Execute(&rotatableObj)
		assert.ErrorIs(t, err, rotatable.ErrNotRotatable)
		rotatableObj.AssertExpectations(t)
	})

	t.Run("error if GetAngularVelocity() returns error", func(t *testing.T) {
		rotatableObj := mocks.Rotatable{}
		rotatableObj.On("GetAngle").Return(0, nil).Once()
		rotatableObj.On("GetAngularVelocity").Return(0, rotatable.ErrNotRotatable).Once()

		err := Rotate{}.Execute(&rotatableObj)
		assert.ErrorIs(t, err, rotatable.ErrNotRotatable)
		rotatableObj.AssertExpectations(t)
	})

	t.Run("error if SetAngle() returns error", func(t *testing.T) {
		rotatableObj := mocks.Rotatable{}
		rotatableObj.On("GetAngle").Return(0, nil).Once()
		rotatableObj.On("GetAngularVelocity").Return(0, nil).Once()
		rotatableObj.On("SetAngle", 0).Return(rotatable.ErrNotRotatable).Once()

		err := Rotate{}.Execute(&rotatableObj)
		assert.ErrorIs(t, err, rotatable.ErrNotRotatable)
		rotatableObj.AssertExpectations(t)
	})

	t.Run("rotates obj with 45 deg direction and -90 ang vel to 315 deg", func(t *testing.T) {
		rotatableObj := mocks.Rotatable{}
		rotatableObj.On("GetAngle").Return(45, nil).Once()
		rotatableObj.On("GetAngularVelocity").Return(-90, nil).Once()
		rotatableObj.On("SetAngle", 315).Return(nil).Once()

		err := Rotate{}.Execute(&rotatableObj)
		assert.NoError(t, err)
		rotatableObj.AssertExpectations(t)
	})

	t.Run("rotates obj with 315 deg direction and 90 ang vel to 45 deg", func(t *testing.T) {
		rotatableObj := mocks.Rotatable{}
		rotatableObj.On("GetAngle").Return(315, nil).Once()
		rotatableObj.On("GetAngularVelocity").Return(90, nil).Once()
		rotatableObj.On("SetAngle", 45).Return(nil).Once()

		err := Rotate{}.Execute(&rotatableObj)
		assert.NoError(t, err)
		rotatableObj.AssertExpectations(t)
	})
}
