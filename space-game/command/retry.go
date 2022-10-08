package command

type RetryCommand struct {
	cmd Command
}

func (c *RetryCommand) Execute() error {
	return c.cmd.Execute()
}

func NewRetryCommand(cmd Command) *RetryCommand {
	return &RetryCommand{cmd: cmd}
}

type Retry2Command struct {
	cmd Command
}

func (c *Retry2Command) Execute() error {
	return c.cmd.Execute()
}

func NewRetry2Command(cmd Command) *Retry2Command {
	return &Retry2Command{cmd: cmd}
}
