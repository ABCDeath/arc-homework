package command

type MacroCommand struct {
	commands []Command
}

func (m *MacroCommand) Execute() error {
	for _, cmd := range m.commands {
		err := cmd.Execute()
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
