package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/fernandomorato/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: cli %s [limit]", cmd.Name)
	}
	limit := 2
	if len(cmd.Args) == 1 {
		arg, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("error setting limit: %v", err)
		}
		limit = arg
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int64(limit),
	})
	if err != nil {
		return fmt.Errorf("could not get posts for user %s: %v", user.Name, err)
	}

	fmt.Println("Retrieved posts:")
	for _, post := range posts {
		fmt.Println(post.Title)
	}
	return nil
}
