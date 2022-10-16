package ioc

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"arc-homework/space-game/command"
)

func Test_registerOp_Execute(t *testing.T) {
	t.Run("adds handler to storage", func(t *testing.T) {
		newOpName := "Foo.Bar"
		storage := &sync.Map{}
		builder := func(ctx context.Context, args ...interface{}) (command.Command, error) { return nil, nil }

		op := NewRegister(storage, newOpName, builder)
		err := op.Execute()

		assert.NoError(t, err)
		_, found := storage.Load(newOpName)
		assert.True(t, found)
	})
}

func Test_newScopeOp_Execute(t *testing.T) {
	t.Run("adds handler to storage", func(t *testing.T) {
		scope := "Scope.1"
		storage := &sync.Map{}

		op := NewNewScopeOp(scope, storage)
		err := op.Execute()

		assert.NoError(t, err)
		_, found := storage.Load(scope)
		assert.True(t, found)
	})
}

func Test_currentScopeOp_Execute(t *testing.T) {
	t.Run("does nothing", func(t *testing.T) {
		op := NewCurrentScopeOp()
		err := op.Execute()

		assert.NoError(t, err)
	})
}
