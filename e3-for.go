package main

import "fmt"

func main() {

	sum := 0
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)


	/////////////////////////////
	for i := 1; i <= 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" * ")
		}
		fmt.Println()
	}
}
