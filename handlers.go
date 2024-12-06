package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SzymonJaroslawski/Gator/internal/database"
)

const TestRSSURL = "https://www.wagslane.dev/index.xml"

func handleUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Usage: %s <feed_url>", cmd.Name)
	}

	err := s.Db.DeleteFollow(context.Background(), database.DeleteFollowParams{
		Url:    cmd.Args[0],
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	return nil
}

func handleFollowing(s *State, cmd Command, user database.User) error {
	followed_feeds, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feed := range followed_feeds {
		fmt.Println("Followed feeds:")
		fmt.Printf("%s\n", feed.FeedName)
	}

	return nil
}

func handleFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Usage: %s <feed_url>", cmd.Name)
	}

	feed, err := s.Db.GetFeedURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User: %s, ID: %s \n", user.Name, user.ID.String())
	fmt.Printf("Tries to follow feed: %s, ID: %s \n", feed.Name, feed.ID.String())

	feed_follow, err := s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Following:")
	fmt.Printf("Feed: %s", feed_follow.FeedName)
	fmt.Printf("By User: %s", feed_follow.UserName)
	return nil
}

func handleFeeds(s *State, cmd Command) error {
	feeds, err := s.Db.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		name := feed.Name
		URL := feed.Url
		user_id := feed.UserID
		user, err := s.Db.GetUserID(context.Background(), user_id)
		if err != nil {
			return err
		}
		user_name := user.Name
		fmt.Printf("Feed: %s URL: %s Created By: %s\n", name, URL, user_name)
	}

	return nil
}

func handleAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("Usage: %s <feed_name> <feed_url>", cmd.Name)
	}

	feed, err := s.Db.InsertFeed(context.Background(), database.InsertFeedParams{
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Name:   cmd.Args[0],
		Url:    cmd.Args[1],
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Feed added successfuly")
	fmt.Println(feed)

	follow_cmd := Command{
		Name: "follow",
		Args: cmd.Args[1:],
	}

	handleFollow(s, follow_cmd, user)

	return nil
}

func handlerAgg(s *State, cmd Command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("Usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBetween, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	log.Printf("Collecting feeds every: %s...", timeBetween)

	ticker := time.NewTicker(timeBetween)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func handlerUsers(s *State, cmd Command) error {
	users, err := s.Db.GetAllUsersName(context.Background())
	if err != nil {
		return fmt.Errorf("Couldn't get users from database: %w", err)
	}

	msgs := fmt.Sprintf("Users: %d\n", len(users))

	for _, user := range users {
		if user == s.Config.CurrentUserName {
			msgs += fmt.Sprintf("* %s (current)\n", user)
			continue
		}

		msgs += fmt.Sprintf("* %s\n", user)
	}

	fmt.Printf("\n%s", msgs)
	return nil
}

func handlerReset(s *State, cmd Command) error {
	err := s.Db.ResetUser(context.Background())
	if err != nil {
		return fmt.Errorf("Couldn't reset the users database: %w", err)
	}

	fmt.Println("Reseted users database successfuly")
	return nil
}

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	if _, err := s.Db.GetUser(context.Background(), name); err != nil {
		return fmt.Errorf("User: %s does not exist", name)
	}

	err := s.Config.SetUser(name)
	if err != nil {
		return fmt.Errorf("Couldn't set current user: %w", err)
	}

	fmt.Println("Switched user successfuly!")
	return nil
}

func handlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	empty := database.User{}
	if user, err := s.Db.GetUser(context.Background(), name); user != empty {
		if user.Name == name {
			fmt.Println("User already exists")
			os.Exit(1)
			return err
		}
	}

	createParams := database.CreateUserParams{
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Name: name,
	}

	_, err := s.Db.CreateUser(context.Background(), createParams)
	if err != nil {
		return err
	}

	fmt.Println("Registered new user successfuly")
	handlerLogin(s, cmd)
	return nil
}
