package main

import "fmt"

type rectangle struct {
	l, w int
}

func area(r rectangle)   { fmt.Println(r.l * r.w) }
func areap(r *rectangle) { fmt.Println(r.l * r.w) }

func main() {
	r := rectangle{
		l: 10,
		w: 5,
	}

	p := &r

	//v-v : allowed
	area(r)  //area(*p)

	//r-r : allowed
	areap(p)  //area(&r)

	//r-v : not allowed
	//area(p) - X - not allowed

	//v-r : not allowed
	//areap(r)

}
