package main

import "fmt"

func main() {

	var n int
	fmt.Printf("Enter the no: ")
	fmt.Scanf("%d", &n)

	if n%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}