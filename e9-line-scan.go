package main

import (
	"bufio"
	"fmt"
	"os"
)

//once
func main0() {

	fmt.Print("Enter text: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	//fmt.Println(input)
	fmt.Println("Input string: ", input.Text())
}

//multiple
func main1() {
	fmt.Print("Enter text: ")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println("Input string: ", input.Text())
		fmt.Print("Enter text: ")
	}
}

//multiple + condition
func main() {
	fmt.Print("Enter text: ")
	c := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println("Input string: ", input.Text())
		fmt.Print("Enter text: ")
		c[input.Text()]++

		if input.Text() == "break" {
			break
		}
	}

	fmt.Println(c)
}
