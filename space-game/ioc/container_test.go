package ioc

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"arc-homework/space-game/command"
	"arc-homework/space-game/command/mocks"
)

var dummyBuilder = func(ctx context.Context, args ...interface{}) (command.Command, error) {
	return nil, nil
}

func Test_initContainerOperations(t *testing.T) {
	basicOperations := []string{Register, NewScope, CurrentScope}

	t.Run("adds basic container operations", func(t *testing.T) {
		ioc := &ioc{
			defaultDeps: &sync.Map{},
			scopeDeps:   &sync.Map{},
		}

		initContainerOperations(ioc)

		for _, opName := range basicOperations {
			_, found := ioc.defaultDeps.Load(opName)
			assert.True(t, found)
		}
	})

	t.Run("Register returns error if name arg has invalid type", func(t *testing.T) {
		ioc := &ioc{
			defaultDeps: &sync.Map{},
			scopeDeps:   &sync.Map{},
		}

		initContainerOperations(ioc)
		registerHandler, _ := ioc.defaultDeps.Load(Register)
		handler := registerHandler.(func(ctx context.Context, args ...interface{}) (command.Command, error))
		_, err := handler(context.Background(), 123, dummyBuilder)
		assert.ErrorIs(t, err, ErrInvalidArgs)
	})

	t.Run("Register retruns error if handler has invalid signature", func(t *testing.T) {
		opName := "Foo.Bar"
		ioc := &ioc{
			defaultDeps: &sync.Map{},
			scopeDeps:   &sync.Map{},
		}

		initContainerOperations(ioc)
		registerHandler, _ := ioc.defaultDeps.Load(Register)
		handler := registerHandler.(func(ctx context.Context, args ...interface{}) (command.Command, error))
		_, err := handler(context.Background(), opName, func(foo int, bar string) {})
		assert.ErrorIs(t, err, ErrInvalidArgs)
	})

	t.Run("Register adds handler to default if no scope in context", func(t *testing.T) {
		opName := "Foo.Bar"
		ioc := &ioc{
			defaultDeps: &sync.Map{},
			scopeDeps:   &sync.Map{},
		}

		initContainerOperations(ioc)
		registerHandler, _ := ioc.defaultDeps.Load(Register)
		handler := registerHandler.(func(ctx context.Context, args ...interface{}) (command.Command, error))
		cmd, err := handler(context.Background(), opName, dummyBuilder)
		assert.NoError(t, err)

		err = cmd.Execute()
		assert.NoError(t, err)

		_, found := ioc.defaultDeps.Load(opName)
		assert.True(t, found)
	})

	t.Run("Register adds handler to specified scope in context", func(t *testing.T) {
		scopeName := "Scope.1"
		ctx := context.WithValue(context.Background(), ContextScopeKey, scopeName)
		opName := "Foo.Bar"
		scopeDeps := &sync.Map{}
		scopeDeps.Store(scopeName, &sync.Map{})
		ioc := &ioc{
			defaultDeps: &sync.Map{},
			scopeDeps:   scopeDeps,
		}

		initContainerOperations(ioc)
		registerHandler, _ := ioc.defaultDeps.Load(Register)
		handler := registerHandler.(func(ctx context.Context, args ...interface{}) (command.Command, error))
		cmd, err := handler(ctx, opName, dummyBuilder)
		assert.NoError(t, err)

		err = cmd.Execute()
		assert.NoError(t, err)

		_, found := ioc.scopeDeps.Load(scopeName)
		require.True(t, found)
		s, _ := ioc.scopeDeps.Load(scopeName)
		storage := s.(*sync.Map)
		_, found = storage.Load(opName)
		assert.True(t, found)
	})

	t.Run("Scope.New returns error if scopeID arg has invalid type", func(t *testing.T) {
		ioc := &ioc{
			defaultDeps: &sync.Map{},
			scopeDeps:   &sync.Map{},
		}

		initContainerOperations(ioc)
		scopeHandler, _ := ioc.defaultDeps.Load(NewScope)
		handler := scopeHandler.(func(ctx context.Context, args ...interface{}) (command.Command, error))
		_, err := handler(context.Background(), 123)
		assert.ErrorIs(t, err, ErrInvalidArgs)
	})

	t.Run("Scope.New creates new scope in storage", func(t *testing.T) {
		scope := "Scope.1"
		ioc := &ioc{
			defaultDeps: &sync.Map{},
			scopeDeps:   &sync.Map{},
		}

		initContainerOperations(ioc)
		scopeHandler, _ := ioc.defaultDeps.Load(NewScope)
		handler := scopeHandler.(func(ctx context.Context, args ...interface{}) (command.Command, error))
		cmd, err := handler(context.Background(), scope)
		assert.NoError(t, err)

		err = cmd.Execute()
		assert.NoError(t, err)
		_, found := ioc.scopeDeps.Load(scope)
		assert.True(t, found)
	})

	t.Run("Scope.Current does nothing", func(t *testing.T) {
		scope := "Scope.1"
		ioc := &ioc{
			defaultDeps: &sync.Map{},
			scopeDeps:   &sync.Map{},
		}

		initContainerOperations(ioc)
		scopeHandler, _ := ioc.defaultDeps.Load(CurrentScope)
		handler := scopeHandler.(func(ctx context.Context, args ...interface{}) (command.Command, error))
		cmd, err := handler(context.Background(), scope)
		assert.NoError(t, err)

		err = cmd.Execute()
		assert.NoError(t, err)
	})
}

