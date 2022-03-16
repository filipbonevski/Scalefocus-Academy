package main

import (
	"fmt"
	"sync"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
	counter     int
	lastElement string
}

func (cp *ConcurrentPrinter) PrintFoo(times int) {
	cp.Add(1)
	go func() {
		defer cp.Done()
		for {
			cp.Lock()
			if cp.counter == times {
				cp.Unlock()
				break
			}
			if cp.lastElement != "foo" {
				fmt.Println("foo")
				cp.lastElement = "foo"
				cp.counter++
			}
			cp.Unlock()
		}
	}()
}

func (cp *ConcurrentPrinter) PrintBar(times int) {
	cp.Add(1)
	go func() {
		defer cp.Done()
		for {
			cp.Lock()
			if cp.counter == times {
				cp.Unlock()
				break
			}
			if cp.lastElement != "bar" {
				fmt.Println("bar")
				cp.lastElement = "bar"
				cp.counter++
			}
			cp.Unlock()
		}
	}()
}

func main() {
	times := 10
	cp := &ConcurrentPrinter{}
	cp.lastElement = "bar"
	cp.PrintFoo(times)
	cp.PrintBar(times)
	cp.Wait()
}
