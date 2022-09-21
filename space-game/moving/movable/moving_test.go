package movable

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"arc-homework/space-game/moving/object"
	"arc-homework/space-game/moving/object/mocks"
	"arc-homework/space-game/moving/vector"
)

func Test_getProperty(t *testing.T) {
	propertyNames := []string{
		PositionPropName,
		DirectionPropName,
		DirectionsNumPropName,
		VelocityPropName,
	}

	for _, propName := range propertyNames {
		t.Run(fmt.Sprintf("error if Object.GetProperty() returns error for %s property", propName), func(t *testing.T) {
			obj := mocks.Object{}
			obj.On("GetProperty", propName).Return(mock.Anything, object.ErrNoProperty).Once()

			_, err := getProperty[int](&obj, propName)
			assert.ErrorIs(t, err, ErrNotMovable)
			obj.AssertExpectations(t)
		})
	}

	t.Run("no error", func(t *testing.T) {
		expectedValue := 42
		obj := mocks.Object{}
		obj.On("GetProperty", PositionPropName).Return(&expectedValue, nil).Once()

		actualValue, err := getProperty[int](&obj, PositionPropName)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, *actualValue)
		obj.AssertExpectations(t)
	})
}

func TestAdapter_GetPosition(t *testing.T) {
	t.Run("error if Object.GetProperty() returns error", func(t *testing.T) {
		obj := mocks.Object{}
		obj.On("GetProperty", PositionPropName).Return(mock.Anything, object.ErrNoProperty).Once()

		adapterObj := New(&obj)

		_, err := adapterObj.GetPosition()
		assert.ErrorIs(t, err, ErrNotMovable)
		obj.AssertExpectations(t)
	})

	t.Run("returns vector", func(t *testing.T) {
		expectedPosition := vector.New(1, 2)
		obj := mocks.Object{}
		obj.On("GetProperty", PositionPropName).Return(&expectedPosition, nil).Once()

		adapterObj := New(&obj)

		actualPosition, err := adapterObj.GetPosition()
		assert.NoError(t, err)
		assert.Equal(t, expectedPosition, actualPosition)
		obj.AssertExpectations(t)
	})
}

func TestAdapter_SetPosition(t *testing.T) {
	t.Run("error if Object.SetProperty() returns error", func(t *testing.T) {
		position := vector.New(1, 2)
		obj := mocks.Object{}
		obj.On("SetProperty", PositionPropName, &position).Return(object.ErrNoProperty).Once()

		adapterObj := New(&obj)

		err := adapterObj.SetPosition(position)
		assert.ErrorIs(t, err, ErrNotMovable)
		obj.AssertExpectations(t)
	})

	t.Run("ok", func(t *testing.T) {
		position := vector.New(1, 2)
		obj := mocks.Object{}
		obj.On("SetProperty", PositionPropName, &position).Return(nil).Once()

		adapterObj := New(&obj)

		err := adapterObj.SetPosition(position)
		assert.NoError(t, err)
		obj.AssertExpectations(t)
	})
}

func TestAdapter_GetVelocity(t *testing.T) {
	direction := 2
	directionNum := 4
	velocity := 5

	t.Run(fmt.Sprintf("error if Object.GetProperty() returns error for %s", DirectionPropName), func(t *testing.T) {
		obj := mocks.Object{}
		obj.On("GetProperty", DirectionPropName).Return(mock.Anything, object.ErrNoProperty).Once()

		adapterObj := New(&obj)

		_, err := adapterObj.GetVelocity()
		assert.ErrorIs(t, err, ErrNotMovable)
		obj.AssertExpectations(t)
	})

	t.Run(fmt.Sprintf("error if Object.GetProperty() returns error for %s", DirectionsNumPropName), func(t *testing.T) {
		obj := mocks.Object{}
		obj.On("GetProperty", DirectionPropName).Return(&direction, nil).Once()
		obj.On("GetProperty", DirectionsNumPropName).Return(mock.Anything, object.ErrNoProperty).Once()

		adapterObj := New(&obj)

		_, err := adapterObj.GetVelocity()
		assert.ErrorIs(t, err, ErrNotMovable)
		obj.AssertExpectations(t)
	})

	t.Run(fmt.Sprintf("error if Object.GetProperty() returns error for %s", VelocityPropName), func(t *testing.T) {
		obj := mocks.Object{}
		obj.On("GetProperty", DirectionPropName).Return(&direction, nil).Once()
		obj.On("GetProperty", DirectionsNumPropName).Return(&directionNum, nil).Once()
		obj.On("GetProperty", VelocityPropName).Return(mock.Anything, object.ErrNoProperty).Once()

		adapterObj := New(&obj)

		_, err := adapterObj.GetVelocity()
		assert.ErrorIs(t, err, ErrNotMovable)
		obj.AssertExpectations(t)
	})

	t.Run("returns vector", func(t *testing.T) {
		obj := mocks.Object{}
		obj.On("GetProperty", DirectionPropName).Return(&direction, nil).Once()
		obj.On("GetProperty", DirectionsNumPropName).Return(&directionNum, nil).Once()
		obj.On("GetProperty", VelocityPropName).Return(&velocity, nil).Once()

		adapterObj := New(&obj)

		actualPosition, err := adapterObj.GetVelocity()
		assert.NoError(t, err)

		expectedVelocity := vector.New(0, velocity)
		assert.Equal(t, expectedVelocity, actualPosition)
		obj.AssertExpectations(t)
	})
}

func TestAdapter_SetVelocity(t *testing.T) {
	t.Run("error if Object.SetProperty() returns error", func(t *testing.T) {
		velocity := vector.New(1, 2)
		obj := mocks.Object{}
		obj.On("SetProperty", VelocityPropName, &velocity).Return(object.ErrNoProperty).Once()

		adapterObj := New(&obj)

		err := adapterObj.SetVelocity(velocity)
		assert.ErrorIs(t, err, ErrNotMovable)
		obj.AssertExpectations(t)
	})

	t.Run("ok", func(t *testing.T) {
		velocity := vector.New(1, 2)
		obj := mocks.Object{}
		obj.On("SetProperty", VelocityPropName, &velocity).Return(nil).Once()

		adapterObj := New(&obj)

		err := adapterObj.SetVelocity(velocity)
		assert.NoError(t, err)
		obj.AssertExpectations(t)
	})
}
