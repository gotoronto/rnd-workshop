Exercise 3
==========

#Goal
This exercise teaches how to use concurrency.

# Task
Implement the `Crawl(url string)` function in the `exercise_3/crawler.go` file.

Some recommendations:
- implement things in smaller chunks
    - Scrape the first url
    - read the responses off of the channel s.Responses
    - iterate the links and scrape each of those
    - keep doing this until there are no new links
    - You can track this by tracking requests (Scrape calls) and responses (channel reads)
    - You can use a map[string]bool to track reuqests
    - when they are equal (no requests in progress) then you can exit
- Please checkout the `exercise_3/crawler_test.go` file to see how it is tested
- Be sure to use go routines and channels

# Resources
- Go Routines https://gobyexample.com/goroutines
- Channels https://gobyexample.com/channels
- Channel Buffering https://gobyexample.com/channel-buffering
- Non-Blocking Channel operations https://gobyexample.com/non-blocking-channel-operations
- Closing channels https://gobyexample.com/closing-channels
- Range over channels https://gobyexample.com/range-over-channels
- Maps https://gobyexample.com/maps

# Note
There is a solution folder, but don't look at it until you have finished the exercise. You will NOT learn much if don't try to solve it yourself.

Feel free to google solutions when you are stuck.

# Example of using go routines and channels
```
package main

import (
  "fmt"
  "net/http"
)

// checkPkg will check existence of a go pkg
func checkPkg(url string, respChan chan string) {
  resp, _ := http.Get(url)
  respChan <- url + ":" + resp.Status
}

func main() {
  pkgs := []string{"os", "net", "fmt", "path"}
  respChan := make(chan string) // a channel to take responses
  // concurrently request each pkg url
  for _, pkg := range pkgs {
    go checkPkg("https://golang.org/pkg/"+pkg, respChan)
  }
  // read each response from the channel
  for respCount := 0; respCount < len(pkgs); respCount++ {
    fmt.Println(<-respChan)
  }
}
```
