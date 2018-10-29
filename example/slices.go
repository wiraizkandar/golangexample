package main

import "fmt"

/*
Almost same like array but not same
*/

func main() {

	fmt.Println("This is slice example")

	// create slice with length 3
	s := make([]string, 3)
	fmt.Println(s)

	// Can set element to slices with position
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	fmt.Println(s)

	// Can copy slice

	c := make([]string, len(s))
	fmt.Println(c)

	// copy (destination , source)
	copy(c, s)
	fmt.Println("This is S as source :", s)
	fmt.Println("This is copied into C :", c)

	// slice[star position , last position]

	l := s[2:5]
	fmt.Println("sl1:", l)
	fmt.Println("sl1:", s[0:2])

	// example two dimension
	twoD := make([][]int, 3)
	fmt.Println(twoD)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
