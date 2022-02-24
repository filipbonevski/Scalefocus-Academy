package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cities, prices := citiesAndPrices()
	groupSlices(cities, prices)
	fmt.Println(groupSlices(cities, prices))
}

func groupSlices(keySlice []string, valueSlice []int) map[string][]int {
	a := make(map[string][]int)

	for i, key := range keySlice {
		a[key] = append(a[key], valueSlice[i])
	}

	return a
}

func citiesAndPrices() ([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	dataPointCount := 100

	// randomly choose cities
	cities := make([]string, dataPointCount)
	for i := range cities {
		cities[i] = cityChoices[rand.Intn(len(cityChoices))]
	}

	prices := make([]int, dataPointCount)
	for i := range prices {
		prices[i] = rand.Intn(100)
	}

	return cities, prices
}
