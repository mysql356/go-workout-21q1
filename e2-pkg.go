package main

import (
	"fmt"
	"log"
	"workout_jan21/rectangle"
)

var a, b = 5, 10

func init() {

	fmt.Println("main pkg")
	if a < 0 {
		log.Println("Len < 0")
	}
	if b < 0 {
		log.Println("Width < 0")
	}
}

func main() {

	fmt.Printf("%2f", float64(rectangle.Area(5, 2)))
}
