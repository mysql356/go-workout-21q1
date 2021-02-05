package main

import "fmt"

type Address struct{ city, state string }
type Contact struct{ email, mobile string }
type Person struct {
	name string
	age  int
	Address
	contact []Contact
}

func main() {
	var p Person
	p.name = "aa"
	p.age = 20
	p.city = "ND"
	p.state = "Delhi"

	fmt.Println(p.name, p.age)
	fmt.Println(p.city)
	fmt.Println(p)

	/////////////////////
	p1 := Person{
		name: "bb",
		age:  30,
		Address: Address{
			city:  "Patna",
			state: "Bihar",
		},
	}
	fmt.Println(p1.name, p1.age)
	fmt.Println(p1.city)
	fmt.Println(p1)

}
