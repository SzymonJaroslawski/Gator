package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"time"

	rss "github.com/SzymonJaroslawski/Gator/internal/RSS"
	"github.com/SzymonJaroslawski/Gator/internal/database"
)

func scrapeFeeds(s *State) {
	feed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Print("Couldn't get next feed to fetch", err)
	}
	scrapeFeed(s.Db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetch(context.Background(), feed.ID)
	if err != nil {
		log.Print("Couldn't scrape feed", err)
		return
	}

	feedData, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed: %v", feedData)
		return
	}

	for _, item := range feedData.Channel.Item {
		publisedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publisedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err := db.CreatePost(context.Background(), database.CreatePostParams{
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: publisedAt.Time,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Print("Couldn't create post: ", err)
			continue
		}
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}
