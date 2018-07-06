package main

import (
	"fmt"
	"os"

	"github.com/gotoronto/rnd-workshop/exercise_2_solution/crawler"
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
	fmt.Println("final list:", crawler.List.URLs)
}
