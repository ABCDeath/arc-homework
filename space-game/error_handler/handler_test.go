package error_handler

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"arc-homework/space-game/command"
	"arc-homework/space-game/command/mocks"
	"arc-homework/space-game/helper"
)

func Test_handler_RegisterHandler(t *testing.T) {
	t.Run("register handler for (cmd, err) puts handler into specific field", func(t *testing.T) {
		callCount, dummyHandler := createDummyHandler()
		dummyErr := errors.New("")

		h := handler{handlerMapping: handlerByCmd{}}
		cmd := mocks.Command{}

		h.RegisterHandler(reflect.TypeOf(&cmd), dummyErr, dummyHandler)
		toErrMapping, exists := h.handlerMapping[reflect.TypeOf(&cmd).Name()]
		require.True(t, exists)

		handle, exists := toErrMapping[dummyErr]
		require.True(t, exists)

		err := handle(&cmd, errors.New(""))
		require.NoError(t, err)
		assert.Equal(t, 1, *callCount)
	})
}

func Test_handler_RegisterDefaultCmdHandler(t *testing.T) {
	t.Run("register default cmd handler puts handler into specific field", func(t *testing.T) {
		callCount, dummyHandler := createDummyHandler()

		h := handler{defaultHandlerByCmd: defaultHandlerByCmd{}}
		cmd := mocks.Command{}

		h.RegisterDefaultCmdHandler(reflect.TypeOf(&cmd), dummyHandler)
		handle, exists := h.defaultHandlerByCmd[reflect.TypeOf(&cmd).Name()]
		require.True(t, exists)

		err := handle(&cmd, errors.New(""))
		require.NoError(t, err)
		assert.Equal(t, 1, *callCount)
	})
}

func Test_handler_RegisterDefaultHandler(t *testing.T) {
	t.Run("register default handler puts handler into specific field", func(t *testing.T) {
		callCount, dummyHandler := createDummyHandler()

		h := handler{}

		h.RegisterDefaultHandler(dummyHandler)
		err := h.defaultHandler(&mocks.Command{}, errors.New(""))
		require.NoError(t, err)

		assert.Equal(t, 1, *callCount)
	})
}

func Test_handler_Handle(t *testing.T) {
	t.Run("error_handler if no handler registered", func(t *testing.T) {
		h := New()

		err := h.Handle(&mocks.Command{}, errors.New(""))
		assert.ErrorIs(t, err, ErrNoHandler)
	})

	t.Run("run default handler if no other found", func(t *testing.T) {
		callCount, dummyHandler := createDummyHandler()
		h := New()
		h.RegisterDefaultHandler(dummyHandler)

		err := h.Handle(&mocks.Command{}, errors.New(""))
		assert.NoError(t, err)
		assert.Equal(t, 1, *callCount)
	})

	t.Run("run command default handler if error_handler specific not found", func(t *testing.T) {
		callCount, dummyHandler := createDummyHandler()
		dummyCmd := &mocks.Command{}
		h := New()
		h.RegisterDefaultCmdHandler(helper.GetStructType(dummyCmd), dummyHandler)

		err := h.Handle(dummyCmd, errors.New(""))
		assert.NoError(t, err)
		assert.Equal(t, 1, *callCount)
	})

	t.Run("run (cmd, err) handler", func(t *testing.T) {
		callCount, dummyHandler := createDummyHandler()
		dummyCmd := &mocks.Command{}
		dummyErr := errors.New("")
		h := New()
		h.RegisterHandler(helper.GetStructType(dummyCmd), dummyErr, dummyHandler)

		err := h.Handle(dummyCmd, dummyErr)
		assert.NoError(t, err)
		assert.Equal(t, 1, *callCount)
	})

	t.Run("multiple commands and handlers", func(t *testing.T) {
		callCount1, dummyHandler1 := createDummyHandler()
		dummyCmd1 := &mocks.Command{}
		dummyErr1 := errors.New("")

		callCount2, dummyHandler2 := createDummyHandler()
		dummyCmd2 := &mocks.Command{}
		dummyErr2 := errors.New("")

		h := New()
		h.RegisterHandler(helper.GetStructType(dummyCmd1), dummyErr1, dummyHandler1)
		h.RegisterHandler(helper.GetStructType(dummyCmd2), dummyErr2, dummyHandler2)

		err := h.Handle(dummyCmd1, dummyErr1)
		assert.NoError(t, err)
		assert.Equal(t, 1, *callCount1)
		assert.Equal(t, 0, *callCount2)

		err = h.Handle(dummyCmd2, dummyErr2)
		assert.NoError(t, err)
		assert.Equal(t, 1, *callCount2)
		assert.Equal(t, 1, *callCount1)
	})
}

func createDummyHandler() (*int, Handle) {
	callCount := 0
	dummyHandler := func(command command.Command, err error) error {
		callCount++

		return nil
	}

	return &callCount, dummyHandler
}
