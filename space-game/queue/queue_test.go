package queue

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_syncQueue_Dequeue(t *testing.T) {
	t.Run("returns error if queue is empty", func(t *testing.T) {
		q := syncQueue[int]{
			objects:           []int{},
			lock:              sync.RWMutex{},
			qIsNotEmptySignal: nil,
		}

		_, err := q.Dequeue()
		assert.Error(t, err, ErrQueueEmpty)
	})

	t.Run("returns object", func(t *testing.T) {
		q := syncQueue[int]{
			objects:           []int{1},
			lock:              sync.RWMutex{},
			qIsNotEmptySignal: nil,
		}

		object, err := q.Dequeue()

		require.NoError(t, err)
		assert.Equal(t, 1, object)
		assert.Empty(t, q.objects)
	})

	t.Run("returns first object", func(t *testing.T) {
		q := syncQueue[int]{
			objects:           []int{1, 2},
			lock:              sync.RWMutex{},
			qIsNotEmptySignal: nil,
		}

		object, err := q.Dequeue()

		require.NoError(t, err)
		assert.Equal(t, 1, object)
		assert.Len(t, q.objects, 1)
	})

	t.Run("enqueue and dequeue", func(t *testing.T) {
		q := syncQueue[int]{
			objects:           []int{},
			lock:              sync.RWMutex{},
			qIsNotEmptySignal: make(chan struct{}, 2),
		}

		q.Enqueue(1)
		object, err := q.Dequeue()

		require.NoError(t, err)
		assert.Equal(t, 1, object)
		assert.Empty(t, q.objects)
	})
}

func Test_syncQueue_DequeueOrWait(t *testing.T) {
	t.Run("waits for object if queue is empty", func(t *testing.T) {
		q := syncQueue[int]{
			objects:           []int{},
			lock:              sync.RWMutex{},
			qIsNotEmptySignal: make(chan struct{}, 2),
		}

		wg := sync.WaitGroup{}
		c := make(chan int)
		t.Cleanup(func() {
			wg.Wait()
			close(c)
		})

		var err error
		go func() {
			var obj int
			obj, err = q.DequeueOrWait(context.Background())
			c <- obj
			wg.Done()
		}()
		wg.Add(1)
		time.Sleep(time.Millisecond)

		assert.Empty(t, c)

		q.Enqueue(1)

		require.NoError(t, err)
		assert.Equal(t, 1, <-c)
		assert.Empty(t, q.objects)
	})

	t.Run("returns object if queue is not empty", func(t *testing.T) {
		q := syncQueue[int]{
			objects:           []int{1},
			lock:              sync.RWMutex{},
			qIsNotEmptySignal: make(chan struct{}, 2),
		}

		actual, err := q.DequeueOrWait(context.Background())
		require.NoError(t, err)
		assert.Equal(t, 1, actual)
		assert.Empty(t, q.objects)
	})

	t.Run("stops waiting if context canceled", func(t *testing.T) {
		q := syncQueue[int]{
			objects:           []int{},
			lock:              sync.RWMutex{},
			qIsNotEmptySignal: make(chan struct{}, 2),
		}

		wg := sync.WaitGroup{}
		ctx, cancelContext := context.WithCancel(context.Background())
		t.Cleanup(func() {
			wg.Wait()
		})

		var err error
		go func() {
			_, err = q.DequeueOrWait(ctx)
			wg.Done()
		}()
		wg.Add(1)
		time.Sleep(time.Millisecond)
		cancelContext()
		wg.Wait()

		assert.Error(t, err, context.Canceled)
	})
}

func Test_syncQueue_Enqueue(t *testing.T) {
	t.Run("puts object into internal slice", func(t *testing.T) {
		q := syncQueue[int]{
			objects:           []int{},
			lock:              sync.RWMutex{},
			qIsNotEmptySignal: make(chan struct{}, 2),
		}

		q.Enqueue(1)

		require.Len(t, q.objects, 1)
		assert.ElementsMatch(t, []int{1}, q.objects)
	})

	t.Run("puts object at the end of the internal slice", func(t *testing.T) {
		q := syncQueue[int]{
			objects:           []int{1},
			lock:              sync.RWMutex{},
			qIsNotEmptySignal: make(chan struct{}, 2),
		}

		q.Enqueue(2)

		require.Len(t, q.objects, 2)
		require.Equal(t, 1, q.objects[0])
		assert.Equal(t, 2, q.objects[1])
	})
}

func Test_syncQueue_isEmpty(t *testing.T) {
	t.Run("true if queue is empty", func(t *testing.T) {
		q := syncQueue[int]{
			objects:           []int{},
			lock:              sync.RWMutex{},
			qIsNotEmptySignal: nil,
		}

		assert.Empty(t, q.objects)
	})

	t.Run("false if queue is not empty", func(t *testing.T) {
		q := syncQueue[int]{
			objects:           []int{1},
			lock:              sync.RWMutex{},
			qIsNotEmptySignal: nil,
		}

		assert.NotEmpty(t, q.objects)
	})
}
