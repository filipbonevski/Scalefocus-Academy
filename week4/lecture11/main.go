package main

import (
	"fmt"
	"sync"
)

func processEven(inputs []int) chan int {
	intChan := make(chan int)
	var wg sync.WaitGroup
	go func() {
		for _, i := range inputs {
			wg.Add(1)

			go func(i int) {
				defer wg.Done()
				if i%2 == 0 {
					intChan <- i
				}
			}(i)
		}
		wg.Wait()
		close(intChan)
	}()

	return intChan
}

func processOdd(inputs []int) chan int {
	intChan := make(chan int)
	var wg sync.WaitGroup
	go func() {
		for _, i := range inputs {
			wg.Add(1)

			go func(i int) {
				defer wg.Done()
				if i%2 == 1 {
					intChan <- i
				}
			}(i)
		}
		wg.Wait()
		close(intChan)
	}()

	return intChan
}

func main() {
	inputs := []int{1, 17, 33, 56, 2, 9}
	chEven := processEven(inputs)
	chOdd := processOdd(inputs)

	for even := range chEven {
		fmt.Println("Even:", even)
	}

	for odd := range chOdd {
		fmt.Println("Odd:", odd)
	}
}
