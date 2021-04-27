package main

import . "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	a := intSeq()
	Println(a(), a(), a())

	a = intSeq()
	Println(a(), a(), a(), a())

	b := intSeq()
	Println(b(), b())

	//Closures : a anonymous functions which can access the variables defined outside of function
	c := 5
	func() {
		Println("c =", c)
	}()
}
