package main

import "fmt"

type Describe interface {
	Desc()
}

type Person struct {
	name string
	age  int
}

func (p Person) Desc() {
	fmt.Println(p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Desc() {

	fmt.Println(a.state, a.country)
}

func main() {
	var d1 Describe
	p1 := Person{"aa", 20}
	//d1 = &p1 //allowed
	d1 = p1 //allowed
	d1.Desc()

	//address
	var d2 Describe
	a := Address{"dl", "india"}
	//d2 = a //not allowed
	d2 = &a
	d2.Desc()
}
