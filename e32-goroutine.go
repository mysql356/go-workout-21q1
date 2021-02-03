package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello")
}

func main() {

	go hello()
	fmt.Println("main")
	time.Sleep(1*time.Second)
}
