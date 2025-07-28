package main

import (
	"log"
	"os"

	"github.com/nacen-dev/gator/internal/command"
	"github.com/nacen-dev/gator/internal/config"
	"github.com/nacen-dev/gator/internal/state"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := state.State{
		Config: &cfg,
	}

	commands := command.Commands{
		RegisteredCommands: map[string]func(*state.State, command.Command) error{},
	}
	commands.Register("login", command.HandlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	err = commands.Run(&s, command.Command{
		Name: commandName,
		Args: commandArgs,
	})

	if err != nil {
		log.Fatalf("unable to run the command")
	}
}
