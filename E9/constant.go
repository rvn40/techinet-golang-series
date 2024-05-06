package main

import "fmt"

func main() {
  const pi = 3.14
	const name string = "Rivan"
    
  fmt.Println(pi)
  fmt.Println(name)


	const (
		gravity     = 9.807 // m/s^2
		daysInWeek  = 7     // days
		hoursInDay  = 24    // hours
		gr          = 1.61  // mathematical constant
	)

	fmt.Println("Acceleration due to gravity:", gravity, "m/s^2")
	fmt.Println("Number of days in a week:", daysInWeek)
	fmt.Println("Number of hours in a day:", hoursInDay)
	fmt.Println("Value of Ggolden ration:", gr)
}