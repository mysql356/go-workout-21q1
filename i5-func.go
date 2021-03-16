package main

import (
	"fmt"
	"math"
)

func main() {

	findType("aaa")
	findType(10)
	findType(math.Pi)
	p := Person{"aa", 20}
	findType(p)
}

func findType(i interface{}) {

	fmt.Printf("====%T===", i)

	switch v := i.(type) {

	case Describe:
		v.Desc()

	case string:
		fmt.Println(i.(string))

	case int:
		fmt.Println(i.(int))

	default:

		fmt.Println("undefined")
	}
}

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
