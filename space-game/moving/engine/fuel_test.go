package engine

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
		FuelAmountPropName,
		FuelBurnRatePropName,
	}

	for _, propName := range propertyNames {
		t.Run(fmt.Sprintf("error if Object.GetProperty() returns error for %s property", propName), func(t *testing.T) {
			obj := mocks.Object{}
			obj.On("GetProperty", propName).Return(mock.Anything, object.ErrNoProperty).Once()

			_, err := getProperty[int](&obj, propName)
			assert.ErrorIs(t, err, ErrFuelIsNotSupported)
			obj.AssertExpectations(t)
		})
	}

	t.Run("no error", func(t *testing.T) {
		expectedValue := 42
		obj := mocks.Object{}
		obj.On("GetProperty", FuelAmountPropName).Return(&expectedValue, nil).Once()

		actualValue, err := getProperty[int](&obj, FuelAmountPropName)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, *actualValue)
		obj.AssertExpectations(t)
	})
}

func Test_adapter_GetFuelAmount(t *testing.T) {
	t.Run("error if Object.GetProperty() returns error", func(t *testing.T) {
		obj := mocks.Object{}
		obj.On("GetProperty", FuelAmountPropName).Return(mock.Anything, object.ErrNoProperty).Once()

		adapterObj := New(&obj)

		_, err := adapterObj.GetFuelAmount()
		assert.ErrorIs(t, err, ErrFuelIsNotSupported)
		obj.AssertExpectations(t)
	})

	t.Run("returns int", func(t *testing.T) {
		expected := 1
		obj := mocks.Object{}
		obj.On("GetProperty", FuelAmountPropName).Return(&expected, nil).Once()

		adapterObj := New(&obj)

		actual, err := adapterObj.GetFuelAmount()
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
		obj.AssertExpectations(t)
	})
}

func Test_adapter_GetFuelBurnRate(t *testing.T) {
	t.Run("error if Object.GetProperty() returns error", func(t *testing.T) {
		obj := mocks.Object{}
		obj.On("GetProperty", FuelBurnRatePropName).Return(mock.Anything, object.ErrNoProperty).Once()

		adapterObj := New(&obj)

		_, err := adapterObj.GetFuelBurnRate()
		assert.ErrorIs(t, err, ErrFuelIsNotSupported)
		obj.AssertExpectations(t)
	})

	t.Run("returns int", func(t *testing.T) {
		expected := 1
		obj := mocks.Object{}
		obj.On("GetProperty", FuelBurnRatePropName).Return(&expected, nil).Once()

		adapterObj := New(&obj)

		actual, err := adapterObj.GetFuelBurnRate()
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
		obj.AssertExpectations(t)
	})
}

func Test_adapter_SetFuelAmount(t *testing.T) {
	t.Run("error if Object.GetProperty() returns error", func(t *testing.T) {
		value := 1
		obj := mocks.Object{}
		obj.On("SetProperty", FuelAmountPropName, value).Return(object.ErrNoProperty).Once()

		adapterObj := New(&obj)

		err := adapterObj.SetFuelAmount(value)
		assert.ErrorIs(t, err, ErrFuelIsNotSupported)
		obj.AssertExpectations(t)
	})

	t.Run("returns int", func(t *testing.T) {
		value := 1
		obj := mocks.Object{}
		obj.On("SetProperty", FuelAmountPropName, value).Return(nil).Once()

		adapterObj := New(&obj)

		err := adapterObj.SetFuelAmount(value)
		assert.NoError(t, err)
		obj.AssertExpectations(t)
	})
}
