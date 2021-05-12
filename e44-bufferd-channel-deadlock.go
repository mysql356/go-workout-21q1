package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 5
	fmt.Println(<-c)

	c <- 5
	fmt.Println(<-c)

}

//https://play.golang.org/p/Qw8lLJilXLK
