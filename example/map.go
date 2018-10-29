package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is Maps Example")

	m := make(map[string]int)

	m["key1"] = 11
	m["key2"] = 22
	m["key3"] = 33

	fmt.Println(m)
	fmt.Println(m["key1"])

	delete(m, "key2")
	fmt.Println("After Delete Key2 : ", m)

}
