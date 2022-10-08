package error_handler

import (
	"errors"
	"fmt"
	"reflect"

	"arc-homework/space-game/command"
	"arc-homework/space-game/helper"
)

var ErrNoHandler = errors.New("no handler registered")

type Handle func(command command.Command, err error) error
type handlerByError map[error]Handle
type handlerByCmd map[string]handlerByError
type defaultHandlerByCmd map[string]Handle

type Handler interface {
	Handle(command command.Command, err error) error
	RegisterHandler(commandType reflect.Type, err error, handler Handle)
	RegisterDefaultCmdHandler(commandType reflect.Type, handler Handle)
	RegisterDefaultHandler(handler Handle)
}

type handler struct {
	handlerMapping      handlerByCmd
	defaultHandlerByCmd defaultHandlerByCmd
	defaultHandler      Handle
}

func (h *handler) Handle(command command.Command, err error) error {
	cmdName := helper.GetStructTypeName(command)
	handlerByErr := h.handlerMapping[cmdName]

	var handlerFunc Handle
	if handlerByErr == nil {
		handlerFunc = h.defaultHandlerByCmd[cmdName]
		if handlerFunc == nil {
			handlerFunc = h.defaultHandler
		}
	} else {
		handlerFunc = handlerByErr[err]
	}

	if handlerFunc == nil {
		return fmt.Errorf("%w for cmd: %s, err: %v", ErrNoHandler, cmdName, err)
	}

	return handlerFunc(command, err)
}

func (h *handler) RegisterHandler(commandType reflect.Type, err error, handler Handle) {
	cmdName := commandType.Name()
	if _, ok := h.handlerMapping[cmdName]; !ok {
		h.handlerMapping[cmdName] = make(handlerByError)
	}

	h.handlerMapping[cmdName][err] = handler
}

func (h *handler) RegisterDefaultCmdHandler(commandType reflect.Type, handler Handle) {
	h.defaultHandlerByCmd[commandType.Name()] = handler
}

func (h *handler) RegisterDefaultHandler(handler Handle) {
	h.defaultHandler = handler
}

func New() Handler {
	return &handler{
		handlerMapping:      handlerByCmd{},
		defaultHandlerByCmd: defaultHandlerByCmd{},
	}
}
