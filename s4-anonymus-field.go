package main

import "fmt"

type Employee struct {
	string
	int
}

func main() {

	e1 := Employee{"aa", 10}
	e2 := &e1

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println((*e2).string, (*e2).int)
}
