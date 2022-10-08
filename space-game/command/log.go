package command

import (
	"log"

	"arc-homework/space-game/helper"
)

const messagePattern = "Command %s returned error_handler: %v\n"

type LogErrorCommand struct {
	logger  *log.Logger
	cmdName string
	err     error
}

func (c *LogErrorCommand) Execute() error {
	c.logger.Printf(messagePattern, c.cmdName, c.err)

	return nil
}

func NewLogErrorCommand(logger *log.Logger, cmd Command, err error) *LogErrorCommand {
	return &LogErrorCommand{
		logger:  logger,
		cmdName: helper.GetStructTypeName(cmd),
		err:     err,
	}
}
