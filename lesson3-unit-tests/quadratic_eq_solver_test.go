package lesson3_unit_tests

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	type args struct {
		a float64
		b float64
		c float64
	}

	tests := []struct {
		name          string
		args          args
		assertRoots   bool
		expectedRoots []float64
		expectedErr   error
	}{
		{
			name:          "x^2 + 1 = 0 has no roots",
			args:          args{a: 1, b: 0, c: 1},
			assertRoots:   true,
			expectedRoots: []float64{},
		},
		{
			name:          "x^2 - 1 = 0 has two roots: 1 and -1)",
			args:          args{a: 1, b: 0, c: -1},
			assertRoots:   true,
			expectedRoots: []float64{1, -1},
		},
		{
			name:          "equation has one root if d is within (-epsilon; epsilon)",
			args:          args{a: 0.5, b: 3, c: 4.5 + 0.25*epsilon},
			assertRoots:   true,
			expectedRoots: []float64{-3, -3},
		},
		{
			name:        "error if a == 0",
			args:        args{a: 0, b: 2, c: 1},
			assertRoots: true,
			expectedErr: ErrInvalidCoefficient,
		},
		{
			name:        "error if a is within (-epsilon; 0]",
			args:        args{a: -epsilon + 1e-15, b: 2, c: 1},
			assertRoots: true,
			expectedErr: ErrInvalidCoefficient,
		},
		{
			name:        "error if a is within [0; epsilon)",
			args:        args{a: epsilon - 1e-15, b: 2, c: 1},
			assertRoots: true,
			expectedErr: ErrInvalidCoefficient,
		},
		{
			name:        "no error if a is close to epsilon, but slightly greater",
			args:        args{a: 1e-42 + epsilon, b: 2, c: 1},
			assertRoots: false,
		},
		{
			name:        "error if a is NaN",
			args:        args{a: math.NaN(), b: 1, c: 1},
			assertRoots: true,
			expectedErr: ErrCoefficientNaN,
		},
		{
			name:        "error if b is NaN",
			args:        args{a: 1, b: math.NaN(), c: 1},
			assertRoots: true,
			expectedErr: ErrCoefficientNaN,
		},
		{
			name:        "error if c is NaN",
			args:        args{a: 1, b: 1, c: math.NaN()},
			assertRoots: true,
			expectedErr: ErrCoefficientNaN,
		},
		{
			name:        "error if a is -Inf",
			args:        args{a: math.Inf(-1), b: 1, c: 1},
			assertRoots: true,
			expectedErr: ErrCoefficientInf,
		},
		{
			name:        "error if b is -Inf",
			args:        args{a: 1, b: math.Inf(-1), c: 1},
			assertRoots: true,
			expectedErr: ErrCoefficientInf,
		},
		{
			name:        "error if c is -Inf",
			args:        args{a: 1, b: 1, c: math.Inf(-1)},
			assertRoots: true,
			expectedErr: ErrCoefficientInf,
		},
		{
			name:        "error if a is +Inf",
			args:        args{a: math.Inf(1), b: 1, c: 1},
			assertRoots: true,
			expectedErr: ErrCoefficientInf,
		},
		{
			name:        "error if b is +Inf",
			args:        args{a: 1, b: math.Inf(1), c: 1},
			assertRoots: true,
			expectedErr: ErrCoefficientInf,
		},
		{
			name:        "error if c is +Inf",
			args:        args{a: 1, b: 1, c: math.Inf(1)},
			assertRoots: true,
			expectedErr: ErrCoefficientInf,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRoots, err := Solve(tt.args.a, tt.args.b, tt.args.c)

			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
				assert.Empty(t, gotRoots)
			} else {
				assert.NoError(t, err)
				if tt.assertRoots {
					assert.ElementsMatch(t, tt.expectedRoots, gotRoots)
				}
			}
		})
	}
}
