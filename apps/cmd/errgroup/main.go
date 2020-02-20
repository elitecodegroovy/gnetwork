package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func doTask() {
	var g errgroup.Group
	var urls = []string{
		"http://nats.org/",
		"https://www.baidu.com/",
		"https://golangnews.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
}

func main() {

	doTask()
}
