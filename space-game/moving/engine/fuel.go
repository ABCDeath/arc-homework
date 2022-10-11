package engine

import (
	"errors"

	"arc-homework/space-game/moving/object"
)

const (
	FuelAmountPropName   = "fuel_amount"
	FuelBurnRatePropName = "fuel_burn_rate"
)

var ErrFuelIsNotSupported = errors.New("object does not have fuel")

type Fuel interface {
	GetFuelAmount() (int, error)
	GetFuelBurnRate() (int, error)
	SetFuelAmount(amount int) error
}

type adapter struct {
	obj object.Object
}

func (a *adapter) GetFuelAmount() (int, error) {
	fuelAmount, err := getProperty[int](a.obj, FuelAmountPropName)
	if err != nil {
		return 0, err
	}

	return *fuelAmount, nil
}

func (a *adapter) GetFuelBurnRate() (int, error) {
	rate, err := getProperty[int](a.obj, FuelBurnRatePropName)
	if err != nil {
		return 0, ErrFuelIsNotSupported
	}

	return *rate, nil
}

func (a *adapter) SetFuelAmount(amount int) error {
	err := a.obj.SetProperty(FuelAmountPropName, amount)
	if err != nil {
		return ErrFuelIsNotSupported
	}

	return nil
}

func getProperty[T any](obj object.Object, name string) (*T, error) {
	value, err := object.GetObjectProperty[T](obj, name)
	if err != nil {
		return nil, ErrFuelIsNotSupported
	}

	return value, nil
}

func New(obj object.Object) *adapter {
	return &adapter{obj: obj}
}
