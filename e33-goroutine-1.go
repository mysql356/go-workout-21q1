package main

import (
	"fmt"
	"runtime"
	"time"
)

func hello(s string) {
	fmt.Println(s)
}

func main() {
	t := time.Now()
	go hello("11")
	go hello("22")
	go hello("33")
	go hello("44")
	go hello("55")

	fmt.Println("NumCPU :", runtime.NumCPU(), " | NumGoroutine :", runtime.NumGoroutine())
	time.Sleep(1 * time.Millisecond)
	fmt.Print("Time Taken : ", time.Since(t))
}

//https://play.golang.org/p/-Djg1Z-SxJl
