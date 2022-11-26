package worker

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"arc-homework/space-game/command"
	"arc-homework/space-game/queue"
)

type dummyCmd struct {
	wait   chan struct{}
	signal chan struct{}
}

func (c *dummyCmd) Execute(_ context.Context) error {
	// channels are preferable for synchronization than sync.Cond
	<-c.wait
	c.signal <- struct{}{}
	close(c.signal)

	return nil
}

func Test_startCommand_Execute(t *testing.T) {
	t.Run("runs worker in a goroutine", func(t *testing.T) {
		wg := sync.WaitGroup{}
		cmdWait := make(chan struct{})
		waitForCmd := make(chan struct{})
		ctx, cancel := context.WithCancel(context.Background())

		t.Cleanup(func() {
			close(cmdWait)
			wg.Wait()
		})

		cmd := dummyCmd{
			wait:   cmdWait,
			signal: waitForCmd,
		}
		cmdQueue := queue.NewSyncQueue[command.Command]()
		cmdQueue.Enqueue(&cmd)

		startCmd := NewStartCommand(&wg, cmdQueue, func(err error) {}, make(SoftStopSignal))
		err := startCmd.Execute(ctx)
		require.NoError(t, err)

		cmdWait <- struct{}{}

	loop:
		for {
			select {
			case <-waitForCmd:
				break loop
			case <-time.After(time.Millisecond):
				assert.Fail(t, "time out")
				break loop
			}
		}

		cancel()
	})
}
