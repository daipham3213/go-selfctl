package command

var commands = make(map[string]func() Command)

type Command struct {
	Name     string
	Alias    string
	Help     string
	Template string
}

func RegisterCommand(name string, alias string, help string, template string) {
	commands[name] = func() Command {
		return Command{
			Name:     name,
			Alias:    alias,
			Help:     help,
			Template: template,
		}
	}
}

func GetCommands() []Command {
	var cmdList []Command
	for _, cmdFunc := range commands {
		cmdList = append(cmdList, cmdFunc())
	}
	return cmdList
}
