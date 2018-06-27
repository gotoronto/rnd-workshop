package main

import (
	"os"

	"github.com/gotoronto/rnd-workshop/crawler"
)

func main() {
	args := os.Args[1:]
	done := make(chan int)
	spider := crawler.New(args...)
	spider.Crawl(done)
}
