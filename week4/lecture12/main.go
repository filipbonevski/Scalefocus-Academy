package main

import (
	"log"
	"time"
)

func main() {
	out := generateThrottled("foo", 2, time.Second)
	for f := range out {
		log.Println(f)
	}
}

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {
	tick := time.NewTicker(clearInterval)

	out := make(chan string, bufferLimit)

	go func() {
		for range tick.C {
			for i := 0; i < bufferLimit; i++ {
				out <- data
			}
		}
	}()
	return out
}
