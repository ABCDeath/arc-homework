package worker

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_softStopCommand_Execute(t *testing.T) {
	t.Run("closes channel", func(t *testing.T) {
		ch := make(SoftStopSignal)
		cmd := NewSoftStopCommand(ch)

		err := cmd.Execute(context.Background())
		assert.NoError(t, err)
		assert.Len(t, ch, 0)

		_, isOpen := <-ch
		assert.False(t, isOpen)
	})
}
