package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gotoronto/rnd-workshop/crawler"
)

func main() {
	args := os.Args[1:]
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	spider := crawler.New(args...)
	spider.Crawl(done)
	fmt.Printf("Scraped %v websites\n", len(spider.List.URLs))
}
