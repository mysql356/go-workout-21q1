package main

import (
	. "fmt"
)

func main() {

	a := 10
	b := []int{10, 20}

	find(10)         //10 []
	find(10, 10)     //10 [10]
	find(10, 10, 20) //10 [10 20]
	find(a, b...)    //10 [10 20]

	find(10, a)    //10 [10]
	find(10, b...) //10 [10 20]

}

func find(x int, y ...int) {
	Println(x, y)
}
