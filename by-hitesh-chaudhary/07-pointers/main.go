package main

import "fmt"

func main() {
	fmt.Println("pointers")
	
	// var ptr *int 
	// fmt.Println("pointers value is ", ptr)
	
	myNum := 23

	var ptr = &myNum
	fmt.Println(ptr)
	fmt.Println(*ptr)
	
	*ptr = *ptr * 2
	fmt.Println(myNum)
}
