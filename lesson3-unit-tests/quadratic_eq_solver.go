package lesson3_unit_tests

import (
	"errors"
	"fmt"
	"math"
)

const epsilon = 1e-12

var (
	ErrInvalidCoefficient = errors.New("invalid coefficient value")
	ErrCoefficientNaN     = fmt.Errorf("%w: can not be NaN", ErrInvalidCoefficient)
	ErrCoefficientInf     = fmt.Errorf("%w: can not be Inf", ErrInvalidCoefficient)
)

// Solve solves a quadratic equation a*x^2 + b*x + c = 0
func Solve(a, b, c float64) (roots []float64, err error) {
	err = validateCoefficients(a, b, c)
	if err != nil {
		return nil, err
	}

	var dSqrt float64

	d := math.Pow(b, 2) - 4*a*c
	switch compare(d, 0) {
	case -1:
		return
	case 0:
		dSqrt = 0
	default:
		dSqrt = math.Sqrt(d)
	}

	roots = make([]float64, 2)
	roots[0] = (-b - dSqrt) / (2 * a)
	roots[1] = (-b + dSqrt) / (2 * a)

	return
}

func validateCoefficients(a, b, c float64) error {
	switch {
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c):
		return ErrCoefficientNaN
	case math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0):
		return ErrCoefficientInf
	case compare(a, 0) == 0:
		return fmt.Errorf("%w: a can not be 0", ErrInvalidCoefficient)
	}

	return nil
}

// compare returns -1 if a < b, 1 if a > b and 0 if a == b
func compare(a, b float64) int {
	diff := a - b
	if math.Abs(diff) < epsilon {
		return 0
	}

	if a < b {
		return -1
	}

	return 1
}
