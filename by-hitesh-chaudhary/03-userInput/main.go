package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok // err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating ,", input)
	fmt.Printf("Type of ths rating is %T " , input)
}