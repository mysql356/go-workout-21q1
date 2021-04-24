package main

import (
	"fmt"
	"os"
)

func argFunc(arg string) {
	if arg == "php" {
		fmt.Println(arg, "is passed")
	} else if arg == "go" {
		fmt.Println(arg, "is passed")
	} else {
		fmt.Println(arg, "is unknown")
	}
}

func main() {
	args := os.Args[1]
	argFunc(args)
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args)
}
