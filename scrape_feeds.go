package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/nacen-dev/gator/internal/database"
)

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("unable to fetch the feed")
		return
	}

	log.Println("Found a feed to fetch")
	scrapeFeed(s.db, feed)

}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("unable to mark the feed %s fetched: %v\n", feed.Name, err)
		return
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}

	fmt.Printf("RSS feed: %v", rssFeed.Channel.Title)
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
		saveScrapedPost(item, feed.ID, db)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))

}

func saveScrapedPost(item RSSItem, feedId uuid.UUID, db *database.Queries) {
	publishedAt := sql.NullTime{}
	if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
		publishedAt = sql.NullTime{
			Time:  t,
			Valid: true,
		}
	}

	_, err := db.CreatePost(context.Background(), database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		FeedID:      feedId,
		Description: sql.NullString{String: item.Description, Valid: true},
		PublishedAt: publishedAt,
		Url:         item.Link,
		Title:       item.Title,
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return
		}
		log.Printf("Couldn't create post: %v", err)
		return
	}
}
