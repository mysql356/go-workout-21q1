package main

import (
	"fmt"
	"time"
)

func hello(s string) {
	fmt.Println(s)
}

func main() {

	go hello("11")
	go hello("22")
	hello("33")
	go hello("44")
	go hello("55")

	fmt.Println("main")
	time.Sleep(250 * time.Millisecond)
}

//https://play.golang.org/p/-Djg1Z-SxJl
