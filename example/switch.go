package main

import "fmt"

func main() {

	// Here's a basic `switch`.
	i := 2
	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Example on how switch use for case for same action
	x := 2
	switch x {
	case 1, 2:
		fmt.Println("We are sharing case")
	case 3:
		fmt.Println("three")
	}

	// In Go switch aslo can be replacement for if else
	y := 3
	z := 10
	switch {
	case y < 4:
		fmt.Println("Y is less than ", z)

	default:
		fmt.Println("Y is more than ", z)
	}

}
