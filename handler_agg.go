package main

import (
	"context"
	"fmt"
	"github.com/ionutcarp/aggregator/internal/database"
	"log"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s...", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("Couldn't get next feeds to fetch: %v\n", err)
		return
	}
	log.Println("Found a feed to fetch")
	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s as fetched: %v\n", feed.Name, err)
		return
	}
	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't fetch feed %s: %v\n", feed.Name, err)
		return
	}
	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}
