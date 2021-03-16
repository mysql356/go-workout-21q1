package main

import (
	"fmt"
	"math"
)

func main() {

	findType("aaa")
	findType(10)
	findType(math.Pi)
	p := person{"aa", 20}
	findType(p)
}

func findType(i interface{}) {

	fmt.Printf("====%T===", i)

	switch v := i.(type) {

	case describe:
		v.desc()

	case string:
		fmt.Println(i.(string))

	case int:
		fmt.Println(i.(int))

	default:

		fmt.Println("undefined")
	}
}

type describe interface {
	desc()
}

type person struct {
	name string
	age  int
}

func (p person) desc() {
	fmt.Println(p.name, p.age)
}
