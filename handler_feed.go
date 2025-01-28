package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ionutcarp/aggregator/internal/config"
	"github.com/ionutcarp/aggregator/internal/database"
	"log"
	"time"
)

func handlerAddFeed(s *state, cmd command) error {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}
	currentUser, err := s.db.GetUser(context.Background(), cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting current user: %v", err)
	}
	if len(cmd.Args) != 2 {
		return fmt.Errorf("expected two arguments <feed name> <feed url>, got %d", len(cmd.Args))
	}
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    currentUser.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	printFeed(feed)

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %v", err)
	}
	for _, feed := range feeds {
		fmt.Printf("Name: %s\n", feed.Name)
		fmt.Printf("URL: %s\n", feed.Url)
		fmt.Printf("User: %v\n", feed.UserName)
	}
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("ID: %s\n", feed.ID)
	fmt.Printf("URL: %s\n", feed.Url)
	fmt.Printf("UserID: %s\n", feed.UserID)
}
