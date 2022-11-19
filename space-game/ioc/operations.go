package ioc

import (
	"context"
	"sync"

	"arc-homework/space-game/command"
)

const (
	Register     = "IoC.Register"
	NewScope     = "IoC.Scope.New"
	CurrentScope = "IoC.Scope.Current"
)

var restrictedToRegisterOperations = map[string]struct{}{
	Register:     {},
	NewScope:     {},
	CurrentScope: {},
}

type registerOp struct {
	name    string
	builder func(context.Context, ...interface{}) (command.Command, error)
	storage *sync.Map
}

func (o *registerOp) Execute(_ context.Context) error {
	o.storage.Store(o.name, o.builder)

	return nil
}

func NewRegister(storage *sync.Map, name string, builder Builder) *registerOp {
	return &registerOp{
		name:    name,
		builder: builder,
		storage: storage,
	}
}

type newScopeOp struct {
	scope   string
	storage *sync.Map
}

func (o *newScopeOp) Execute(_ context.Context) error {
	o.storage.Store(o.scope, &sync.Map{})

	return nil
}

func NewNewScopeOp(scope string, storage *sync.Map) *newScopeOp {
	return &newScopeOp{
		scope:   scope,
		storage: storage,
	}
}

type currentScopeOp struct{}

func (o *currentScopeOp) Execute(_ context.Context) error {
	// there is no goroutine or thread local storage in golang
	// and golang does not have a mutable context
	// and IoC or Command interfaces do not allow us to return a context.Context
	// by the way returning context.Context is not right, only passing as an argument is allowed
	// https://go.dev/blog/context
	return nil
}

func NewCurrentScopeOp() *currentScopeOp {
	return &currentScopeOp{}
}
