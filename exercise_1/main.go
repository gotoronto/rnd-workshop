package main

import (
	"fmt"
	"os"

	"github.com/gotoronto/rnd-workshop/exercise_1/crawler"
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
	fmt.Println("final list:", crawler.URLs)
}
