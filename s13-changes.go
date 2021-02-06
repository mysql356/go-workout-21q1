package main

import "fmt"

type Employee struct {
	name string
	age  int
}

func (e Employee) tch(s string) {
	e.name = s
	fmt.Println(e)
}

func (e *Employee) pch(s string) {
	e.name = s
	fmt.Println(e)
}

func main0() {
	e := Employee{"aa", 20}
	fmt.Println(e)
	e.tch("bb")
	fmt.Println(e)
	(&e).pch("cc")
	fmt.Println(e)
}

func main() {
	e := Employee{"aa", 20}
	p := &e
	fmt.Println(e)

	//v-v
	e.tch("bb")
	fmt.Println(e)
	fmt.Println("---")

	//r-r
	p.pch("cc")
	fmt.Println(e)
	fmt.Println("---")

	//r-v
	p.tch("dd")
	fmt.Println(e)
	fmt.Println("---")

	//v-r
	e.pch("ee")
	fmt.Println(e)

}
