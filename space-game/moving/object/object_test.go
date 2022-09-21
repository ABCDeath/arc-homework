package object

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"arc-homework/space-game/moving/object/mocks"
)

func TestGetObjectProperty(t *testing.T) {
	propName := "some_property"

	t.Run("error if Object.GetProperty() returns error for", func(t *testing.T) {
		obj := mocks.Object{}
		obj.On("GetProperty", propName).Return(mock.Anything, ErrNoProperty).Once()

		_, err := GetObjectProperty[int](&obj, propName)
		assert.ErrorIs(t, err, ErrNoProperty)
		obj.AssertExpectations(t)
	})

	t.Run("error if can not cast Object.GetProperty() returned value", func(t *testing.T) {
		obj := mocks.Object{}
		obj.On("GetProperty", propName).Return(struct{}{}, nil).Once()

		_, err := GetObjectProperty[int](&obj, propName)
		assert.ErrorIs(t, err, ErrPropertyType)
		obj.AssertExpectations(t)
	})

	t.Run("no error", func(t *testing.T) {
		expectedValue := 42
		obj := mocks.Object{}
		obj.On("GetProperty", propName).Return(&expectedValue, nil).Once()

		actualValue, err := GetObjectProperty[int](&obj, propName)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, *actualValue)
		obj.AssertExpectations(t)
	})
}
