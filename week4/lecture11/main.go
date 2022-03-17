package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func processEven(inputs []int) chan int {
	intChan := make(chan int)
	var wg sync.WaitGroup

	for _, i := range inputs {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

			intChan <- i
			if i%2 == 0 {
				log.Println("input: ", i)
				// intChan <- i
			}

		}(i)
		<-intChan
	}

	close(intChan)
	wg.Wait()
	return intChan
}

func processOdd(inputs []int) chan int {
	intChan := make(chan int)
	var wg sync.WaitGroup

	for _, i := range inputs {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

			intChan <- i
			if i%2 == 1 {
				log.Println("input: ", i)
				// intChan <- i
			}

		}(i)
		<-intChan
	}

	close(intChan)
	wg.Wait()
	return intChan

}

func main() {
	inputs := []int{1, 17, 34, 56, 2, 8}
	chEven := processEven(inputs)
	chOdd := processOdd(inputs)

	for even := range chEven {
		fmt.Println(even)
	}

	for odd := range chOdd {
		fmt.Println(odd)
	}
}
