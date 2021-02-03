package main

import "fmt"

func main() {
	c := make(chan int)
	go prod(c)

	for {
		a, b := <-c
		if b == false {
			break
		}

		fmt.Println(a, "==", b)
	}
}

func prod(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}
