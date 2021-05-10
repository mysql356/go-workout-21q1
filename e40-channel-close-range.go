package main

import "fmt"

func main() {
	c := make(chan int)
	go prod(c)

	for a := range c {
		fmt.Println(a)
	}
}

func prod(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}
//https://play.golang.org/p/Hgjay901R3t
//https://play.golang.org/p/erQJGsuh37U 
