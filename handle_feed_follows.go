package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nacen-dev/gator/internal/database"
)

func handlerFollowFeed(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}
	feedUrl := cmd.Args[0]
	user, err := s.db.GetUserByName(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to get the user to be added for the feed")
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("unable to get the feed to follow by the url given")
	}

	feedFollow, err := s.db.FollowFeed(context.Background(), database.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't follow the feed for the user: %w", err)
	}
	fmt.Println("Feed follow created:")
	fmt.Printf("* User:          %s\n", feedFollow.FeedName)
	fmt.Printf("* Feed:          %s\n", feedFollow.UserName)

	return nil
}

func handlerGetFeedFollowsForCurrentUser(s *state, cmd command) error {
	user, err := s.db.GetUserByName(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to get the user")
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get the list of followed feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feeds followed for user %s:\n", user.Name)
	for _, feed := range feeds {
		fmt.Printf("* Feed:          %s\n", feed.FeedName)
	}

	return nil
}
