package main

import "fmt"

type Employee struct {
	fn, ln      string
	age, salary int
}

func main() {

	e1 := Employee{"aa", "aaa", 20, 3000}
	e2 := &Employee{"aa", "aaa", 20, 3000}

	e3 := &e1

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)

	fmt.Println(e1.fn, e1.ln, e1.age, e1.salary)
	fmt.Println((*e2).fn)
	fmt.Println((*e3).fn, (*e3).ln)

}
