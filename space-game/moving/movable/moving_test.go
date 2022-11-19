package movable

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"arc-homework/space-game/moving/object"
	"arc-homework/space-game/moving/object/mocks"
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
