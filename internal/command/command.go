package command

import (
	"errors"

	"github.com/nacen-dev/gator/internal/state"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	RegisteredCommands map[string]func(*state.State, Command) error
}

func (c *Commands) Run(s *state.State, cmd Command) error {
	functionToRun, ok := c.RegisteredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return functionToRun(s, cmd)
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) {
	c.RegisteredCommands[name] = f
}
