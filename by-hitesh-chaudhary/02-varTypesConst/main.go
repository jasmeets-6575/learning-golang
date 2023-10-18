package main

import "fmt"

func main() {
	var username string = "Jasmeet"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)
	
	var smallFloat float64 = 255.232341412412412345667
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)
}