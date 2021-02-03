package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	c := make(chan int)

	for input.Scan() {

		i, _ := strconv.Atoi(input.Text())
		go sum(i, c)

		fmt.Println(<-c)

	}
}

func sum(i int, c chan int) {

	c <- i + 10
}
