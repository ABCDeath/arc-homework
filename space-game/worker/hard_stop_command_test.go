package worker

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_hardStopCommand_Execute(t *testing.T) {
	t.Run("calls cancel function", func(t *testing.T) {
		var timeoutHappened bool
		wg := sync.WaitGroup{}
		t.Cleanup(func() {
			wg.Wait()
		})

		routine := func(ctx context.Context) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case <-time.After(10 * time.Millisecond):
					timeoutHappened = true
				}
			}
		}

		ctx, cancel := context.WithCancel(context.Background())
		wg.Add(1)
		cmd := NewHardStopCommand(cancel)
		go routine(ctx)

		err := cmd.Execute(ctx)
		assert.NoError(t, err)
		assert.False(t, timeoutHappened)
	})
}
