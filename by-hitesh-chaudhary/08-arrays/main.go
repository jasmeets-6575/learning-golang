package main

import "fmt"

func main() {
	fmt.Println("Welcome ")
	
	var fruitList [4]string 
	
	fruitList[0] = "Apple"
	fruitList[1] = "banana"
	fruitList[3] = "Orange"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))
	
	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println(len(vegList))
}