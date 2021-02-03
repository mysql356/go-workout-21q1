package main

import "fmt"

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

	fmt.Println("Hello ----", i)
	done <- i
}
