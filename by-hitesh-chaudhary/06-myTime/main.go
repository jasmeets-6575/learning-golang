package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to tme study")

	t := time.Now()

	fmt.Println(t.Year())
 	fmt.Println(t.Month())
 	fmt.Println(t.Day())
 	fmt.Println(t.Hour())
 	fmt.Println(t.Minute())
 	fmt.Println(t.Second())

	presentTime := time.Now()
	
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020,time.October,10,23,11,0,0,time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006  Monday"))
}