package main

import (
	"fmt"
	"sort"
	"time"
)

type timeSlice []time.Time

func (s timeSlice) Less(i, j int) bool { return s[i].Before(s[j]) }
func (s timeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s timeSlice) Len() int           { return len(s) }

func sortDates(format string, dates ...string) ([]string, error) {
	var sortedDates []string
	var dateSlice timeSlice

	for _, date := range dates {
		formattedDate, err := time.Parse(format, date)
		if err != nil {
			fmt.Println("Wrong date input: ", date)
			return nil, err
		}
		dateSlice = append(dateSlice, formattedDate)
	}

	sort.Sort(dateSlice)

	for _, date := range dateSlice {
		sortedDates = append(sortedDates, date.Format("2 Jan 2006"))
	}
	return sortedDates, nil
}

func main() {
	dates := []string{"Sep-14-2018", "Dec-03-2021", "Mar-18-2020", "Apr-18-2005", "May-01-2010", "Jan-12-2017"}
	fmt.Println("Before sorting: ", dates)
	format := "Jan-02-2006"
	fmt.Println("After sorting: ")
	fmt.Println(sortDates(format, dates...))
}
