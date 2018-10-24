package main

import "fmt"

func main(){
	fmt.Println(1 + 3)
	fmt.Println("1 + 3 = ", 1 + 3)

	if(1 + 3 == 4){
		fmt.Println(true)
	}else{
		fmt.Println(false)
	}

	if 1 + 3 == 4{
		fmt.Println(true)
	}else{
		fmt.Println(false)
	}
}