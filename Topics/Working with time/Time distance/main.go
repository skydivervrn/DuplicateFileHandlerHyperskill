package main

import (
	"fmt"
	"time"
)

const (
	HoursPerDay = 24
	DaysPerYear = 365
	Century     = 100
)

func main() {
	var year, month, day int
	fmt.Scan(&year, &month, &day)

	// Create the `date` using the input `year`, `month`, and `day`:
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	yearDuration := time.Hour * HoursPerDay * DaysPerYear
	centuryDuration := yearDuration * Century

	now := time.Now()

	// Write the missing code below to compare the `date` and `now` variables:
	switch {
	case date.After(time.Now()):
		fmt.Println("This date will be in the future")
	case now.Sub(date) < centuryDuration:
		fmt.Println("This date was in the past hundred years")
	default:
		fmt.Println("It was a long time ago...")
	}
}
