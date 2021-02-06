package main

import "fmt"

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	n1 := myInt(5)
	n2 := myInt(10)

	//method + function calling
	sum := n1.add(n2)
	fmt.Println(sum)
}
