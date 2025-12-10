package commands

type CliCommand struct {
	Name        string
	Description string
	Options     map[string]bool
	Callback    func(...string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"calcsize": {
			Name:        "calcsize",
			Description: "calculates the size of child directories",
			Options:     make(map[string]bool),
			Callback:    calcSize,
		},
		"help": {
			Name:        "help",
			Description: "how to use filecustodian",
			Options:     make(map[string]bool),
			Callback:    help,
		},
	}
}
