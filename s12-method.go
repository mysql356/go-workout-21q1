package main

import "fmt"

type rectangle struct {
	l, w int
}

func (r rectangle) area()   { fmt.Println(r.l * r.w) }
func (r *rectangle) areap() { fmt.Println(r.l * r.w) }

func main() {
	r := rectangle{
		l: 10,
		w: 5,
	}

	p := &r

	//v-v : allowed
	r.area()

	//r-r : allowed
	p.areap()

	//r-v :  allowed
	p.area() 

	//v-r :  allowed
	r.areap()

}
