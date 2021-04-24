package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Print("Enter text: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	//fmt.Println(input)
	fmt.Println("Input string: ", input.Text())
}

func main0() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(input.Text())
	}
}

func main2() {

	c := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(input.Text())
		c[input.Text()]++

		//	if input.Text() == "break" {
		//		break
		//	}
	}

	fmt.Println(c)
}
