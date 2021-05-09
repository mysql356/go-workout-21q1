package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan int)
	fmt.Println("Hi")

	for i := 1; i <= 5; i++ {
		go hello(i, done)
	}

	for i := 1; i <= 5; i++ {

		fmt.Println("jobs result ----", <-done)
	}

	fmt.Println("main")
}

func hello(i int, done chan int) {
        //to maintain the order of execution
	var j time.Duration = time.Duration(i)
	time.Sleep(j * 250 * time.Millisecond)
	fmt.Println("Hello ----", i, "---", j)
	done <- i
}
//https://play.golang.org/p/amsY7UfbRjO
