package moving

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"arc-homework/space-game/command"
	"arc-homework/space-game/moving/engine"
	"arc-homework/space-game/moving/engine/mocks"
)

func TestCheckFuel_Execute(t *testing.T) {
	t.Run("error if GetFuelAmount() returns error", func(t *testing.T) {
		fuelObj := mocks.Fuel{}
		fuelObj.On("GetFuelAmount").Return(0, engine.ErrFuelIsNotSupported).Once()

		cmd := NewCheckFuel(&fuelObj)
		err := cmd.Execute()
		assert.ErrorIs(t, err, engine.ErrFuelIsNotSupported)
		fuelObj.AssertExpectations(t)
	})

	t.Run("ErrCommand if fuel amount is 0", func(t *testing.T) {
		fuelObj := mocks.Fuel{}
		fuelObj.On("GetFuelAmount").Return(0, nil).Once()

		cmd := NewCheckFuel(&fuelObj)
		err := cmd.Execute()
		assert.ErrorIs(t, err, command.ErrCommand)
		fuelObj.AssertExpectations(t)
	})

	t.Run("ErrCommand if fuel amount is 0", func(t *testing.T) {
		fuelObj := mocks.Fuel{}
		fuelObj.On("GetFuelAmount").Return(42, nil).Once()

		cmd := NewCheckFuel(&fuelObj)
		err := cmd.Execute()
		assert.NoError(t, err)
		fuelObj.AssertExpectations(t)
	})
}

func TestBurnFuel_Execute(t *testing.T) {
	t.Run("error if GetFuelAmount() returns error", func(t *testing.T) {
		fuelObj := mocks.Fuel{}
		fuelObj.On("GetFuelAmount").Return(0, engine.ErrFuelIsNotSupported).Once()

		cmd := NewBurnFuel(&fuelObj)
		err := cmd.Execute()
		assert.ErrorIs(t, err, engine.ErrFuelIsNotSupported)
		fuelObj.AssertExpectations(t)
	})

	t.Run("error if GetFuelBurnRate() returns error", func(t *testing.T) {
		fuelObj := mocks.Fuel{}
		fuelObj.On("GetFuelAmount").Return(1, nil).Once()
		fuelObj.On("GetFuelBurnRate").Return(0, engine.ErrFuelIsNotSupported).Once()

		cmd := NewBurnFuel(&fuelObj)
		err := cmd.Execute()
		assert.ErrorIs(t, err, engine.ErrFuelIsNotSupported)
		fuelObj.AssertExpectations(t)
	})

	t.Run("error if SetFuelAmount() returns error", func(t *testing.T) {
		fuelObj := mocks.Fuel{}
		fuelObj.On("GetFuelAmount").Return(1, nil).Once()
		fuelObj.On("GetFuelBurnRate").Return(1, nil).Once()
		fuelObj.On("SetFuelAmount", 0).Return(engine.ErrFuelIsNotSupported).Once()

		cmd := NewBurnFuel(&fuelObj)
		err := cmd.Execute()
		assert.ErrorIs(t, err, engine.ErrFuelIsNotSupported)
		fuelObj.AssertExpectations(t)
	})

	t.Run("decrease fuel amount by burn rate", func(t *testing.T) {
		amount, rate := 5, 2
		fuelObj := mocks.Fuel{}
		fuelObj.On("GetFuelAmount").Return(amount, nil).Once()
		fuelObj.On("GetFuelBurnRate").Return(rate, nil).Once()
		fuelObj.On("SetFuelAmount", amount-rate).Return(nil).Once()

		cmd := NewBurnFuel(&fuelObj)
		err := cmd.Execute()
		assert.NoError(t, err)
		fuelObj.AssertExpectations(t)
	})
}
