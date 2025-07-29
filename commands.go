package main

import (
	"errors"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	RegisteredCommands map[string]func(*state, command) error
}

func (c *commands) Run(s *state, cmd command) error {
	functionToRun, ok := c.RegisteredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return functionToRun(s, cmd)
}

func (c *commands) Register(name string, f func(*state, command) error) {
	c.RegisteredCommands[name] = f
}
