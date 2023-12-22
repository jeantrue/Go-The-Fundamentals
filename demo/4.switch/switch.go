package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {

	switch p := price(); {
	case p < 2:
		fmt.Println("price < 1, Cheap item")
	case p < 10:
		fmt.Println("price < 10, Moderately price item")
	default:
		fmt.Println("Price < 10, Expensive item")
	}

	ticket := Economy
	switch ticket {
	case Economy:
		fmt.Println("Economy seating")
	case Business:
		fmt.Println("Business seating")
	case FirstClass:
		fmt.Println("FirstClass seating")
	default:
		fmt.Println("Other seating")
	}
}
