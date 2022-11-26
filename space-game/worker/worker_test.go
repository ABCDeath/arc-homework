package worker

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"arc-homework/space-game/command"
	commandMocks "arc-homework/space-game/command/mocks"
	"arc-homework/space-game/queue"
	"arc-homework/space-game/queue/mocks"
)

func Test_worker_getCommandAndExecute(t *testing.T) {
	t.Run("calls non-blocking Queue.Dequeue if stopIfEmpty == true", func(t *testing.T) {
		cmd := commandMocks.Command{}
		cmd.On("Execute", mock.Anything).Return(nil)

		q := mocks.Queue[command.Command]{}
		q.On("Dequeue").Return(&cmd, nil).Once()

		w := worker{
			wg:               &sync.WaitGroup{},
			commandQueue:     &q,
			errHandler:       nil,
			stopIfQueueEmpty: nil,
			stopIfEmpty:      true,
		}

		err := w.getCommandAndExecute(context.Background())
		assert.NoError(t, err)
		q.AssertExpectations(t)
	})

	t.Run("calls Queue.DequeueOrWait if stopIfEmpty == false", func(t *testing.T) {
		cmd := commandMocks.Command{}
		cmd.On("Execute", mock.Anything).Return(nil)

		q := mocks.Queue[command.Command]{}
		q.On("DequeueOrWait", context.Background()).Return(&cmd, nil).Once()

		w := worker{
			wg:               &sync.WaitGroup{},
			commandQueue:     &q,
			errHandler:       nil,
			stopIfQueueEmpty: nil,
			stopIfEmpty:      false,
		}

		err := w.getCommandAndExecute(context.Background())
		assert.NoError(t, err)
		q.AssertExpectations(t)
	})

	t.Run("calls Command.Execute", func(t *testing.T) {
		cmd := commandMocks.Command{}
		cmd.On("Execute", mock.Anything).Return(nil)

		q := mocks.Queue[command.Command]{}
		q.On("Dequeue").Return(&cmd, nil).Once()

		w := worker{
			wg:               &sync.WaitGroup{},
			commandQueue:     &q,
			errHandler:       nil,
			stopIfQueueEmpty: nil,
			stopIfEmpty:      true,
		}

		err := w.getCommandAndExecute(context.Background())
		assert.NoError(t, err)
		cmd.AssertExpectations(t)
	})
}

func Test_worker_Run(t *testing.T) {
	t.Run("stops immediately if context canceled", func(t *testing.T) {
		q := queue.NewSyncQueue[command.Command]()
		q.Enqueue(&commandMocks.Command{})
		wg := sync.WaitGroup{}
		ctx, cancel := context.WithCancel(context.Background())
		stop := make(SoftStopSignal)

		w := worker{
			wg:               &wg,
			commandQueue:     q,
			errHandler:       nil,
			stopIfQueueEmpty: stop,
			stopIfEmpty:      false,
		}

		wg.Add(1)
		go func() {
			err := w.Run(ctx)
			assert.ErrorIs(t, err, context.Canceled)
		}()

		cancel()
		wg.Wait()
		_, err := q.Dequeue()
		assert.NoError(t, err)
	})

	t.Run("stops if queue is empty and stop channel is closed", func(t *testing.T) {
		cmd := commandMocks.Command{}
		cmd.On("Execute", mock.Anything).Return(nil)
		q := queue.NewSyncQueue[command.Command]()
		q.Enqueue(&cmd)

		wg := sync.WaitGroup{}
		ctx := context.Background()
		stop := make(SoftStopSignal)

		w := worker{
			wg:               &wg,
			commandQueue:     q,
			errHandler:       nil,
			stopIfQueueEmpty: stop,
			stopIfEmpty:      false,
		}

		wg.Add(1)
		go func() {
			err := w.Run(ctx)
			assert.NoError(t, err)
		}()

		close(stop)
		wg.Wait()
		_, err := q.Dequeue()
		assert.ErrorIs(t, err, queue.ErrQueueEmpty)
		cmd.AssertExpectations(t)
	})

	t.Run("gets command from queue and executes it", func(t *testing.T) {
		cmd := commandMocks.Command{}
		cmd.On("Execute", mock.Anything).Return(nil)
		q := queue.NewSyncQueue[command.Command]()
		q.Enqueue(&cmd)

		wg := sync.WaitGroup{}
		ctx, cancel := context.WithCancel(context.Background())
		stop := make(SoftStopSignal)

		t.Cleanup(func() {
			cancel()
		})

		w := worker{
			wg:               &wg,
			commandQueue:     q,
			errHandler:       nil,
			stopIfQueueEmpty: stop,
			stopIfEmpty:      false,
		}

		wg.Add(1)
		go func() {
			err := w.Run(ctx)
			assert.NoError(t, err)
		}()

		close(stop)

		wg.Wait()
		_, err := q.Dequeue()
		assert.Error(t, err, queue.ErrQueueEmpty)
		cmd.AssertExpectations(t)
	})

	t.Run("call error handler callback if Command.Execute returns error", func(t *testing.T) {
		cmdErr := errors.New("")
		cmd := commandMocks.Command{}
		cmd.On("Execute", mock.Anything).Return(cmdErr)
		q := queue.NewSyncQueue[command.Command]()
		q.Enqueue(&cmd)

		wg := sync.WaitGroup{}
		ctx, cancel := context.WithCancel(context.Background())
		stop := make(SoftStopSignal)

		t.Cleanup(func() {
			cancel()
		})

		called := false
		callback := func(err error) {
			called = true
			assert.Error(t, err, cmdErr)
		}

		w := worker{
			wg:               &wg,
			commandQueue:     q,
			errHandler:       callback,
			stopIfQueueEmpty: stop,
			stopIfEmpty:      false,
		}

		wg.Add(1)
		go func() {
			err := w.Run(ctx)
			assert.NoError(t, err)
		}()

		close(stop)

		wg.Wait()
		_, err := q.Dequeue()
		assert.Error(t, err, queue.ErrQueueEmpty)
		cmd.AssertExpectations(t)
		assert.True(t, called)
	})
}
