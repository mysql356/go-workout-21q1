package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	n := 10
	var wg sync.WaitGroup
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All Done")
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println(i)

	sum := 0
	for j := 1; j <= 10; j++ {
		sum += j * i
	}

	time.Sleep(5 * time.Second)
	fmt.Println(i, "G-E", sum)
	wg.Done()
}
