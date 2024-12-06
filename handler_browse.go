package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/SzymonJaroslawski/Gator/internal/database"
)

func handleBrowse(s *State, cmd Command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		if arg_limit, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = arg_limit
		} else {
			return fmt.Errorf("Invalid limit: %w", err)
		}
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Found %d posts for user %s\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
