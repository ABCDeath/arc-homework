package command

import "context"

type RetryCommand struct {
	cmd Command
}

func (c *RetryCommand) Execute(ctx context.Context) error {
	return c.cmd.Execute(ctx)
}

func NewRetryCommand(cmd Command) *RetryCommand {
	return &RetryCommand{cmd: cmd}
}

type Retry2Command struct {
	cmd Command
}

func (c *Retry2Command) Execute(ctx context.Context) error {
	return c.cmd.Execute(ctx)
}

func NewRetry2Command(cmd Command) *Retry2Command {
	return &Retry2Command{cmd: cmd}
}
