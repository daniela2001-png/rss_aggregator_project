package main

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/daniela2001-png/rss_aggregator_project/internal/database"
)

func scrapeFeed(wg *sync.WaitGroup, feed database.Feed, db database.Queries, conn *sql.DB) {
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
	// TODO: Create the posts table and save items/posts there
	posts := ConvertRSSItemsListToDatabaseCreatePostParams(rssFeed.Channel.Item, feed.ID)
	err = db.InsertPostsBulk(context.Background(), conn, posts)
	if err != nil {
		log.Println("error inserting posts: ", err)
		return
	}
}

func startScraping(db database.Queries, concurrency int, timeBetweenRequest time.Duration, conn *sql.DB) {
	log.Printf("Scraping on %d goroutines every %s duration", concurrency, timeBetweenRequest)
	if timeBetweenRequest < 0 {
		log.Fatal("duration must be greater than zero")
	}
	// Unlike timers, which are used for timeouts, "tickers" repeat the execution of a task every n seconds or minutes
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		// context.Background is the global context when we do not have the scoped context
		feeds, err := db.GetNextFeedSToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("error fetching feeds from feeds table")
			continue
		}
		var wg sync.WaitGroup
		log.Println("The lenght of feeds is equal to: ", len(feeds))
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(&wg, feed, db, conn)
		}
		wg.Wait()
	}
}