func TestIoC_Concurrently(t *testing.T) {
	ioc := &ioc{
		defaultDeps: &sync.Map{},
		scopeDeps:   &sync.Map{},
	}

	initContainerOperations(ioc)

	t.Run("use different scopes in goroutines", func(t *testing.T) {
		scope1, scope2, scope3 := "Scope.1", "Scope.2", "Scope.3"
		handler1, handler2, handler3 := "foo1", "foo2", "foo3"
		handler1Mock, handler2Mock, handler3Mock := mocks.Command{}, mocks.Command{}, mocks.Command{}
		handler1Mock.On("Execute").Return(nil).Once()
		handler2Mock.On("Execute").Return(nil).Once()
		handler3Mock.On("Execute").Return(nil).Once()
		wg := sync.WaitGroup{}

		go func() {
			defer wg.Done()

			handle := func(ctx context.Context, args ...interface{}) (command.Command, error) {
				return &handler1Mock, nil
			}

			ctx := context.WithValue(context.Background(), ContextScopeKey, scope1)

			cmd, err := ioc.Resolve(ctx, NewScope, scope1)
			require.NoError(t, err)
			err = cmd.Execute()
			require.NoError(t, err)

			cmd, err = ioc.Resolve(ctx, CurrentScope, scope1)
			require.NoError(t, err)
			err = cmd.Execute()
			require.NoError(t, err)

			cmd, err = ioc.Resolve(ctx, Register, handler1, handle)
			require.NoError(t, err)
			err = cmd.Execute()
			require.NoError(t, err)

			cmd, err = ioc.Resolve(ctx, handler1)
			require.NoError(t, err)
			err = cmd.Execute()
			require.NoError(t, err)
		}()
		wg.Add(1)

		go func() {
			defer wg.Done()

			handle := func(ctx context.Context, args ...interface{}) (command.Command, error) {
				return &handler2Mock, nil
			}

			ctx := context.WithValue(context.Background(), ContextScopeKey, scope3)

			cmd, err := ioc.Resolve(ctx, NewScope, scope2)
			require.NoError(t, err)
			err = cmd.Execute()
			require.NoError(t, err)

			cmd, err = ioc.Resolve(ctx, CurrentScope, scope2)
			require.NoError(t, err)
			err = cmd.Execute()
			require.NoError(t, err)

			cmd, err = ioc.Resolve(ctx, Register, handler2, handle)
			require.NoError(t, err)
			err = cmd.Execute()
			require.NoError(t, err)

			cmd, err = ioc.Resolve(ctx, handler2)
			require.NoError(t, err)
			err = cmd.Execute()
			require.NoError(t, err)
		}()
		wg.Add(1)

		go func() {
			defer wg.Done()

			handle := func(ctx context.Context, args ...interface{}) (command.Command, error) {
				return &handler3Mock, nil
			}

			ctx := context.WithValue(context.Background(), ContextScopeKey, scope3)

			cmd, err := ioc.Resolve(ctx, NewScope, scope3)
			require.NoError(t, err)
			err = cmd.Execute()
			require.NoError(t, err)

			cmd, err = ioc.Resolve(ctx, CurrentScope, scope3)
			require.NoError(t, err)
			err = cmd.Execute()
			require.NoError(t, err)

			cmd, err = ioc.Resolve(ctx, Register, handler3, handle)
			require.NoError(t, err)
			err = cmd.Execute()
			require.NoError(t, err)

			cmd, err = ioc.Resolve(ctx, handler3)
			require.NoError(t, err)
			err = cmd.Execute()
			require.NoError(t, err)
		}()
		wg.Add(1)

		wg.Wait()

		handler1Mock.AssertExpectations(t)
		handler2Mock.AssertExpectations(t)
		handler3Mock.AssertExpectations(t)
	})
}
