package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	l, w int
}

type Circle struct {
	r float64
}

func (e Rectangle) Area() int {
	return e.l * e.w
}

func (e Circle) Area() float64 {
	return e.r * e.r * math.Pi
}

func main() {

	r := Rectangle{10, 20}
	c := Circle{12}

	fmt.Println(r.Area())
	fmt.Println(c.Area())
}
