package queue

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var (
	ErrQueue      = errors.New("")
	ErrQueueEmpty = fmt.Errorf("%wQueue is empty", ErrQueue)
)

type Queue[T any] interface {
	Enqueue(object T)
	Dequeue() (T, error)
	DequeueOrWait(ctx context.Context) (T, error)
}

type syncQueue[T any] struct {
	objects           []T
	lock              sync.RWMutex
	qIsNotEmptySignal chan struct{}
}

func (q *syncQueue[T]) Enqueue(object T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.isEmpty() {
		defer func() {
			q.qIsNotEmptySignal <- struct{}{}
		}()
	}

	q.objects = append(q.objects, object)
}

func (q *syncQueue[T]) Dequeue() (T, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.isEmpty() {
		return *new(T), ErrQueueEmpty
	}

	object := q.objects[0]
	q.objects = q.objects[1:]

	return object, nil
}

func (q *syncQueue[T]) DequeueOrWait(ctx context.Context) (T, error) {
	var object T
	var err error

	object, err = q.Dequeue()
	for errors.Is(err, ErrQueueEmpty) {
		select {
		case <-ctx.Done():
			return *new(T), ctx.Err()
		case <-q.qIsNotEmptySignal:
			object, err = q.Dequeue()
		}
	}

	if err != nil {
		return *new(T), err
	}

	return object, nil
}

func (q *syncQueue[T]) isEmpty() bool {
	return len(q.objects) == 0
}

func NewSyncQueue[T any]() Queue[T] {
	return &syncQueue[T]{
		objects:           []T{},
		lock:              sync.RWMutex{},
		qIsNotEmptySignal: make(chan struct{}, 2),
	}
}
