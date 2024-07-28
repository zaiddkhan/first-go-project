package main

import (
	"context"
	"github.com/zaiddkhan/first-go-project/internal/database"
	"log"
	"sync"
	"time"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenScraping time.Duration,
) {
	log.Println("Starting scraping ...")
	ticker := time.NewTicker(timeBetweenScraping)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println(err)
			continue
		}
		wq := &sync.WaitGroup{}
		for _, feed := range feeds {
			wq.Add(1)

			go scrapeFeed(db, wq, feed)
		}
		wq.Wait()
	}

}

func scrapeFeed(db *database.Queries, wq *sync.WaitGroup, feed database.Feed) {
	defer wq.Done()

	_, err := db.MakeFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println(err)
		return
	}
	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println(err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Println(item.Title)
	}
	log.Println("Scraping feed from", rssFeed)

}
