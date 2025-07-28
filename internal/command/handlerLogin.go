package command

import (
	"fmt"

	"github.com/nacen-dev/gator/internal/state"
)

func HandlerLogin(s *state.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	userName := cmd.Args[0]
	err := s.Config.SetUser(userName)
	if err != nil {
		return fmt.Errorf("couldn't set the user %w", err)
	}
	fmt.Printf("User %v has been set\n", userName)
	return nil
}
