package main

import (
	"bufio"
	"fmt"
	"os"
)

func main0() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(input.Text())
	}
}

func main1() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input)
}

func main() {

	c := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(input.Text())
		c[input.Text()]++

		if input.Text() == "break" {
			break
		}
	}

	fmt.Println(c)
}
