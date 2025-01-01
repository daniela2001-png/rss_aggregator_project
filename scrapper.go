package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/daniela2001-png/rss_aggregator_project/internal/database"
)

func scrapeFeed(wg *sync.WaitGroup, feed database.Feed, db database.Queries) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("error marking feed as fetched: ", err)
		return
	}
	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("error fetching feed from RSS api: ", err)
		return
	}
	// TODO: Creat tha posts table and save items/posts there
	for _, item := range rssFeed.Channel.Item {
		log.Println("Found post: ", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}

func startScraping(db database.Queries, concurrency int, timeBetweenRequest time.Duration) {

	log.Printf("Scraping on %d goroutines every %s duration", concurrency, timeBetweenRequest)
	if timeBetweenRequest < 0 {
		log.Fatal("duration must be greater than zero")
	}
	// Unlike timers, which are used for timeouts, "tickers" repeat the execution of a task every n seconds
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		// context.Background is the global context when we do not have the scoped context
		feeds, err := db.GetNextFeedSToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("error fetching feeds from feeds table")
			continue
		}
		var wg sync.WaitGroup
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(&wg, feed, db)
		}
		wg.Wait()
	}
}
