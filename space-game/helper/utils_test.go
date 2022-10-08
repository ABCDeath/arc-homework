package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStructTypeName(t *testing.T) {
	t.Run("returns struct type name passed by pointer", func(t *testing.T) {
		type SomeTestStructType struct{}
		expectedName := "SomeTestStructType"

		actualName := GetStructTypeName(&SomeTestStructType{})
		assert.Equal(t, expectedName, actualName)
	})
}
