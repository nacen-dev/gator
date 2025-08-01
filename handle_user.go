package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nacen-dev/gator/internal/database"
)

func HandlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	userName := cmd.Args[0]
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	})

	if err != nil {
		return fmt.Errorf("couldn't create the user: %w", err)
	}

	setUserError := s.config.SetUser(user.Name)

	if setUserError != nil {
		return fmt.Errorf("couldn't set the user: %w", err)
	}

	fmt.Printf("User %v has been created and set\n", user.Name)
	printUser(user)

	return nil
}

func HandlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	userName := cmd.Args[0]
	user, err := s.db.GetUserByName(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("user %s doesn't exist", userName)
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set the user %w", err)
	}
	fmt.Printf("User %v has been set\n", user.Name)
	return nil
}

func HandlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable to reset users data")
	}
	return nil
}

func HandlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable to get users")
	}
	for _, user := range users {
		if user.Name == s.config.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}
	}
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
