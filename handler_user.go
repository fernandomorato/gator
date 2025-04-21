package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/fernandomorato/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: cli %s <name>", cmd.Name)
	}
	username := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("user %s doesn't exist: %v", username, err)
	}

	err = s.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set user %s: %v", username, err)
	}

	fmt.Printf("user set to: %s\n", username)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: cli %s <name>", cmd.Name)
	}
	username := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("error creating user %s: %v", username, err)
	}

	err = s.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set user %s: %v", username, err)
	}

	log.Printf("created user %s successfully!", username)
	log.Println(user)
	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: cli %s", cmd.Name)
	}

	err := s.db.TruncateUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting users: %v", err)
	}
	log.Print("database reset!")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: cli %s", cmd.Name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting users: %v", err)
	}

	for _, user := range users {
		fmt.Printf("* %s", user.Name)
		if s.config.CurrentUserName == user.Name {
			fmt.Printf(" (current)")
		}
		fmt.Println()
	}
	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: cli %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed %q: %v", url, err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating following feed: %v", err)
	}

	fmt.Printf("user %s is now following feed %s\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: cli %s", cmd.Name)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting followed fields for user %s: %v", user.Name, err)
	}

	fmt.Printf("user %s is following:\n", user.Name)
	for _, feedFollow := range feedFollows {
		fmt.Printf("- %s\n", feedFollow.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: cli %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    url,
	})
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %v", err)
	}

	fmt.Println("feed unfollowed")
	return nil
}
