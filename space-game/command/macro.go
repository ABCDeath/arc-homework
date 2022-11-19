package command

import "context"

type MacroCommand struct {
	commands []Command
}

func (m *MacroCommand) Execute(ctx context.Context) error {
	for _, cmd := range m.commands {
		err := cmd.Execute(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewMacroCommand(commands ...Command) *MacroCommand {
	return &MacroCommand{
		commands: commands,
	}
}
