package moving

import (
	"fmt"

	"arc-homework/space-game/command"
	"arc-homework/space-game/moving/engine"
)

type CheckFuel struct {
	fuelObj engine.Fuel
}

func (f *CheckFuel) Execute() error {
	fuelAmount, err := f.fuelObj.GetFuelAmount()
	if err != nil {
		return err
	}

	if fuelAmount <= 0 {
		return fmt.Errorf("%w: not enough fuel", command.ErrCommand)
	}

	return nil
}

func NewCheckFuel(obj engine.Fuel) *CheckFuel {
	return &CheckFuel{fuelObj: obj}
}

type BurnFuel struct {
	fuelObj engine.Fuel
}

func (f *BurnFuel) Execute() error {
	fuelAmount, err := f.fuelObj.GetFuelAmount()
	if err != nil {
		return err
	}

	burnRate, err := f.fuelObj.GetFuelBurnRate()
	if err != nil {
		return err
	}

	return f.fuelObj.SetFuelAmount(fuelAmount - burnRate)
}

func NewBurnFuel(obj engine.Fuel) *BurnFuel {
	return &BurnFuel{fuelObj: obj}
}
