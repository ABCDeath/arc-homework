package game_runner

import (
	"context"

	"arc-homework/space-game/command"
	errhandler "arc-homework/space-game/error_handler"
)

type Queue struct {
	q          []command.Command
	errHandler errhandler.Handler
}

func (q *Queue) Enqueue(cmd command.Command) {
	q.q = append(q.q, cmd)
}

func (q *Queue) Dequeue() command.Command {
	if q.IsEmpty() {
		return nil
	}

	cmd := q.q[0]
	q.q = q.q[1:]

	return cmd
}

func (q *Queue) IsEmpty() bool {
	return len(q.q) == 0
}

func (q *Queue) Run() {
	for !q.IsEmpty() {
		q.dequeAndRun()
	}
}

func (q *Queue) dequeAndRun() {
	cmd := q.Dequeue()

	err := cmd.Execute(context.Background())
	if err != nil {
		e := q.errHandler.Handle(cmd, err)
		if e != nil {
			panic(e)
		}
	}
}

func NewQueue(errHandler errhandler.Handler) *Queue {
	return &Queue{
		q:          []command.Command{},
		errHandler: errHandler,
	}
}
