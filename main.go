package main

import (
	"fmt"
	"os"

	"github.com/gotoronto/rnd-workshop/crawler"
)

func main() {
	args := os.Args[1:]
	for _, url := range args {
		if _, err := crawler.Add(url); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Added %s\n", url)
		}
	}
}
