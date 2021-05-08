package main

import "fmt"

func main() {

	//syncronous
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("main")

	////////////////////////
	go hello(done)
	<-done
	fmt.Println("main")	
}

func hello(done chan bool) {

	fmt.Println("Hello")
	done <- true
}

//https://play.golang.org/p/MDSGQEYO1gE
