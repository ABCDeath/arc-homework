package command

import "errors"

var (
	ErrCommand = errors.New("command error")
)

type Command interface {
	Execute() error
}
