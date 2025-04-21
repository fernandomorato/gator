package main

import (
	"context"
	"fmt"

	"github.com/fernandomorato/gator/rss"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: cli %s", cmd.Name)
	}

	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error fetching feed: %v", err)
	}
	fmt.Println(feed)
	return nil
}
