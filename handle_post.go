package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/nacen-dev/gator/internal/database"
)

func handlerBrowse(s *state, c command, user database.User) error {
	limit := 2
	if len(c.Args) == 1 {

		if limitArg, err := strconv.Atoi(c.Args[0]); err != nil {
			limit = limitArg
		} else {
			return fmt.Errorf("invalid limit %w", err)
		}
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("unable to get posts for the user")
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		printPost(post)
	}
	return nil
}

func printPost(post database.GetPostsForUserRow) {
	fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
	fmt.Printf("--- %s ---\n", post.Title)
	fmt.Printf("    %v\n", post.Description.String)
	fmt.Printf("Link: %s\n", post.Url)
	fmt.Println("=====================================")
}
