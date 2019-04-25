package main

import "fmt"

var glob int32 = 44
var start int
var end int

func main() {
	print2("Wednesday", "Sunday")
}

func print2(startDay, endDay string) {

	for i := checkDay(startDay); i <= checkDay(endDay); i++ {
		switch i {
		case 7:
			fmt.Println("Sunday")
		case 6:
			fmt.Println("Saturday")
		case 5:
			fmt.Println("Friday")
		case 4:
			fmt.Println("Thursday")
		case 3:
			fmt.Println("Wednesday")
		case 2:
			fmt.Println("Tuesday")
		case 1:
			fmt.Println("Monday")
		}
	}
}

func checkDay(startDay string) int {
	switch startDay {
	case "Sunday":
		return 7
	case "Saturday":
		return 6
	case "Friday":
		return 5
	case "Thursday":
		return 4
	case "Wednesday":
		return 3
	case "Tuesday":
		return 2
	case "Monday":
		return 1
	default:
		return 0
	}
}
