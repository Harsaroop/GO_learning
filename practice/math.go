package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("I am learning go at:", n)

	t := time.Date(2021, time.December, 30, 5, 0, 0, 0, time.UTC)
	fmt.Println("I will be in India at: ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Thu Dec 30 05:00:00 2021")
	fmt.Printf("The type of  parsed Time is: %T", parsedTime)
}
