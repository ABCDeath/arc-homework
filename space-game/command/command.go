package command

import (
	"context"
	"errors"
)

var (
	ErrCommand = errors.New("command error")
)

type Command interface {
	Execute(ctx context.Context) error
}
