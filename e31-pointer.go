package main

import . "fmt"

func main() {

	n := 255
	var a *int = &n
	Printf("Type of a =  %T\n", a)
	Println("address of a = ", a)
	Println("a = ", *a)

	b := new(int)
	b = &n
	Printf("Type of b =  %T\n", b)
	Println("address of b = ", b)
	Println("a = ", *b)
}
