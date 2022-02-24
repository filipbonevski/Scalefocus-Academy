package main

import "fmt"

func main() {
	var month int
	var year int

	fmt.Println("Enter month")
	fmt.Scanf("%d\n", &month)

	fmt.Println("Enter year")
	fmt.Scanf("%d\n", &year)

	if month > 0 && month < 13 {
		var days, bool = daysInMonth(month, year)
		if bool {
			fmt.Printf("The month you entered has %d days", days)
		}
	} else {
		fmt.Println("Please enter a valid month")
	}

}

func daysInMonth(month int, year int) (int, bool) {
	switch {
	case month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12:
		return 31, true
	case month == 4 || month == 6 || month == 9 || month == 11:
		return 30, true
	case month == 2 && ((year % 4) != 0):
		return 28, true
	case month == 2 && ((year%100) == 0 && (year%400) == 0) || ((year % 100) != 0):
		return 29, true
	default:
		return 28, true
	}
}
