package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/nacen-dev/gator/internal/config"
	"github.com/nacen-dev/gator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

func main() {
	cfg, err := config.Read()

	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("unable to connect to the database")
	}
	dbQueries := database.New(db)

	s := state{
		config: &cfg,
		db:     dbQueries,
	}

	commands := commands{
		RegisteredCommands: map[string]func(*state, command) error{},
	}
	commands.Register("login", HandlerLogin)
	commands.Register("register", HandlerRegister)
	commands.Register("reset", HandlerReset)
	commands.Register("users", HandlerGetUsers)
	commands.Register("agg", handlerAgg)
	commands.Register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.Register("feeds", handlerGetFeeds)
	commands.Register("follow", middlewareLoggedIn(handlerFollowFeed))
	commands.Register("following", middlewareLoggedIn(handlerGetFeedFollowsForCurrentUser))
	commands.Register("unfollow", middlewareLoggedIn(handlerUnfollowFeed))

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	err = commands.Run(&s, command{
		Name: commandName,
		Args: commandArgs,
	})

	if err != nil {

		log.Fatalf("unable to run the command %v", err.Error())
	}
}
