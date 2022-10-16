package ioc

import (
	"errors"
	"fmt"
)

var (
	ErrIoC                = errors.New("IoC error")
	ErrNotFound           = fmt.Errorf("%w: is not registered", ErrIoC)
	ErrInvalidArgs        = fmt.Errorf("%w: invalid args", ErrIoC)
	ErrRegisterRestricted = fmt.Errorf("%w: this operation register is not allowed", ErrIoC)
)
