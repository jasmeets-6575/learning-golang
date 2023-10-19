package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
	// "time"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// var numOne int = 2
	// var numTwo float64 = 4

	// fmt.Println("the sum is ", numOne + int(numTwo))
	
	// // random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	// random from crypto
	myRandomNum , _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}