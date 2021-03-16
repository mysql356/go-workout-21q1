package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("\n------\n%T %v", i, i)
}

func main() {

	//1
	a := "aaa"
	describe(a)
	//2
	b := 10
	describe(b)

	//3
	c := struct {
		n string
	}{
		n: "aaa",
	}
	describe(c)

	//4
	var d interface{} = 10
	fmt.Println("\n======")
	assert(d)
}

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}
