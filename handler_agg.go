package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fernandomorato/gator/internal/database"
	"github.com/fernandomorato/gator/rss"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: cli %s <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error parsing time_between_reqs: %v", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	feedRecord, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed to fetch: %v", err)
	}

	feed, err := rss.FetchFeed(context.Background(), feedRecord.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed %s: %v", feedRecord.Name, err)
	}
	s.db.MarkFeedFetched(context.Background(), feedRecord.ID)

	fmt.Printf("Fetched feed %s. The items are:\n", feedRecord.Name)
	for _, item := range feed.Channel.Item {
		fmt.Printf("- %s\n", item.Title)
	}

	fmt.Println()
	fmt.Println("Saving data...")
	for _, item := range feed.Channel.Item {
		date, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			return fmt.Errorf("error parsing item publication date: %v", err)
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: date,
			FeedID: feedRecord.ID,
		})
		if err != nil {
			return fmt.Errorf("could not create post: %v", err)
		}
	}
	fmt.Println("All data saved!")

	return nil
}
