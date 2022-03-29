package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func pingURL(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("Got response for %s: %d\n", url, resp.StatusCode)
	return nil
}

func main() {
	var concurrencyLevel int
	flag.IntVar(&concurrencyLevel, "c", 2, "the concurrency level of parsing")
	flag.Parse()

	urls := flag.Args()
	if len(urls) < 1 {
		fmt.Println("URLs are missing")
		return
	}

	processQueue := make(chan string, concurrencyLevel)
	done := make(chan string)

	go func() {
		for _, url := range urls {
			processQueue <- url

			go func(url string) {
				pingURL(url)
				<-processQueue
				done <- url
			}(url)
		}
	}()

	for range urls {
		log.Println("Done: ", <-done)
	}
}
