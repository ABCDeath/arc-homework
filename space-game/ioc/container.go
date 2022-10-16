package ioc

import (
	"context"
	"fmt"
	"sync"

	"arc-homework/space-game/command"
)

type ContextScopeKeyType string

const ContextScopeKey = ContextScopeKeyType("ioc.scope")

type IoC interface {
	Resolve(ctx context.Context, name string, params ...interface{}) (command.Command, error)
}

type Builder func(ctx context.Context, args ...interface{}) (command.Command, error)

type ioc struct {
	defaultDeps *sync.Map
	scopeDeps   *sync.Map
}

func (i *ioc) Resolve(ctx context.Context, name string, params ...interface{}) (command.Command, error) {
	var storage *sync.Map
	if _, found := restrictedToRegisterOperations[name]; found {
		storage = i.defaultDeps
	} else {
		st, err := i.getStorage(ctx)
		if err != nil {
			return nil, err
		}
		storage = st
	}

	builder, exists := storage.Load(name)
	if !exists {
		return nil, fmt.Errorf("%s: %w", name, ErrNotFound)
	}

	builderCast := builder.(func(ctx context.Context, args ...interface{}) (command.Command, error))

	return builderCast(ctx, params...)
}

func (i *ioc) getStorage(ctx context.Context) (*sync.Map, error) {
	scopePtr := ctx.Value(ContextScopeKey)
	if scopePtr == nil {
		return i.defaultDeps, nil
	}

	scopeName, ok := scopePtr.(string)
	if !ok {
		return nil, fmt.Errorf("%w: context scope type cast error", ErrIoC)
	}

	if storage, exists := i.scopeDeps.Load(scopeName); !exists {
		return i.defaultDeps, nil
	} else {
		return storage.(*sync.Map), nil
	}
}

func New() IoC {
	ioc := &ioc{
		defaultDeps: &sync.Map{},
		scopeDeps:   &sync.Map{},
	}

	initContainerOperations(ioc)

	return ioc
}

func initContainerOperations(ioc *ioc) {
	ioc.defaultDeps.Store(Register, func(ctx context.Context, args ...interface{}) (command.Command, error) {
		name, ok := args[0].(string)
		if !ok {
			return nil, fmt.Errorf("%w: args[0] expected: string", ErrInvalidArgs)
		}

		// can not re-define Register or Scope.New/Scope.Set
		if _, found := restrictedToRegisterOperations[name]; found {
			return nil, ErrRegisterRestricted
		}

		// can not use Builder type alias here
		builder, ok := args[1].(func(ctx context.Context, args ...interface{}) (command.Command, error))
		if !ok {
			return nil, fmt.Errorf("%w: args[1] expected: Builder", ErrInvalidArgs)
		}

		storage, err := ioc.getStorage(ctx)
		if err != nil {
			panic(err)
		}

		return NewRegister(storage, name, builder), nil
	})

	ioc.defaultDeps.Store(NewScope, func(ctx context.Context, args ...interface{}) (command.Command, error) {
		scope, ok := args[0].(string)
		if !ok {
			return nil, fmt.Errorf("%w: args[0] expected: string", ErrInvalidArgs)
		}

		return NewNewScopeOp(scope, ioc.scopeDeps), nil
	})

	ioc.defaultDeps.Store(CurrentScope, func(ctx context.Context, args ...interface{}) (command.Command, error) {
		return NewCurrentScopeOp(), nil
	})
}
