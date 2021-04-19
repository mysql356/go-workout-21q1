package main

import (
	"fmt"
)

type Employee struct {
	firstName, lastName string
	age, salary         int
}

type Emp struct {
	firstName, lastName string
	age, salary         int
}

func main() {

	//creating structure using field names
	emp1 := Employee{
		firstName: "bb",
		age:       25,
		salary:    500,
		lastName:  "aaa",
	}

	//structA := A{a: 42, b: "foo"}
	//structB := B{A: structA}
	emp2 := Emp(emp1)

	fmt.Println("Employee 1", emp1) //Employee 1 {Sam Anderson 25 500}
	fmt.Println("Employee 2", emp2) //Employee 2 {Thomas Paul 29 800}

	fmt.Printf("%T %v", emp1, emp1)

	fmt.Printf("\n%T %v", emp2, emp2)
}
