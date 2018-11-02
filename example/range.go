/*
Desc : This sample for go range
*/
package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 2, 3}
	sum := 0

	// Print Value
	for _, num := range nums {
		sum += num
		fmt.Println("Num = ", num)
	}
	fmt.Println("Total Sum : ", sum)

	sum = 0
	// Loop with index
	for index, num := range nums {
		sum += num
		fmt.Println("Index = ", index)
		fmt.Println("Num = ", num)
	}
	fmt.Println("Total Sum : ", sum)

	// Looping in map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// print key
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// loop in string will return unicode value .
	// What is unicode ? : https://unicode-table.com/en/#control-character
	for i, c := range "wiraizkandar" {
		fmt.Println(i, c)
	}

}
