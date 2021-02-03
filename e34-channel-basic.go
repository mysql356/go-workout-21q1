package main

import "fmt"

func main() {
	var a chan int
	a = make(chan int)
	fmt.Printf("%T %v", a, a)
}
