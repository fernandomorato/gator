package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fernandomorato/gator/rss"
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
	return nil
}
