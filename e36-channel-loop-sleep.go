package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan int)
	fmt.Println("Hi")

	for i := 1; i <= 10; i++ {
		go hello(i, done)
	}

	for i := 1; i <= 10; i++ {

		fmt.Println("jobs result ----", <-done)
	}

	fmt.Println("main")
}

func hello(i int, done chan int) {

	var j time.Duration = time.Duration(i)
	time.Sleep(j * 250 * time.Millisecond)
	fmt.Println("Hello ----", i)
	done <- i
}