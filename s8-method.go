package main

import "fmt"

type Employee struct {
	name, city string
	salary     int
}

//method
func (e Employee) display() {
	fmt.Println(e.name, e.city, e.salary)
}

//function
func display(e Employee) {
	fmt.Println(e.name, e.city, e.salary)
}

func main() {
	emp := Employee{
		name:   "aa",
		city:   "aaa",
		salary: 2000,
	}
	fmt.Println(emp)
	display(emp)

	emp.display()
	fmt.Println(emp.name, emp.city, emp.salary, emp.display) //aa aaa 2000 0x4a60c0
	c := emp.display
	fmt.Println(c) //0x4a60c0
	c()            //aa aaa 2000
}
