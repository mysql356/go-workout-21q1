package main

import "fmt"

type Employee struct {
	fn, ln      string
	age, salary int
}

func main() {

	//1
	e1 := Employee{"aa", "aaa", 20, 3000}
	e2 := &Employee{"aa", "aaa", 20, 3000}

	e3 := &e1

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)

	fmt.Println(e1.fn, e1.ln, e1.age, e1.salary)
	fmt.Println((*e2).fn)
	fmt.Println((*e3).fn, (*e3).ln)

	//2
	e4 := []Employee{}
	e4 = []Employee{
		{},
		{},
		{"bb", "bb", 20, 3000},
	}
	fmt.Println(e4)
	fmt.Println(e4[2].ln)

	//3
	e5 := []Employee{}
	e5 = []Employee{
		e1, *e2,
	}
	fmt.Println(e5)
	fmt.Println(e5[1].ln)

	//4-pointer
	e41 := &[]Employee{}
	e41 = &[]Employee{
		{},
		{},
		{"bb", "bbb", 20, 3000},
	}
	fmt.Println(e41)
	fmt.Println((*e41)[2].ln)
}
