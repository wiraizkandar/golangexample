package main

import "fmt"

func main() {

	// Example array in golang
	fmt.Println("This array example!")

	// Basically it almost same as other array in other language
	var a [5]int
	fmt.Println(a)

	// Set value to index
	a[0] = 1
	a[2] = 3
	fmt.Println(a)

	// To get length of an array
	// len(array)
	fmt.Println(len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	for i := 0; i < len(a); i++ {
		fmt.Println("Index a:", a[i])
		fmt.Println("Index b:", b[i])
	}

}
